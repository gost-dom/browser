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
	vm     *goja.Runtime
	clock  *clock.Clock
	window html.Window
	// globals      map[string]function
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
	// panic("Not implemented")
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
	constructor := c.vm.ToValue(class.callback).(*goja.Object)
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

type gojaFunctionCallback = func(call goja.FunctionCall, r *goja.Runtime) *goja.Object

func (class *gojaClass) callback(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	class.installInstance(call.This)
	class.cb(newGojaCallbackContext(class.ctx, call))
	return nil
}

func (class *gojaClass) installInstance(this *goja.Object) {
	for _, v := range class.instanceAttrs {
		v.install(this)
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

type gojaClass struct {
	ctx            *GojaContext
	cb             js.FunctionCallback[jsTypeParam]
	prototype      *goja.Object
	indexedHandler *gojaIndexedHandler
	instanceAttrs  map[string]attributeHandler
}

func (c *gojaClass) CreateIndexedHandler(cb js.HandlerGetterCallback[jsTypeParam, int]) {
	c.indexedHandler = &gojaIndexedHandler{cb}
}

func (c *gojaClass) CreateNamedHandler(cb ...js.NamedHandlerOption[jsTypeParam]) {}

func (c *gojaClass) CreateInstanceAttribute(
	name string,
	getter js.FunctionCallback[jsTypeParam],
	setter js.FunctionCallback[jsTypeParam],
) {
	c.instanceAttrs[name] = attributeHandler{c.ctx, name, getter, setter}
}

func (c gojaClass) CreatePrototypeMethod(
	name string,
	cb js.FunctionCallback[jsTypeParam],
) {
	if err := c.prototype.Set(name, wrapJSCallback(c.ctx, cb)); err != nil {
		panic(err)
	}
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

func (c gojaClass) CreatePrototypeAttribute(
	name string,
	getter js.FunctionCallback[jsTypeParam],
	setter js.FunctionCallback[jsTypeParam],
) {
	attr := attributeHandler{c.ctx, name, getter, setter}
	attr.install(c.prototype)
}

func (c gojaClass) CreateIteratorMethod(cb js.FunctionCallback[jsTypeParam]) {
	c.prototype.SetSymbol(goja.SymIterator, wrapJSCallback(c.ctx, cb))
}

func (c *gojaClass) NewInstance(native any) (js.Object[jsTypeParam], error) {
	obj := c.ctx.vm.CreateObject(c.prototype)
	c.installInstance(obj)
	c.ctx.storeInternal(native, obj)
	return newGojaObject(c.ctx, obj), nil
}

type gojaCallbackContext struct{}

func newGojaCallbackContext(
	ctx *GojaContext,
	call goja.ConstructorCall,
) *callbackContext {
	return &callbackContext{ctx, call.This, call.Arguments, 0}
}
