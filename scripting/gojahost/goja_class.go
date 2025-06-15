package gojahost

import (
	"github.com/dop251/goja"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type gojaClass struct {
	ctx            *GojaContext
	cb             js.FunctionCallback[jsTypeParam]
	prototype      *goja.Object
	indexedHandler *js.IndexedHandlerCallbacks[jsTypeParam]
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
	c.indexedHandler = &oo
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
	c.ctx.storeInternal(native, obj)
	c.installInstance(&obj)

	return newGojaObject(c.ctx, obj), nil
}

// gojaDynamicArray implements [goja.DynamicArray], serving as an indexed
// property handler.
type gojaDynamicArray struct {
	ctx   *GojaContext
	this  *goja.Object
	scope gojaCallbackScope
	cbs   js.IndexedHandlerCallbacks[jsTypeParam]
}

func (o gojaDynamicArray) Get(index int) goja.Value {
	if o.cbs.Getter == nil {
		return nil
	}
	res, err := o.cbs.Getter(o.scope, index)
	if err == js.NotIntercepted {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return toGojaValue(res)
}

func (o gojaDynamicArray) Set(index int, value goja.Value) bool {
	panic("gojaDynamicArray.Set: not implemented")
}

func (o gojaDynamicArray) Len() int {
	if o.cbs.Len == nil {
		return 0
	}
	res, err := o.cbs.Len(o.scope)
	if err == js.NotIntercepted {
		return 0
	}
	if err != nil {
		panic(err)
	}
	return res
}

func (o gojaDynamicArray) SetLen(int) bool {
	panic("gojaDynamicArray.SetLen: not implemented")
}

// gojaDynamicObject implements [goja.DynamicObject], serving as a named
// property handler.
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
	f := o.scope.ValueFactory()
	s := f.NewString(key)
	res, err := o.cbs.Getter(o.scope, s)
	if err == js.NotIntercepted {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return toGojaValue(res)
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
	if o.cbs.Enumerator == nil {
		return nil
	}
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
