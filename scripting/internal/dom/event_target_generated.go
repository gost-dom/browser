// This file is generated. Do not edit.

package dom

import (
	event "github.com/gost-dom/browser/dom/event"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type EventTarget[T any] struct{}

func NewEventTarget[T any](scriptHost js.ScriptEngine[T]) EventTarget[T] {
	return EventTarget[T]{}
}

func (wrapper EventTarget[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w EventTarget[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("addEventListener", EventTarget_addEventListener)
	jsClass.CreateOperation("removeEventListener", EventTarget_removeEventListener)
	jsClass.CreateOperation("dispatchEvent", EventTarget_dispatchEvent)
}

func EventTargetConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return CreateEventTarget(cbCtx)
}

func EventTarget_addEventListener[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	callback, errArg2 := js.ConsumeArgument(cbCtx, "callback", codec.ZeroValue, decodeEventListener)
	options, errArg3 := js.ConsumeArgument(cbCtx, "options", codec.ZeroValue, decodeEventListenerOptions)
	err = gosterror.First(errArg1, errArg2, errArg3)
	if err != nil {
		return nil, err
	}
	instance.AddEventListener(type_, callback, options...)
	return nil, nil
}

func EventTarget_removeEventListener[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	callback, errArg2 := js.ConsumeArgument(cbCtx, "callback", codec.ZeroValue, decodeEventListener)
	options, errArg3 := js.ConsumeArgument(cbCtx, "options", codec.ZeroValue, decodeEventListenerOptions)
	err = gosterror.First(errArg1, errArg2, errArg3)
	if err != nil {
		return nil, err
	}
	instance.RemoveEventListener(type_, callback, options...)
	return nil, nil
}

func EventTarget_dispatchEvent[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	event, errArg1 := js.ConsumeArgument(cbCtx, "event", nil, decodeEvent)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.DispatchEvent(event)
	return codec.EncodeBoolean(cbCtx, result)
}
