package dom

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type customEvent[T any] struct{}

func NewCustomEvent[T any](scriptHost js.ScriptEngine[T]) *customEvent[T] {
	return &customEvent[T]{}
}

func (w customEvent[T]) Constructor(info js.CallbackContext[T]) (js.Value[T], error) {
	arg, ok := info.ConsumeArg()
	if !ok {
		return info.ReturnWithTypeError("Must have at least one constructor argument")
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
			fmt.Printf("!!BUBBLES: %#v\n", bubbles)
			e.Bubbles = bubbles.Boolean()
			e.Cancelable = cancelable.Boolean()
			data.Detail = detail
		}
	}
	e.Data = data
	info.This().SetNativeValue(e)
	return nil, nil
}

func (w customEvent[T]) Initialize(class js.Class[T]) {
	class.CreatePrototypeAttribute("detail", w.detail, nil)
}

func (w customEvent[T]) detail(info js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[*event.Event](info.Instance())
	if err != nil {
		return nil, err
	}
	if data, ok := instance.Data.(event.CustomEventInit); ok {
		detail, _ := data.Detail.(js.Value[T])
		return detail, nil
	}
	return nil, fmt.Errorf(
		"Data for custom event was not a CustomEventInit. %s",
		constants.BUG_ISSUE_URL,
	)
}
