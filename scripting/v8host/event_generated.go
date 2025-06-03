// This file is generated. Do not edit.

package v8host

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("Event", "", createEventPrototype)
}

func createEventPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newEventV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return constructor
}

func (w eventV8Wrapper) installPrototype(jsClass v8Class) {
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

func (w eventV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Event.Constructor")
	type_, errArg1 := consumeArgument(cbCtx, "type", nil, w.decodeString)
	eventInitDict, errArg2 := consumeArgument(cbCtx, "eventInitDict", w.defaultEventInit, w.decodeEventInit)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	return w.CreateInstance(cbCtx, type_, eventInitDict)
}

func (w eventV8Wrapper) stopPropagation(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Event.stopPropagation")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.StopPropagation()
	return nil, nil
}

func (w eventV8Wrapper) preventDefault(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Event.preventDefault")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.PreventDefault()
	return nil, nil
}

func (w eventV8Wrapper) type_(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Event.type_")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Type
	return w.toString_(cbCtx, result)
}

func (w eventV8Wrapper) target(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Event.target")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Target
	return w.toEventTarget(cbCtx, result)
}

func (w eventV8Wrapper) currentTarget(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Event.currentTarget")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.CurrentTarget
	return w.toEventTarget(cbCtx, result)
}

func (w eventV8Wrapper) bubbles(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Event.bubbles")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Bubbles
	return w.toBoolean(cbCtx, result)
}

func (w eventV8Wrapper) cancelable(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Event.cancelable")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Cancelable
	return w.toBoolean(cbCtx, result)
}

func (w eventV8Wrapper) defaultPrevented(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Event.defaultPrevented")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.DefaultPrevented
	return w.toBoolean(cbCtx, result)
}
