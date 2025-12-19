package fetch

import "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	e.CreateFunction("fetch", Fetch)
	Bootstrap(e)
}
