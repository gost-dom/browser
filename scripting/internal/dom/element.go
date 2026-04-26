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

func Element_className[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	if className, ok := instance.GetAttribute("class"); ok {
		return codec.EncodeString(cbCtx, className)
	}
	return cbCtx.Null(), nil
}

func Element_setClassName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	className, err := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	if err != nil {
		return nil, err
	}
	instance.SetAttribute("class", className)
	return nil, nil
}
