package sobekhost

import (
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type gojaValue struct {
	ctx   *GojaContext
	value sobek.Value
}

type jsTypeParam = gojaValue

func toGojaValue(val js.Value[jsTypeParam]) sobek.Value {
	if val == nil {
		return sobek.Undefined()
	}
	return val.Self().value
}

// newGojaValue createa a js.Value[T] wrapping goja value v. This is safe to use
// on nil values, returning nil if v is nil.
func newGojaValue(ctx *GojaContext, v sobek.Value) js.Value[jsTypeParam] {
	if v == nil {
		return nil
	}
	return gojaValue{ctx, v}
}

func (v gojaValue) AsFunction() (js.Function[jsTypeParam], bool) {
	f, ok := sobek.AssertFunction(v.value)
	return gojaFunction{v, f}, ok
}

func (v gojaValue) AsObject() (js.Object[jsTypeParam], bool) {
	if o := v.value.ToObject(v.ctx.vm); o != nil {
		return newGojaObject(v.ctx, o), true
	}
	return nil, false
}

func (v gojaValue) IsNull() bool { return sobek.IsNull(v.value) }

func (v gojaValue) IsUndefined() bool { return sobek.IsUndefined(v.value) }
func (v gojaValue) IsString() bool    { return sobek.IsString(v.value) }
func (v gojaValue) IsObject() bool    { return sobek.IsNull(v.value) }

func (v gojaValue) IsBoolean() bool { _, ok := v.AsObject(); return ok }
func (v gojaValue) Self() gojaValue { return v }

func (v gojaValue) StrictEquals(other js.Value[jsTypeParam]) bool {
	return v.value.StrictEquals(other.Self().value)
}

func (v gojaValue) IsFunction() bool {
	_, ok := sobek.AssertFunction(v.value)
	return ok
}

func (v gojaValue) String() string { return v.value.String() }
func (v gojaValue) Boolean() bool  { return v.value.ToBoolean() }
func (v gojaValue) Int32() int32   { return int32(v.value.ToInteger()) }
func (v gojaValue) Uint32() uint32 { return uint32(v.value.ToInteger()) }

type gojaObject struct {
	gojaValue
	obj *sobek.Object
}

func newGojaObject(c *GojaContext, o *sobek.Object) js.Object[jsTypeParam] {
	return gojaObject{gojaValue{c, o}, o}
}

func (o gojaObject) Get(key string) (js.Value[jsTypeParam], error) {
	v := o.obj.Get(key)
	if v == nil {
		v = sobek.Undefined()
	}
	return newGojaValue(o.ctx, v), nil
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

type gojaFunction struct {
	gojaValue
	f sobek.Callable
}

func (f gojaFunction) Call(
	this js.Object[jsTypeParam],
	args ...js.Value[jsTypeParam],
) (js.Value[jsTypeParam], error) {
	v := make([]sobek.Value, len(args))
	for i, a := range args {
		v[i] = a.Self().value
	}
	res, err := f.f(this.Self().value, v...)
	return newGojaValue(f.ctx, res), err
}
