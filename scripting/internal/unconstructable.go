package internal

import (
	"github.com/gost-dom/browser/scripting/internal/js"
)

type unconstructable[T any] struct{}

func NewUnconstructable[T any](host js.ScriptEngine[T]) unconstructable[T] {
	return unconstructable[T]{}
}

func (w unconstructable[T]) Constructor(cb js.CallbackContext[T]) (js.Value[T], error) {
	return nil, cb.NewTypeError("Illegal constructor")
}
func (w unconstructable[T]) Initialize(c js.Class[T]) {}
