package v8host

//*
import (
	"github.com/gost-dom/v8go"
)

type jsValue = *v8Value
type jsFunction = *v8Function
type jsObject = *v8Object

func assertV8Value(v jsValue) *v8Value    { return v }
func assertV8Object(v jsObject) *v8Object { return v }

/*/

import (
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

type jsValue = js.Value
type jsFunction = js.Function
type jsObject = js.Object

func assertV8Value(v jsValue) *v8Value {
	if r, ok := v.(*v8Value); ok {
		return r
	}
	panic("Expected a V8 Value")
}

func assertV8Object(v jsObject) *v8Object {
	if r, ok := v.(*v8Object); ok {
		return r
	}
	panic("Expected a V8 Object")
}

//*/

type v8Value struct{ *v8go.Value }

func (v v8Value) String() string { return v.Value.String() }
func (v v8Value) Int32() int32   { return v.Value.Int32() }

func (v v8Value) AsFunction() (jsFunction, bool) {
	f, err := v.Value.AsFunction()
	if err == nil {
		return &v8Function{v, f}, true
	}
	return nil, false
}

func (v v8Value) AsObject() (jsObject, bool) {
	o, err := v.Value.AsObject()
	if err == nil {
		return &v8Object{v, o}, true
	}
	return nil, false
}

type v8Function struct {
	v8Value
	v8fn *v8go.Function
}

func (f v8Function) Call(this jsObject, args ...jsValue) (jsValue, error) {
	v8Args := make([]v8go.Valuer, len(args))
	for i, a := range args {
		v8Args[i] = assertV8Value(a)
	}
	var res jsValue
	v, err := f.v8fn.Call(assertV8Object(this).Object, v8Args...)
	if err == nil {
		res = &v8Value{v}
	}
	return res, err
}

/* -------- v8Object -------- */

type v8Object struct {
	v8Value
	*v8go.Object
}

// NativeValue returns the native Go value if any that this JS object is
// wrapping. I.e., for a JS HTMLFormElement, this will return the Go
// HTMLFormElement implementation. Returns nil when no native value is being
// wrapped by this object.
func (o *v8Object) NativeValue() any {
	if o.Object.InternalFieldCount() == 0 {
		return nil
	}
	internal := o.Object.GetInternalField(0)
	defer internal.Release()

	if !internal.IsExternal() {
		return nil
	}

	return internal.ExternalHandle().Value()
}

func (o *v8Object) Get(name string) (jsValue, error) {
	res, err := o.Object.Get(name)
	if err != nil {
		return nil, err
	}
	return &v8Value{res}, nil
}
