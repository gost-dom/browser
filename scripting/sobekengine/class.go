package sobekengine

import (
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type class struct {
	ctx            *scriptContext
	cb             js.CallbackFunc[jsTypeParam]
	name           string
	prototype      *sobek.Object
	indexedHandler *js.IndexedHandlerCallbacks[jsTypeParam]
	instanceAttrs  map[string]attributeHandler

	namedHandlerCallbacks *js.NamedHandlerCallbacks[jsTypeParam]
}

func (c *class) assertValid() {
	if c.indexedHandler != nil || c.namedHandlerCallbacks != nil {
		if c.indexedHandler != nil && c.namedHandlerCallbacks != nil {
			panic(
				"gost-dom/sobek: Sobek doesn't support both a named and indexed handler on the same class",
			)
		}
		if len(c.instanceAttrs) > 0 {
			panic(
				"gost-dom/sobek: Sobek doesn't support instance attribute accessors or methods when handlers are defined",
			)
		}
	}
}

func (c *class) CreateIndexedHandler(opts ...js.IndexedHandlerOption[jsTypeParam]) {
	var oo js.IndexedHandlerCallbacks[jsTypeParam]
	for _, o := range opts {
		o(&oo)
	}
	c.indexedHandler = &oo
	c.assertValid()
}

func (c *class) CreateNamedHandler(opts ...js.NamedHandlerOption[jsTypeParam]) {
	var cbs js.NamedHandlerCallbacks[jsTypeParam]
	for _, o := range opts {
		o(&cbs)
	}
	c.namedHandlerCallbacks = &cbs
	c.assertValid()
}

func (c *class) CreateInstanceAttribute(
	name string,
	getter js.CallbackFunc[jsTypeParam],
	setter js.CallbackFunc[jsTypeParam],
) {
	c.instanceAttrs[name] = attributeHandler{c.ctx, name, getter, setter}
	c.assertValid()
}

func (c class) CreateOperation(
	name string,
	cb js.CallbackFunc[jsTypeParam],
) {

	if err := c.prototype.Set(name, wrapJSCallback(c.ctx, cb.WithLog(c.name, name))); err != nil {
		panic(err)
	}
}

func (c class) CreatePrototypeAttribute(
	name string,
	getter js.CallbackFunc[jsTypeParam],
	setter js.CallbackFunc[jsTypeParam],
) {
	attr := attributeHandler{c.ctx, name, getter, setter}
	attr.install(c.prototype)
}

func (c class) CreateIteratorMethod(cb js.CallbackFunc[jsTypeParam]) {
	c.prototype.SetSymbol(
		sobek.SymIterator,
		wrapJSCallback(c.ctx, cb.WithLog(c.name, "Symbol.Iterator")),
	)
}

func (c *class) NewInstance(native any) (js.Object[jsTypeParam], error) {
	obj := c.ctx.vm.CreateObject(c.prototype)
	c.ctx.storeInternal(native, obj)
	c.installInstance(&obj, native)

	return newObject(c.ctx, obj), nil
}

// dynamicArray implements [sobek.DynamicArray], serving as an indexed
// property handler.
type dynamicArray struct {
	ctx   *scriptContext
	this  *sobek.Object
	scope callbackScope
	cbs   js.IndexedHandlerCallbacks[jsTypeParam]
}

func (o dynamicArray) Get(index int) sobek.Value {
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
	return unwrapValue(res)
}

func (o dynamicArray) Set(index int, value sobek.Value) bool {
	panic("gost-dom/sobek: dynamicArray.Set: not implemented")
}

func (o dynamicArray) Len() int {
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

func (o dynamicArray) SetLen(int) bool {
	panic("gost-dom/sobek: dynamicArray.SetLen: not implemented")
}

// dynamicObject implements [sobek.DynamicObject], serving as a named
// property handler.
type dynamicObject struct {
	ctx   *scriptContext
	this  *sobek.Object
	scope callbackScope
	cbs   js.NamedHandlerCallbacks[jsTypeParam]
}

func (o dynamicObject) Get(key string) sobek.Value {
	if o.cbs.Getter == nil {
		return nil
	}
	s := o.scope.NewString(key)
	res, err := o.cbs.Getter(o.scope, s)
	if err == js.NotIntercepted {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return unwrapValue(res)
}

func (o dynamicObject) Delete(key string) (res bool) {
	var err error
	if o.cbs.Deleter != nil {
		res, err = o.cbs.Deleter(o.scope, o.scope.NewString(key))
		if err == js.NotIntercepted {
			return false
		}
		if err != nil {
			panic(err)
		}
	}
	return res
}

func (o dynamicObject) Has(key string) (res bool) {
	if o.cbs.Getter == nil {
		panic("Must have a getter")
	}
	_, err := o.cbs.Getter(o.scope, o.scope.NewString(key))
	if err == js.NotIntercepted {
		return false
	}
	if err != nil {
		panic(err)
	}
	return true
}

func (o dynamicObject) Keys() []string {
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

func (o dynamicObject) Set(key string, val sobek.Value) bool {
	if o.cbs.Setter == nil {
		return false
	}
	err := o.cbs.Setter(o.scope, o.scope.NewString(key), newValue(o.ctx, val))
	if err == js.NotIntercepted {
		return false
	}
	if err != nil {
		panic(err)
	}
	return true
}
