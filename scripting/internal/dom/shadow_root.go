package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

type shadowRoot[T any] struct{}

func newShadowRoot[T any](host js.ScriptEngine[T]) shadowRoot[T] {
	return shadowRoot[T]{}
}

func shadowRootConstructor[T any](cb js.CallbackContext[T]) (js.Value[T], error) {
	return nil, cb.NewTypeError("Illegal constructor")
}
func InitializeShadowRoot[T any](c js.Class[T]) {}
