package dom

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func CreateDocument[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	res := html.NewHTMLDocument(nil)
	return codec.EncodeConstructedValue(cbCtx, res)
}

func encodeHTMLCollection[T any](
	cbCtx js.CallbackContext[T],
	c dom.HTMLCollection,
) (js.Value[T], error) {
	return cbCtx.Constructor("HTMLCollection").NewInstance(c)
}
