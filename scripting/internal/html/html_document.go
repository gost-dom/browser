package html

import (
	"github.com/gost-dom/browser/scripting/internal/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLDocumentV8Wrapper[T any] struct {
	dom.DocumentV8Wrapper[T]
}

func NewHTMLDocumentV8Wrapper[T any](host js.ScriptEngine[T]) HTMLDocumentV8Wrapper[T] {
	return HTMLDocumentV8Wrapper[T]{*dom.NewDocumentV8Wrapper(host)}
}

func (w HTMLDocumentV8Wrapper[T]) Constructor(c js.CallbackContext[T]) (js.Value[T], error) {
	return nil, c.ValueFactory().NewTypeError("illegal constructor")
}

func (w HTMLDocumentV8Wrapper[T]) Initialize(c js.Class[T]) {
	w.DocumentV8Wrapper.Initialize(c)
}
