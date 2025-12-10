package sobekengine

import (
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type object struct {
	value
	obj *sobek.Object
}

func newObject(c *scriptContext, o *sobek.Object) jsObject {
	return object{value{c, o}, o}
}

func (o object) Iterator() (js.Value[jsTypeParam], error) {
	v := o.obj.GetSymbol(sobek.SymIterator)
	if v == nil {
		v = sobek.Undefined()
	}
	return newValue(o.ctx, v), nil
}

func (o object) Get(key string) (js.Value[jsTypeParam], error) {
	v := o.obj.Get(key)
	if v == nil {
		v = sobek.Undefined()
	}
	return newValue(o.ctx, v), nil
}

func (o object) Set(key string, v js.Value[jsTypeParam]) error {
	o.obj.Set(key, unwrapValue(v))
	return nil
}

func (o object) Keys() ([]string, error) { return o.obj.Keys(), nil }

func (o object) NativeValue() any {
	instance := o.obj.GetSymbol(o.ctx.wrappedGoObj)
	if instance == nil {
		return nil
	}
	return instance.Export()
}

func (o object) SetNativeValue(value any) {
	o.obj.DefineDataPropertySymbol(
		o.ctx.wrappedGoObj,
		o.ctx.vm.ToValue(value),
		sobek.FLAG_FALSE, // Writable
		sobek.FLAG_FALSE, // Configurable
		sobek.FLAG_FALSE, // Enumerable
	)
}
