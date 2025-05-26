package v8host

import (
	"github.com/gost-dom/v8go"
)

type jsValue = *v8Value

//type jsValue = js.Value

type v8Value struct{ *v8go.Value }

func (v v8Value) String() string { return v.Value.String() }
func (v v8Value) Int32() int32   { return v.Value.Int32() }

type v8Function struct {
	v8Value
	v8fn *v8go.Function
}
