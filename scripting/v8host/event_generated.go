// This file is generated. Do not edit.

package v8host

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("Event", "", newEventV8Wrapper)
}

func (wrapper eventV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w eventV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
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

func (w eventV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Event.Constructor")
	type_, errArg1 := consumeArgument(cbCtx, "type", nil, w.decodeString)
	eventInitDict, errArg2 := consumeArgument(cbCtx, "eventInitDict", w.defaultEventInit, w.decodeEventInit)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	return w.CreateInstance(cbCtx, type_, eventInitDict)
}

func (w eventV8Wrapper[T]) stopPropagation(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Event.stopPropagation")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.StopPropagation()
	return nil, nil
}

func (w eventV8Wrapper[T]) preventDefault(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Event.preventDefault")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.PreventDefault()
	return nil, nil
}

func (w eventV8Wrapper[T]) type_(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Event.type_")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Type
	return w.toString_(cbCtx, result)
}

func (w eventV8Wrapper[T]) target(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Event.target")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Target
	return w.toEventTarget(cbCtx, result)
}

func (w eventV8Wrapper[T]) currentTarget(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Event.currentTarget")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.CurrentTarget
	return w.toEventTarget(cbCtx, result)
}

func (w eventV8Wrapper[T]) bubbles(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Event.bubbles")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Bubbles
	return w.toBoolean(cbCtx, result)
}

func (w eventV8Wrapper[T]) cancelable(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Event.cancelable")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Cancelable
	return w.toBoolean(cbCtx, result)
}

func (w eventV8Wrapper[T]) defaultPrevented(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Event.defaultPrevented")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.DefaultPrevented
	return w.toBoolean(cbCtx, result)
}
