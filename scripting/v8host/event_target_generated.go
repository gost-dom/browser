// This file is generated. Do not edit.

package v8host

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("EventTarget", "", newEventTargetV8Wrapper)
}

type eventTargetV8Wrapper struct {
	handleReffedObject[event.EventTarget, jsTypeParam]
}

func newEventTargetV8Wrapper(scriptHost jsScriptEngine) *eventTargetV8Wrapper {
	return &eventTargetV8Wrapper{newHandleReffedObject[event.EventTarget](scriptHost)}
}

func (wrapper eventTargetV8Wrapper) initialize(jsClass jsClass) {
	wrapper.installPrototype(jsClass)
}

func (w eventTargetV8Wrapper) installPrototype(jsClass jsClass) {
	jsClass.CreatePrototypeMethod("addEventListener", w.addEventListener)
	jsClass.CreatePrototypeMethod("removeEventListener", w.removeEventListener)
	jsClass.CreatePrototypeMethod("dispatchEvent", w.dispatchEvent)
}

func (w eventTargetV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.constructor")
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
