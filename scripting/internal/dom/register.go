package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

func Register[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "DOMException", "", newDOMException)
	Bootstrap(e)
}
