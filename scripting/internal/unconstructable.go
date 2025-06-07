package internal

import (
	"github.com/gost-dom/browser/scripting/internal/js"
)

type unconstructableV8Wrapper[T any] struct{}

func NewUnconstructableV8Wrapper[T any](host js.ScriptEngine[T]) unconstructableV8Wrapper[T] {
	return unconstructableV8Wrapper[T]{}
}

func (w unconstructableV8Wrapper[T]) Constructor(cb js.CallbackContext[T]) (js.Value[T], error) {
	return nil, cb.ValueFactory().NewTypeError("Illegal constructor")
}
func (w unconstructableV8Wrapper[T]) Initialize(c js.Class[T]) {}
