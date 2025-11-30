package sobekengine

import (
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type value struct {
	ctx   *scriptContext
	value sobek.Value
}

type jsTypeParam = value

// unwrapValue returns the underlying JS value of v. Returns undefined if v is
// nil.
func unwrapValue(v js.Value[jsTypeParam]) sobek.Value {
	if v == nil {
		return sobek.Undefined()
	}
	return v.Self().value
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
		return newObject(v.ctx, o), true
	}
	return nil, false
}

func (v value) IsNull() bool { return sobek.IsNull(v.value) }

func (v value) IsUndefined() bool { return sobek.IsUndefined(v.value) }
func (v value) IsString() bool    { return sobek.IsString(v.value) }

func (v value) Self() value { return v }

func (v value) StrictEquals(other js.Value[jsTypeParam]) bool {
	return v.value.StrictEquals(unwrapValue(other))
}

func (v value) IsFunction() bool {
	_, ok := sobek.AssertFunction(v.value)
	return ok
}

func (v value) String() string { return v.value.String() }
func (v value) Boolean() bool  { return v.value.ToBoolean() }
func (v value) Int32() int32   { return int32(v.value.ToInteger()) }
func (v value) Uint32() uint32 { return uint32(v.value.ToInteger()) }

type object struct {
	value
	obj *sobek.Object
}

func newObject(c *scriptContext, o *sobek.Object) js.Object[jsTypeParam] {
	return object{value{c, o}, o}
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
		v[i] = unwrapValue(a)
	}
	res, err := f.f(unwrapValue(this), v...)
	return newValue(f.ctx, res), err
}

/* -------- sobekError -------- */

type sobekError struct {
	*sobek.Object
}

func (e sobekError) Error() string {
	return e.Get("message").String()
}
