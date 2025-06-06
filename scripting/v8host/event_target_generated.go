// This file is generated. Do not edit.

package v8host

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type eventTargetV8Wrapper[T any] struct{}

func newEventTargetV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *eventTargetV8Wrapper[T] {
	return &eventTargetV8Wrapper[T]{}
}

func (wrapper eventTargetV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w eventTargetV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("addEventListener", w.addEventListener)
	jsClass.CreatePrototypeMethod("removeEventListener", w.removeEventListener)
	jsClass.CreatePrototypeMethod("dispatchEvent", w.dispatchEvent)
}

func (w eventTargetV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.Constructor")
	return w.CreateInstance(cbCtx)
}

func (w eventTargetV8Wrapper[T]) addEventListener(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.addEventListener")
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	callback, errArg2 := js.ConsumeArgument(cbCtx, "callback", zeroValue, w.decodeEventListener)
	options, errArg3 := js.ConsumeArgument(cbCtx, "options", w.defaultEventListenerOptions, w.decodeEventListenerOptions)
	err := errors.Join(errArg1, errArg2, errArg3)
	if err != nil {
		return nil, err
	}
	instance.AddEventListener(type_, callback, options...)
	return nil, nil
}

func (w eventTargetV8Wrapper[T]) removeEventListener(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.removeEventListener")
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	callback, errArg2 := js.ConsumeArgument(cbCtx, "callback", zeroValue, w.decodeEventListener)
	options, errArg3 := js.ConsumeArgument(cbCtx, "options", w.defaultEventListenerOptions, w.decodeEventListenerOptions)
	err := errors.Join(errArg1, errArg2, errArg3)
	if err != nil {
		return nil, err
	}
	instance.RemoveEventListener(type_, callback, options...)
	return nil, nil
}

func (w eventTargetV8Wrapper[T]) dispatchEvent(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.dispatchEvent")
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	event, errArg1 := js.ConsumeArgument(cbCtx, "event", nil, w.decodeEvent)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.DispatchEvent(event)
	return codec.EncodeBoolean(cbCtx, result)
}
