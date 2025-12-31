// This file is generated. Do not edit.

package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

type DocumentFragment[T any] struct {
	nonElementParentNode NonElementParentNode[T]
	parentNode           ParentNode[T]
}

func NewDocumentFragment[T any](scriptHost js.ScriptEngine[T]) DocumentFragment[T] {
	return DocumentFragment[T]{
		NewNonElementParentNode(scriptHost),
		NewParentNode(scriptHost),
	}
}

func (wrapper DocumentFragment[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w DocumentFragment[T]) installPrototype(jsClass js.Class[T]) {
	w.nonElementParentNode.installPrototype(jsClass)
	w.parentNode.installPrototype(jsClass)
}
