package dom

import (
	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w DocumentFragment[T]) Constructor(ctx js.CallbackContext[T]) (js.Value[T], error) {
	result := dom.NewDocumentFragment(ctx.Scope().Window().Document())
	return codec.EncodeConstrucedValue(ctx, result)
}
