package dom

import (
	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (e Element[T]) classList(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	cl := instance.ClassList()
	tokenList := cbCtx.Constructor("DOMTokenList")
	return tokenList.NewInstance(cl)
}

func (e *Element[T]) toNamedNodeMap(
	cbCtx js.CallbackContext[T],
	n dom.NamedNodeMap,
) (js.Value[T], error) {
	return codec.EncodeEntity(cbCtx, n)
}

func (e *Element[T]) decodeElement(s js.Scope[T], v js.Value[T]) (dom.Element, error) {
	return codec.DecodeAs[dom.Element](s, v)
}
