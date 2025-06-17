package gojahost

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/scripting/internal/js"

	"github.com/dop251/goja"
)

type GojaContext struct {
	vm           *goja.Runtime
	clock        *clock.Clock
	window       html.Window
	classes      map[string]*gojaClass
	wrappedGoObj *goja.Symbol
	cachedNodes  map[int32]goja.Value
}

func (c *GojaContext) Clock() html.Clock { return c.clock }

func (i *GojaContext) Close() {}

func (i *GojaContext) logger() *slog.Logger {
	if i.window == nil {
		return nil
	}
	return i.window.Logger()
}

func (i *GojaContext) run(str string) (goja.Value, error) {
	res, err := i.vm.RunString(str)
	i.clock.Tick()
	return res, err
}

func (i *GojaContext) Run(str string) error {
	_, err := i.run(str)
	return err
}

func (i *GojaContext) Eval(str string) (res any, err error) {
	if gojaVal, err := i.run(str); err == nil {
		return gojaVal.Export(), nil
	} else {
		return nil, err
	}
}

func (i *GojaContext) EvalCore(str string) (res any, err error) {
	return i.vm.RunString(str)
}

func (i *GojaContext) RunFunction(str string, arguments ...any) (res any, err error) {
	var f goja.Value
	if f, err = i.vm.RunString(str); err == nil {
		if c, ok := goja.AssertFunction(f); !ok {
			err = errors.New("GojaContext.RunFunction: script is not a function")
		} else {
			values := make([]goja.Value, len(arguments))
			for i, a := range arguments {
				var ok bool
				if values[i], ok = a.(goja.Value); !ok {
					err = fmt.Errorf("GojaContext.RunFunction: argument %d was not a goja Value", i)
				}
			}
			res, err = c(goja.Undefined(), values...)
		}
	}
	return
}

// Export create a native Go value out of a javascript value. The value argument
// must be a [goja.Value] instance.
//
// This function is intended to be used only for test purposes. The value has an
// [any] type as the tests are not supposed to know the details of the
// underlying engine.
//
// The value is expected to be the ourput of [RunFunction] or [EvalCore]
//
// An error will be returned if the value is not a goja Value, or the value
// could not be converted to a native Go object
func (i *GojaContext) Export(value any) (res any, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("GojaContext.Export: %v", r)
		}
	}()
	if gv, ok := value.(goja.Value); ok {
		res = gv.Export()
	} else {
		err = fmt.Errorf("GojaContext.Export: Value not a goja value: %v", value)
	}
	return
}

func (m *GojaContext) createLocationInstance() *goja.Object {
	location, err := m.classes["Location"].NewInstance(m.window.Location())
	if err != nil {
		panic(err)
	}
	return location.(gojaObject).obj
}

func (c *GojaContext) CreateFunction(name string, cb js.FunctionCallback[jsTypeParam]) {
	c.vm.Set(name, wrapJSCallback(c, cb))
}

func (c *GojaContext) CreateClass(
	name string, extends js.Class[jsTypeParam],
	cb js.FunctionCallback[jsTypeParam],
) js.Class[jsTypeParam] {
	class := &gojaClass{ctx: c, cb: cb, instanceAttrs: make(map[string]attributeHandler)}
	constructor := c.vm.ToValue(class.constructorCb).(*goja.Object)
	constructor.DefineDataProperty(
		"name",
		c.vm.ToValue(name),
		goja.FLAG_NOT_SET,
		goja.FLAG_NOT_SET,
		goja.FLAG_NOT_SET,
	)
	class.prototype = constructor.Get("prototype").(*goja.Object)
	c.vm.Set(name, constructor)
	c.classes[name] = class

	if extends != nil {
		if superclass, ok := extends.(*gojaClass); ok {
			class.prototype.SetPrototype(superclass.prototype)
		} else {
			panic(fmt.Sprintf("Superclass not installed for %s. extends: %+v", name, extends))
		}
	}

	return class
}

func (class *gojaClass) constructorCb(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	class.installInstance(&call.This, nil)
	class.cb(newGojaCallbackContext(class.ctx, call))
	return nil
}

func (class *gojaClass) installInstance(this **goja.Object, native any) {
	for _, v := range class.instanceAttrs {
		v.install(*this)
	}

	if class.namedHandlerCallbacks != nil {
		// This implementation is somewhat fragile if the object need own
		// properties. See comment below.
		obj := *this
		proto := *this
		*this = class.ctx.vm.NewDynamicObject(&gojaDynamicObject{
			ctx:   class.ctx,
			cbs:   *class.namedHandlerCallbacks,
			this:  obj,
			scope: newCallbackScope(class.ctx, proto, native),
		})
		(*this).SetPrototype(class.prototype)
	}
	if class.indexedHandler != nil {
		// TODO: Fix prototype for indexed property handlers. Due to lack of
		// support for internal values in goja, and because a "Dynamic Object"
		// cannot have own symbol properties, an artificial prototype is
		// inserted between the instance and the correct prototype, in order to
		// be able to retrieve the internal instance.
		proto := *this
		*this = class.ctx.vm.NewDynamicArray(&gojaDynamicArray{
			ctx:   class.ctx,
			cbs:   *class.indexedHandler,
			this:  *this,
			scope: newCallbackScope(class.ctx, proto, native),
		})
		(*this).SetPrototype(proto)
	}
}

type gojaIndexedHandler struct {
	cb js.HandlerGetterCallback[jsTypeParam, int]
}

type attributeHandler struct {
	ctx    *GojaContext
	name   string
	getter js.FunctionCallback[jsTypeParam]
	setter js.FunctionCallback[jsTypeParam]
}

func (h attributeHandler) install(object *goja.Object) {
	var g, s goja.Value
	if h.getter != nil {
		g = wrapJSCallback(h.ctx, h.getter)
	}
	if h.setter != nil {
		s = wrapJSCallback(h.ctx, h.setter)
	}
	object.DefineAccessorProperty(h.name, g, s,
		goja.FLAG_TRUE, // configurable
		goja.FLAG_TRUE, // enumerable
	)
}

func newGojaCallbackContext(
	ctx *GojaContext,
	call goja.ConstructorCall,
) *callbackContext {
	return &callbackContext{
		newCallbackScope(ctx, call.This, nil),
		call.Arguments, 0,
	}
}
