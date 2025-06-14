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

	namedHandlerCallbacks *js.NamedHandlerCallbacks[jsTypeParam]
}

func (c *gojaClass) assertValid() {
	if c.indexedHandler != nil || c.namedHandlerCallbacks != nil {
		if c.indexedHandler != nil && c.namedHandlerCallbacks != nil {
			panic("Goja mapper doesn't support both a named and indexed handler on the same class")
		}
		if len(c.instanceAttrs) > 0 {
			panic(
				"Goja mapper doesn't support instance attribute accessors or methods when handlers are defined",
			)
		}
	}
}

func (c *gojaClass) CreateIndexedHandler(opts ...js.IndexedHandlerOption[jsTypeParam]) {
	var oo js.IndexedHandlerCallbacks[jsTypeParam]
	for _, o := range opts {
		o(&oo)
	}
	c.indexedHandler = &gojaIndexedHandler{oo.Getter}
	c.assertValid()
}

func (c *gojaClass) CreateNamedHandler(opts ...js.NamedHandlerOption[jsTypeParam]) {
	var cbs js.NamedHandlerCallbacks[jsTypeParam]
	for _, o := range opts {
		o(&cbs)
	}
	c.namedHandlerCallbacks = &cbs
	c.assertValid()
}

func (c *gojaClass) CreateInstanceAttribute(
	name string,
	getter js.FunctionCallback[jsTypeParam],
	setter js.FunctionCallback[jsTypeParam],
) {
	c.instanceAttrs[name] = attributeHandler{c.ctx, name, getter, setter}
	c.assertValid()
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
	if c.namedHandlerCallbacks != nil {
		dynObj := gojaDynamicObject{
			ctx: c.ctx,
			cbs: *c.namedHandlerCallbacks,
		}
		obj := c.ctx.vm.NewDynamicObject(dynObj)
		dynObj.this = obj
		obj.SetPrototype(c.prototype)
		return newGojaObject(c.ctx, obj), nil
	}
	obj := c.ctx.vm.CreateObject(c.prototype)
	c.installInstance(obj)
	c.ctx.storeInternal(native, obj)
	return newGojaObject(c.ctx, obj), nil
}

func newGojaCallbackContext(
	ctx *GojaContext,
	call goja.ConstructorCall,
) *callbackContext {
	return &callbackContext{
		gojaCallbackScope{ctx,
			call.This,
		}, call.Arguments, 0}
}

type gojaDynamicObject struct {
	ctx   *GojaContext
	this  *goja.Object
	scope gojaCallbackScope
	cbs   js.NamedHandlerCallbacks[jsTypeParam]
}

func (o gojaDynamicObject) Get(key string) goja.Value {
	if o.cbs.Getter == nil {
		return nil
	}
	res, err := o.cbs.Getter(o.scope, o.scope.ValueFactory().NewString(key))
	if err == js.NotIntercepted {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return res.Self().value
}

func (o gojaDynamicObject) Delete(key string) (res bool) {
	var err error
	if o.cbs.Deleter != nil {
		res, err = o.cbs.Deleter(o.scope, o.scope.ValueFactory().NewString(key))
		if err == js.NotIntercepted {
			return false
		}
		if err != nil {
			panic(err)
		}
	}
	return res
}

func (o gojaDynamicObject) Has(key string) (res bool) {
	if o.cbs.Getter == nil {
		panic("Must have a getter")
	}
	_, err := o.cbs.Getter(o.scope, o.scope.ValueFactory().NewString(key))
	if err == js.NotIntercepted {
		return false
	}
	if err != nil {
		panic(err)
	}
	return true
}

func (o gojaDynamicObject) Keys() []string {
	keys, err := o.cbs.Enumerator(o.scope)
	if err == js.NotIntercepted {
		return nil
	}
	if err != nil {
		panic(err)
	}
	res := make([]string, len(keys))
	for i, k := range keys {
		res[i] = k.String()
	}
	return res
}

func (o gojaDynamicObject) Set(key string, val goja.Value) bool {
	if o.cbs.Setter == nil {
		return false
	}
	err := o.cbs.Setter(o.scope, o.scope.ValueFactory().NewString(key), newGojaValue(o.ctx, val))
	if err == js.NotIntercepted {
		return false
	}
	if err != nil {
		panic(err)
	}
	return true
}
