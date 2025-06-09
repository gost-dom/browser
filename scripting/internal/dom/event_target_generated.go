// This file is generated. Do not edit.

package dom

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type EventTargetV8Wrapper[T any] struct{}

func NewEventTargetV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *EventTargetV8Wrapper[T] {
	return &EventTargetV8Wrapper[T]{}
}

func (wrapper EventTargetV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w EventTargetV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("addEventListener", w.addEventListener)
	jsClass.CreatePrototypeMethod("removeEventListener", w.removeEventListener)
	jsClass.CreatePrototypeMethod("dispatchEvent", w.dispatchEvent)
}

func (w EventTargetV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.Constructor")
	return w.CreateInstance(cbCtx)
}

func (w EventTargetV8Wrapper[T]) addEventListener(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.addEventListener")
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	callback, errArg2 := js.ConsumeArgument(cbCtx, "callback", codec.ZeroValue, w.decodeEventListener)
	options, errArg3 := js.ConsumeArgument(cbCtx, "options", w.defaultEventListenerOptions, w.decodeEventListenerOptions)
	err := errors.Join(errArg1, errArg2, errArg3)
	if err != nil {
		return nil, err
	}
	instance.AddEventListener(type_, callback, options...)
	return nil, nil
}

func (w EventTargetV8Wrapper[T]) removeEventListener(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.removeEventListener")
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	callback, errArg2 := js.ConsumeArgument(cbCtx, "callback", codec.ZeroValue, w.decodeEventListener)
	options, errArg3 := js.ConsumeArgument(cbCtx, "options", w.defaultEventListenerOptions, w.decodeEventListenerOptions)
	err := errors.Join(errArg1, errArg2, errArg3)
	if err != nil {
		return nil, err
	}
	instance.RemoveEventListener(type_, callback, options...)
	return nil, nil
}

func (w EventTargetV8Wrapper[T]) dispatchEvent(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: EventTarget.dispatchEvent")
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	event, errArg1 := js.ConsumeArgument(cbCtx, "event", nil, w.decodeEvent)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.DispatchEvent(event)
	return codec.EncodeBoolean(cbCtx, result)
}
