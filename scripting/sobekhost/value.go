package sobekhost

import (
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type value struct {
	ctx   *scriptContext
	value sobek.Value
}

type jsTypeParam = value

func toValue(val js.Value[jsTypeParam]) sobek.Value {
	if val == nil {
		return sobek.Undefined()
	}
	return val.Self().value
}

// newValue createa a js.Value[T] wrapping sobek value v. This is safe to use
// on nil values, returning nil if v is nil.
func newValue(ctx *scriptContext, v sobek.Value) js.Value[jsTypeParam] {
	if v == nil {
		return nil
	}
	return value{ctx, v}
}

func (v value) AsFunction() (js.Function[jsTypeParam], bool) {
	f, ok := sobek.AssertFunction(v.value)
	return function{v, f}, ok
}

func (v value) AsObject() (js.Object[jsTypeParam], bool) {
	if o := v.value.ToObject(v.ctx.vm); o != nil {
		return newGojaObject(v.ctx, o), true
	}
	return nil, false
}

func (v value) IsNull() bool { return sobek.IsNull(v.value) }

func (v value) IsUndefined() bool { return sobek.IsUndefined(v.value) }
func (v value) IsString() bool    { return sobek.IsString(v.value) }
func (v value) IsObject() bool    { return sobek.IsNull(v.value) }

func (v value) IsBoolean() bool { _, ok := v.AsObject(); return ok }
func (v value) Self() value     { return v }

func (v value) StrictEquals(other js.Value[jsTypeParam]) bool {
	return v.value.StrictEquals(other.Self().value)
}

func (v value) IsFunction() bool {
	_, ok := sobek.AssertFunction(v.value)
	return ok
}

func (v value) String() string { return v.value.String() }
func (v value) Boolean() bool  { return v.value.ToBoolean() }
func (v value) Int32() int32   { return int32(v.value.ToInteger()) }
func (v value) Uint32() uint32 { return uint32(v.value.ToInteger()) }

type gojaObject struct {
	value
	obj *sobek.Object
}

func newGojaObject(c *scriptContext, o *sobek.Object) js.Object[jsTypeParam] {
	return gojaObject{value{c, o}, o}
}

func (o gojaObject) Get(key string) (js.Value[jsTypeParam], error) {
	v := o.obj.Get(key)
	if v == nil {
		v = sobek.Undefined()
	}
	return newValue(o.ctx, v), nil
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
		sobek.FLAG_FALSE, // Writable
		sobek.FLAG_FALSE, // Configurable
		sobek.FLAG_FALSE, // Enumerable
	)
}

type function struct {
	value
	f sobek.Callable
}

func (f function) Call(
	this js.Object[jsTypeParam],
	args ...js.Value[jsTypeParam],
) (js.Value[jsTypeParam], error) {
	v := make([]sobek.Value, len(args))
	for i, a := range args {
		v[i] = a.Self().value
	}
	res, err := f.f(this.Self().value, v...)
	return newValue(f.ctx, res), err
}
