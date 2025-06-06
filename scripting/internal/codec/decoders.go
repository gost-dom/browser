package codec

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func DecodeString[T any](cbCtx js.CallbackContext[T], val js.Value[T]) (string, error) {
	return val.String(), nil
}

func DecodeBoolean[T any](_ js.CallbackContext[T], val js.Value[T]) (bool, error) {
	return val.Boolean(), nil
}

func DecodeInt[T any](_ js.CallbackContext[T], val js.Value[T]) (int, error) {
	return int(val.Int32()), nil
}

func DecodeNode[T any](ctx js.CallbackContext[T], val js.Value[T]) (dom.Node, error) {
	if obj, ok := val.AsObject(); ok {
		if node, ok := obj.NativeValue().(dom.Node); ok {
			return node, nil
		}
	}
	return nil, ctx.ValueFactory().NewTypeError("Value is not a node")
}

func DecodeHTMLElement[T any](
	ctx js.CallbackContext[T],
	val js.Value[T],
) (html.HTMLElement, error) {
	if obj, ok := val.AsObject(); ok {
		if res, ok := obj.NativeValue().(html.HTMLElement); ok {
			return res, nil
		}
	}
	return nil, ctx.ValueFactory().NewTypeError("Value is not a node")
}
