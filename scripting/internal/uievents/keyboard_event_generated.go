// This file is generated. Do not edit.

package uievents

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	uievents "github.com/gost-dom/browser/internal/uievents"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (wrapper KeyboardEvent[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w KeyboardEvent[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("getModifierState", w.getModifierState)
	jsClass.CreateAttribute("key", w.key, nil)
	jsClass.CreateAttribute("code", w.code, nil)
	jsClass.CreateAttribute("location", w.location, nil)
	jsClass.CreateAttribute("ctrlKey", w.ctrlKey, nil)
	jsClass.CreateAttribute("shiftKey", w.shiftKey, nil)
	jsClass.CreateAttribute("altKey", w.altKey, nil)
	jsClass.CreateAttribute("metaKey", w.metaKey, nil)
	jsClass.CreateAttribute("repeat", w.repeat, nil)
	jsClass.CreateAttribute("isComposing", w.isComposing, nil)
}

func KeyboardEventConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	type_, errType := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	options, errOpts := js.ConsumeArgument(cbCtx, "options", codec.ZeroValue, codec.DecodeJsObject)
	err = gosterror.First(errType, errOpts)
	if err != nil {
		return nil, err
	}
	var data uievents.KeyboardEventInit
	e := event.Event{Type: type_}
	if options != nil {
		err = errors.Join(
			codec.DecodeEvent(cbCtx, options, &e),
			decodeKeyboardEventInit(cbCtx, options, &data))
		if err != nil {
			return nil, err
		}
	}
	e.Data = data
	return codec.EncodeConstructedValue(cbCtx, &e)
}

func (w KeyboardEvent[T]) getModifierState(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) key(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var eventInit uievents.KeyboardEventInit
	eventInit, err = codec.RetrieveEventInit[uievents.KeyboardEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	result := eventInit.Key
	return codec.EncodeString(cbCtx, result)
}

func (w KeyboardEvent[T]) code(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.code: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) location(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.location: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) ctrlKey(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.ctrlKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) shiftKey(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.shiftKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) altKey(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.altKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) metaKey(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.metaKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) repeat(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.repeat: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) isComposing(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.isComposing: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
