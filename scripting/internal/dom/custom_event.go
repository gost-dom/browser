package dom

import (
	"errors"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func decodeCustomEventInit[T any](
	s js.Scope[T],
	v js.Value[T],
	init *event.CustomEventInit,
) (err error) {
	if obj, ok := v.AsObject(); ok {
		init.Detail, err = obj.Get("detail")
	}
	return
}

func customEventConstructor[T any](info js.CallbackContext[T]) (js.Value[T], error) {
	arg, ok := info.ConsumeArg()
	if !ok {
		return nil, info.NewTypeError("Must have at least one constructor argument")
	}
	data := event.CustomEventInit{}
	e := &event.Event{Type: arg.String()}
	if options, ok := info.ConsumeArg(); ok {
		if obj, ok := options.AsObject(); ok {
			bubbles, err1 := obj.Get("bubbles")
			cancelable, err2 := obj.Get("cancelable")
			detail, err3 := obj.Get("detail")
			err := errors.Join(err1, err2, err3)
			if err != nil {
				return nil, err
			}
			e.Bubbles = bubbles.Boolean()
			e.Cancelable = cancelable.Boolean()
			data.Detail = detail
		}
	}
	e.Data = data
	info.This().SetNativeValue(e)
	return nil, nil
}

func encodeAny[T any](s js.Scope[T], v any) (js.Value[T], error) {
	return codec.EncodeAny(s, v)
}

func customEvent_detail[T any](info js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[*event.Event](info.Instance())
	if err != nil {
		return nil, err
	}
	if data, ok := instance.Data.(event.CustomEventInit); ok {
		detail, _ := data.Detail.(js.Value[T])
		return detail, nil
	}
	return nil, errors.Join(
		errors.New("Data for custom event was not a CustomEventInit"),
		constants.ErrGostDomBug,
	)
}
