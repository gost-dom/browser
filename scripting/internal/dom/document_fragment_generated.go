// This file is generated. Do not edit.

package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

type DocumentFragmentV8Wrapper[T any] struct {
	parentNode *ParentNodeV8Wrapper[T]
}

func NewDocumentFragmentV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *DocumentFragmentV8Wrapper[T] {
	return &DocumentFragmentV8Wrapper[T]{NewParentNodeV8Wrapper(scriptHost)}
}

func (wrapper DocumentFragmentV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w DocumentFragmentV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	w.parentNode.installPrototype(jsClass)
}
