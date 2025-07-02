package codec

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func ZeroValue[T any]() (res T) { return }

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
	if val.IsNull() {
		return nil, nil
	}
	if obj, ok := val.AsObject(); ok {
		if node, ok := obj.NativeValue().(dom.Node); ok {
			return node, nil
		}
	}
	return nil, ctx.NewTypeError("Value is not a node")
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
	return nil, ctx.NewTypeError("Value is not a node")
}

type EventInit struct {
	Bubbles    bool
	Cancelable bool
	Init       any
}

func DecodeEventInit[T any](
	_ js.CallbackContext[T],
	val js.Value[T],
) (EventInit, error) {
	options, ok := val.AsObject()
	if !ok {
		return EventInit{}, errors.New("Not an event init object")
	}

	bubbles, err1 := options.Get("bubbles")
	cancelable, err2 := options.Get("cancelable")
	err := errors.Join(err1, err2)
	if err != nil {
		return EventInit{}, err
	}
	init := EventInit{
		Bubbles:    bubbles.Boolean(),
		Cancelable: cancelable.Boolean(),
	}
	return init, nil
}

func DecodeFunction[T any](cbCtx js.CallbackContext[T], val js.Value[T]) (js.Function[T], error) {
	if f, ok := val.AsFunction(); ok {
		return f, nil
	}
	return nil, cbCtx.NewTypeError("Must be a function")
}
