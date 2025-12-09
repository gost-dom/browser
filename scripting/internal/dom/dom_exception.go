package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

type domException[T any] struct{}

func newDOMException[T any](scriptHost js.ScriptEngine[T]) *domException[T] {
	return &domException[T]{}
}

func (w domException[T]) Constructor(info js.CallbackContext[T]) (js.Value[T], error) {
	return nil, nil
}

func (w domException[T]) Initialize(jsClass js.Class[T]) {}
