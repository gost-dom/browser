package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

func Register[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "DOMException", "", newDOMException)
	Bootstrap(reg)
}
