package dom

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
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

func Document_implementation[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	implementation, ok := entity.ComponentType[*html.DOMImplementation](instance)
	if !ok {
		implementation = &html.DOMImplementation{}
		entity.SetComponentType(instance, implementation)
	}
	return codec.EncodeEntityScopedWithPrototype(cbCtx, implementation, "DOMImplementation")
}

func Document_createRange[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	r := &dom.Range{}
	r.SetStart(instance, 0)
	r.SetEnd(instance, 0)
	return encodeRange(cbCtx, r)
}
