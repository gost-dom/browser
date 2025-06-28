// This file is generated. Do not edit.

package dom

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
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
	jsClass.CreatePrototypeMethod("stopPropagation", w.stopPropagation)
	jsClass.CreatePrototypeMethod("preventDefault", w.preventDefault)
	jsClass.CreatePrototypeAttribute("type", w.type_, nil)
	jsClass.CreatePrototypeAttribute("target", w.target, nil)
	jsClass.CreatePrototypeAttribute("currentTarget", w.currentTarget, nil)
	jsClass.CreatePrototypeAttribute("eventPhase", w.eventPhase, nil)
	jsClass.CreatePrototypeAttribute("bubbles", w.bubbles, nil)
	jsClass.CreatePrototypeAttribute("cancelable", w.cancelable, nil)
	jsClass.CreatePrototypeAttribute("defaultPrevented", w.defaultPrevented, nil)
}

func (w Event[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Event.Constructor")
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	eventInitDict, errArg2 := js.ConsumeArgument(cbCtx, "eventInitDict", codec.ZeroValue, codec.DecodeEventInit)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	return w.CreateInstance(cbCtx, type_, eventInitDict)
}

func (w Event[T]) stopPropagation(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Event.stopPropagation")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.StopPropagation()
	return nil, nil
}

func (w Event[T]) preventDefault(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Event.preventDefault")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.PreventDefault()
	return nil, nil
}

func (w Event[T]) type_(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Event.type_")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return codec.EncodeString(cbCtx, result)
}

func (w Event[T]) target(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Event.target")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return w.toEventTarget(cbCtx, result)
}

func (w Event[T]) currentTarget(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Event.currentTarget")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CurrentTarget
	return w.toEventTarget(cbCtx, result)
}

func (w Event[T]) bubbles(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Event.bubbles")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Bubbles
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Event[T]) cancelable(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Event.cancelable")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Cancelable
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Event[T]) defaultPrevented(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Event.defaultPrevented")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.DefaultPrevented
	return codec.EncodeBoolean(cbCtx, result)
}
