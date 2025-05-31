// This file is generated. Do not edit.

package v8host

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("EventTarget", "", createEventTargetPrototype)
}

type eventTargetV8Wrapper struct {
	handleReffedObject[event.EventTarget, jsTypeParam]
}

func newEventTargetV8Wrapper(scriptHost *V8ScriptHost) *eventTargetV8Wrapper {
	return &eventTargetV8Wrapper{newHandleReffedObject[event.EventTarget](scriptHost)}
}

func createEventTargetPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newEventTargetV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w eventTargetV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("addEventListener", wrapV8Callback(w.scriptHost, w.addEventListener))
	prototypeTmpl.Set("removeEventListener", wrapV8Callback(w.scriptHost, w.removeEventListener))
	prototypeTmpl.Set("dispatchEvent", wrapV8Callback(w.scriptHost, w.dispatchEvent))
}

func (w eventTargetV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.Constructor")
	return w.CreateInstance(cbCtx)
}

func (w eventTargetV8Wrapper) addEventListener(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.addEventListener")
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	type_, errArg1 := consumeArgument(cbCtx, "type", nil, w.decodeString)
	callback, errArg2 := consumeArgument(cbCtx, "callback", zeroValue, w.decodeEventListener)
	options, errArg3 := consumeArgument(cbCtx, "options", w.defaultEventListenerOptions, w.decodeEventListenerOptions)
	err := errors.Join(errArg1, errArg2, errArg3)
	if err != nil {
		return nil, err
	}
	instance.AddEventListener(type_, callback, options...)
	return nil, nil
}

func (w eventTargetV8Wrapper) removeEventListener(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.removeEventListener")
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	type_, errArg1 := consumeArgument(cbCtx, "type", nil, w.decodeString)
	callback, errArg2 := consumeArgument(cbCtx, "callback", zeroValue, w.decodeEventListener)
	options, errArg3 := consumeArgument(cbCtx, "options", w.defaultEventListenerOptions, w.decodeEventListenerOptions)
	err := errors.Join(errArg1, errArg2, errArg3)
	if err != nil {
		return nil, err
	}
	instance.RemoveEventListener(type_, callback, options...)
	return nil, nil
}

func (w eventTargetV8Wrapper) dispatchEvent(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.dispatchEvent")
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	event, errArg1 := consumeArgument(cbCtx, "event", nil, w.decodeEvent)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.DispatchEvent(event)
	return w.toBoolean(cbCtx, result)
}
