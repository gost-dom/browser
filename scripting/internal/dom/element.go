package dom

import (
	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func encodeDOMTokenList[T any](cbCtx js.Scope[T], val dom.DOMTokenList) (js.Value[T], error) {
	tokenList := cbCtx.Constructor("DOMTokenList")
	return tokenList.NewInstance(val)
}

func encodeNamedNodeMap[T any](
	cbCtx js.CallbackContext[T],
	n dom.NamedNodeMap,
) (js.Value[T], error) {
	return codec.EncodeEntity(cbCtx, n)
}

func decodeElement[T any](s js.Scope[T], v js.Value[T]) (dom.Element, error) {
	return codec.DecodeAs[dom.Element](s, v)
}
