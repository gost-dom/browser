package streams

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureWindowRealm[T any](e js.ScriptEngine[T]) {
	Bootstrap(e)
}
