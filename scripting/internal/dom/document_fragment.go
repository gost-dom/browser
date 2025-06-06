package dom

import (
	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type DocumentFragmentV8Wrapper[T any] struct {
	parentNode *ParentNodeV8Wrapper[T]
}

func NewDocumentFragmentV8Wrapper[T any](host js.ScriptEngine[T]) DocumentFragmentV8Wrapper[T] {
	return DocumentFragmentV8Wrapper[T]{
		NewParentNodeV8Wrapper(host),
	}
}

func (w DocumentFragmentV8Wrapper[T]) Constructor(ctx js.CallbackContext[T]) (js.Value[T], error) {
	result := dom.NewDocumentFragment(ctx.Scope().Window().Document())
	return codec.EncodeConstrucedValue(ctx, result)
}

func (w DocumentFragmentV8Wrapper[T]) Initialize(class js.Class[T]) {
	w.parentNode.installPrototype(class)
}
