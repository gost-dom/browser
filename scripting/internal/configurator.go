package internal

import (
	"github.com/gost-dom/browser/scripting/internal/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func Configure[T any](host js.ScriptEngine[T]) {
	dom.Configure(host)
}
