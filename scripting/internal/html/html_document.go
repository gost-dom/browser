package html

import (
	"github.com/gost-dom/browser/scripting/internal/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLDocument[T any] struct {
	dom.Document[T]
}

func NewHTMLDocument[T any](host js.ScriptEngine[T]) HTMLDocument[T] {
	return HTMLDocument[T]{*dom.NewDocument(host)}
}

func (w HTMLDocument[T]) Constructor(c js.CallbackContext[T]) (js.Value[T], error) {
	return nil, c.NewTypeError("illegal constructor")
}
