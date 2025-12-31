package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

func Register[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "DOMException", "", InitializeDomException, domExceptionConstructor)
	Bootstrap(e)
}
