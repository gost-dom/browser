package fetch

import "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	Bootstrap(e)
	window, ok := e.Class("Window")
	if !ok {
		panic("Window not installed")
	}
	InitializeWindowOrWorkerGlobalScope(window)
}
