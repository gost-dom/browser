package dom

import "github.com/gost-dom/browser/scripting/internal/js"

func Configure[T any](host js.ScriptEngine[T]) {
	installDOMParser(host)
}
