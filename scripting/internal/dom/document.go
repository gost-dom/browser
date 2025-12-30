package dom

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func CreateDocument[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	res := html.NewHTMLDocument(nil)
	return codec.EncodeConstructedValue(cbCtx, res)
}

func (w *Document[T]) createElement(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	var name string
	name, err1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	instance, err2 := js.As[dom.Document](cbCtx.Instance())
	err := errors.Join(err1, err2)
	if err != nil {
		return nil, err
	}
	return codec.EncodeEntity(cbCtx, instance.CreateElement(name))
}

func encodeHTMLCollection[T any](
	cbCtx js.CallbackContext[T],
	c dom.HTMLCollection,
) (js.Value[T], error) {
	return cbCtx.Constructor("HTMLCollection").NewInstance(c)
}
