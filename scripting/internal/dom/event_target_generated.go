// This file is generated. Do not edit.

package dom

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type EventTarget[T any] struct{}

func NewEventTarget[T any](scriptHost js.ScriptEngine[T]) *EventTarget[T] {
	return &EventTarget[T]{}
}

func (wrapper EventTarget[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w EventTarget[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("addEventListener", w.addEventListener)
	jsClass.CreatePrototypeMethod("removeEventListener", w.removeEventListener)
	jsClass.CreatePrototypeMethod("dispatchEvent", w.dispatchEvent)
}

func (w EventTarget[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: EventTarget.Constructor", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return w.CreateInstance(cbCtx)
}

func (w EventTarget[T]) addEventListener(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: EventTarget.addEventListener", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	callback, errArg2 := js.ConsumeArgument(cbCtx, "callback", codec.ZeroValue, w.decodeEventListener)
	options, errArg3 := js.ConsumeArgument(cbCtx, "options", w.defaultEventListenerOptions, w.decodeEventListenerOptions)
	err = errors.Join(errArg1, errArg2, errArg3)
	if err != nil {
		return nil, err
	}
	instance.AddEventListener(type_, callback, options...)
	return nil, nil
}

func (w EventTarget[T]) removeEventListener(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: EventTarget.removeEventListener", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, errInst := js.As[event.EventTarget](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	callback, errArg2 := js.ConsumeArgument(cbCtx, "callback", codec.ZeroValue, w.decodeEventListener)
	options, errArg3 := js.ConsumeArgument(cbCtx, "options", w.defaultEventListenerOptions, w.decodeEventListenerOptions)
	err = errors.Join(errArg1, errArg2, errArg3)
	if err != nil {
		return nil, err
	}
	instance.RemoveEventListener(type_, callback, options...)
	return nil, nil
}

func (w EventTarget[T]) dispatchEvent(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: EventTarget.dispatchEvent", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
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
