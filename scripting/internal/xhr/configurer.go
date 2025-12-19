package xhr

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	Bootstrap(e)
}
