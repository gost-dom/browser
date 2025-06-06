package v8host

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type customEventV8Wrapper struct {
}

func newCustomEventV8Wrapper(scriptHost js.ScriptEngine[jsTypeParam]) *customEventV8Wrapper {
	return &customEventV8Wrapper{}
}

func (w customEventV8Wrapper) Constructor(info jsCallbackContext) (jsValue, error) {
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
			e.Bubbles = bubbles.Boolean()
			e.Cancelable = cancelable.Boolean()
			data.Detail = detail
		}
	}
	e.Data = data
	info.This().SetNativeValue(e)
	return nil, nil
}

func (w customEventV8Wrapper) Initialize(class jsClass) {
	class.CreatePrototypeAttribute("detail", w.detail, nil)
}

func (w customEventV8Wrapper) detail(info jsCallbackContext) (jsValue, error) {
	instance, err := js.As[*event.Event](info.Instance())
	if err != nil {
		return nil, err
	}
	if data, ok := instance.Data.(event.CustomEventInit); ok {
		detail, _ := data.Detail.(jsValue)
		return detail, nil
	}
	return nil, fmt.Errorf(
		"Data for custom event was not a CustomEventInit. %s",
		constants.BUG_ISSUE_URL,
	)
}
