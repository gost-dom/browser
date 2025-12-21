package fileapi

import (
	"github.com/gost-dom/browser/scripting/internal/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func ConfigureWindowRealm[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "File", "", dom.NewEvent)
}
