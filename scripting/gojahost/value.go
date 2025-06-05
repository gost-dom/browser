package gojahost

import (
	"github.com/dop251/goja"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type gojaValue struct {
	ctx   *GojaContext
	value goja.Value
}
type jsTypeParam = gojaValue

// newGojaValue createa a js.Value[T] wrapping goja value v. This is safe to use
// on nil values, returning nil if v is nil.
func newGojaValue(ctx *GojaContext, v goja.Value) js.Value[jsTypeParam] {
	if v == nil {
		return nil
	}
	return gojaValue{ctx, v}
}

func (v gojaValue) AsFunction() (js.Function[jsTypeParam], bool) { panic("TODO") }

func (v gojaValue) AsObject() (js.Object[jsTypeParam], bool) {
	if o := v.value.ToObject(v.ctx.vm); o != nil {
		return newGojaObject(v.ctx, o), true
	}
	return nil, false
}

func (v gojaValue) IsNull() bool { return goja.IsNull(v.value) }

func (v gojaValue) IsUndefined() bool { return goja.IsUndefined(v.value) }
func (v gojaValue) IsString() bool    { return goja.IsString(v.value) }
func (v gojaValue) IsObject() bool    { return goja.IsNull(v.value) }

func (v gojaValue) IsBoolean() bool { _, ok := v.AsObject(); return ok }
func (v gojaValue) Self() gojaValue { return v }

func (v gojaValue) StrictEquals(other js.Value[jsTypeParam]) bool {
	return v.value.StrictEquals(other.Self().value)
}

func (v gojaValue) IsFunction() bool {
	_, ok := goja.AssertFunction(v.value)
	return ok
}

func (v gojaValue) String() string { return v.value.String() }
func (v gojaValue) Boolean() bool  { return v.value.ToBoolean() }
func (v gojaValue) Int32() int32   { return int32(v.value.ToInteger()) }
func (v gojaValue) Uint32() uint32 { return uint32(v.value.ToInteger()) }

type gojaObject struct {
	gojaValue
	obj *goja.Object
}

func newGojaObject(c *GojaContext, o *goja.Object) js.Object[jsTypeParam] {
	return gojaObject{gojaValue{c, o}, o}
}

func (o gojaObject) Get(key string) (js.Value[jsTypeParam], error) {
	return newGojaValue(o.ctx, o.obj.Get(key)), nil
}

func (o gojaObject) Set(key string, v js.Value[jsTypeParam]) error {
	o.obj.Set(key, v.Self().value)
	return nil
}

func (o gojaObject) Keys() ([]string, error) { return o.obj.Keys(), nil }

func (o gojaObject) NativeValue() any {
	instance := o.obj.GetSymbol(o.ctx.wrappedGoObj)
	if instance == nil {
		return nil
	}
	return instance.Export()
}

func (o gojaObject) SetNativeValue(value any) {
	o.obj.DefineDataPropertySymbol(
		o.ctx.wrappedGoObj,
		o.ctx.vm.ToValue(value),
		goja.FLAG_FALSE, // Writable
		goja.FLAG_FALSE, // Configurable
		goja.FLAG_FALSE, // Enumerable
	)
}
