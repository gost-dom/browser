package sobekengine

import (
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type jsTypeParam = value
type jsValue = js.Value[jsTypeParam]
type jsObject = js.Object[jsTypeParam]
type jsArray = js.Array[jsTypeParam]
type jsFunction = js.Function[jsTypeParam]
type jsError = js.Error[jsTypeParam]

type value struct {
	ctx   *scriptContext
	value sobek.Value
}

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

func (v value) AsObject() (jsObject, bool) {
	if o := v.value.ToObject(v.ctx.vm); o != nil {
		return newObject(v.ctx, o), true
	}
	return nil, false
}

func (v value) IsNull() bool { return sobek.IsNull(v.value) }

func (v value) IsUndefined() bool { return sobek.IsUndefined(v.value) }
func (v value) IsString() bool    { return sobek.IsString(v.value) }

func (v value) IsBoolean() bool {
	// Sobek doesn't expose an IsBoolean function, so resort to calling 'typeof'
	// in JS.
	return v.ctx.typeOf(v) == "boolean"
}

func (v value) IsSymbol() bool {
	// Sobek doesn't expose an IsSymbol function, so resort to calling 'typeof'
	// in JS.
	return v.ctx.typeOf(v) == "symbol"
}
func (v value) IsObject() bool {
	// Sobek doesn't expose an IsObject function, so resort to calling 'typeof'
	// in JS.
	return v.ctx.typeOf(v) == "object"
}

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
