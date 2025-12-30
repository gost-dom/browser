// This file is generated. Do not edit.

package dom

import (
	event "github.com/gost-dom/browser/dom/event"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Event[T any] struct{}

func NewEvent[T any](scriptHost js.ScriptEngine[T]) *Event[T] {
	return &Event[T]{}
}

func (wrapper Event[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Event[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("stopPropagation", Event_stopPropagation)
	jsClass.CreateOperation("preventDefault", Event_preventDefault)
	jsClass.CreateAttribute("type", Event_type, nil)
	jsClass.CreateAttribute("target", Event_target, nil)
	jsClass.CreateAttribute("currentTarget", Event_currentTarget, nil)
	jsClass.CreateAttribute("eventPhase", Event_eventPhase, nil)
	jsClass.CreateAttribute("bubbles", Event_bubbles, nil)
	jsClass.CreateAttribute("cancelable", Event_cancelable, nil)
	jsClass.CreateAttribute("defaultPrevented", Event_defaultPrevented, nil)
}

func EventConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	type_, errType := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	options, errOpts := js.ConsumeArgument(cbCtx, "options", codec.ZeroValue, codec.DecodeJsObject)
	err = gosterror.First(errType, errOpts)
	if err != nil {
		return nil, err
	}
	e := event.Event{Type: type_}
	if options != nil {
		err = codec.DecodeEvent(cbCtx, options, &e)
		if err != nil {
			return nil, err
		}
	}
	return codec.EncodeConstructedValue(cbCtx, &e)
}

func Event_stopPropagation[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.StopPropagation()
	return nil, nil
}

func Event_preventDefault[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.PreventDefault()
	return nil, nil
}

func Event_type[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return codec.EncodeString(cbCtx, result)
}

func Event_target[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return encodeEventTarget(cbCtx, result)
}

func Event_currentTarget[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CurrentTarget
	return encodeEventTarget(cbCtx, result)
}

func Event_eventPhase[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.EventPhase
	return encodeEventPhase(cbCtx, result)
}

func Event_bubbles[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Bubbles
	return codec.EncodeBoolean(cbCtx, result)
}

func Event_cancelable[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Cancelable
	return codec.EncodeBoolean(cbCtx, result)
}

func Event_defaultPrevented[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.DefaultPrevented
	return codec.EncodeBoolean(cbCtx, result)
}
