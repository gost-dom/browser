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
	jsClass.CreateOperation("stopPropagation", w.stopPropagation)
	jsClass.CreateOperation("preventDefault", w.preventDefault)
	jsClass.CreateAttribute("type", w.type_, nil)
	jsClass.CreateAttribute("target", w.target, nil)
	jsClass.CreateAttribute("currentTarget", w.currentTarget, nil)
	jsClass.CreateAttribute("eventPhase", w.eventPhase, nil)
	jsClass.CreateAttribute("bubbles", w.bubbles, nil)
	jsClass.CreateAttribute("cancelable", w.cancelable, nil)
	jsClass.CreateAttribute("defaultPrevented", w.defaultPrevented, nil)
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

func (w Event[T]) stopPropagation(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.StopPropagation()
	return nil, nil
}

func (w Event[T]) preventDefault(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.PreventDefault()
	return nil, nil
}

func (w Event[T]) type_(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return codec.EncodeString(cbCtx, result)
}

func (w Event[T]) target(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return encodeEventTarget(cbCtx, result)
}

func (w Event[T]) currentTarget(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CurrentTarget
	return encodeEventTarget(cbCtx, result)
}

func (w Event[T]) bubbles(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Bubbles
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Event[T]) cancelable(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Cancelable
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Event[T]) defaultPrevented(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.DefaultPrevented
	return codec.EncodeBoolean(cbCtx, result)
}
