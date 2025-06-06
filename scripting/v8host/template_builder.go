package v8host

import (
	"github.com/gost-dom/browser/scripting/internal/js"
)

type unconstructableV8Wrapper[T any] struct{}

func newUnconstructableV8Wrapper(
	host js.ScriptEngine[jsTypeParam],
) unconstructableV8Wrapper[jsTypeParam] {
	return unconstructableV8Wrapper[jsTypeParam]{}
}

func (w unconstructableV8Wrapper[T]) Constructor(cb js.CallbackContext[T]) (js.Value[T], error) {
	return nil, cb.ValueFactory().NewTypeError("Illegal constructor")
}
func (w unconstructableV8Wrapper[T]) Initialize(c jsClass) {}
