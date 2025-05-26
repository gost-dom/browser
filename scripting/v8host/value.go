package v8host

import (
	"github.com/gost-dom/v8go"
)

// "github.com/gost-dom/browser/scripting/internal/js"

type jsValue = *v8Value
type jsFunction = *v8Function
type jsObject = *v8go.Object

// type jsValue = js.Value
// type jsFunction = js.Function

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
	v, err := f.v8fn.Call(this, v8Args...)
	if err == nil {
		res = &v8Value{v}
	}
	return res, err
}

func assertV8Value(v jsValue) *v8Value {
	return v
	// if r, ok := v.(*v8Value); ok {
	// 	return r
	// }
	// panic("Expected a V8 Value")
}
