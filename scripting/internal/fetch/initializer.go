package fetch

import "github.com/gost-dom/browser/scripting/internal/js"

func Configure[T any](host js.ScriptEngine[T]) {
	host.CreateFunction("fetch", Fetch)
}
