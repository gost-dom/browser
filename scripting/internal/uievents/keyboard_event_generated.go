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

func InitializeKeyboardEvent[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("getModifierState", KeyboardEvent_getModifierState)
	jsClass.CreateAttribute("key", KeyboardEvent_key, nil)
	jsClass.CreateAttribute("code", KeyboardEvent_code, nil)
	jsClass.CreateAttribute("location", KeyboardEvent_location, nil)
	jsClass.CreateAttribute("ctrlKey", KeyboardEvent_ctrlKey, nil)
	jsClass.CreateAttribute("shiftKey", KeyboardEvent_shiftKey, nil)
	jsClass.CreateAttribute("altKey", KeyboardEvent_altKey, nil)
	jsClass.CreateAttribute("metaKey", KeyboardEvent_metaKey, nil)
	jsClass.CreateAttribute("repeat", KeyboardEvent_repeat, nil)
	jsClass.CreateAttribute("isComposing", KeyboardEvent_isComposing, nil)
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

func KeyboardEvent_getModifierState[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.KeyboardEvent_getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func KeyboardEvent_key[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var eventInit uievents.KeyboardEventInit
	eventInit, err = codec.RetrieveEventInit[uievents.KeyboardEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	result := eventInit.Key
	return codec.EncodeString(cbCtx, result)
}

func KeyboardEvent_code[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var eventInit uievents.KeyboardEventInit
	eventInit, err = codec.RetrieveEventInit[uievents.KeyboardEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	result := eventInit.Code
	return codec.EncodeString(cbCtx, result)
}

func KeyboardEvent_location[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var eventInit uievents.KeyboardEventInit
	eventInit, err = codec.RetrieveEventInit[uievents.KeyboardEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	result := eventInit.Location
	return codec.EncodeInt(cbCtx, result)
}

func KeyboardEvent_ctrlKey[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var eventInit uievents.KeyboardEventInit
	eventInit, err = codec.RetrieveEventInit[uievents.KeyboardEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	result := eventInit.CtrlKey
	return codec.EncodeBoolean(cbCtx, result)
}

func KeyboardEvent_shiftKey[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var eventInit uievents.KeyboardEventInit
	eventInit, err = codec.RetrieveEventInit[uievents.KeyboardEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	result := eventInit.ShiftKey
	return codec.EncodeBoolean(cbCtx, result)
}

func KeyboardEvent_altKey[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var eventInit uievents.KeyboardEventInit
	eventInit, err = codec.RetrieveEventInit[uievents.KeyboardEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	result := eventInit.AltKey
	return codec.EncodeBoolean(cbCtx, result)
}

func KeyboardEvent_metaKey[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var eventInit uievents.KeyboardEventInit
	eventInit, err = codec.RetrieveEventInit[uievents.KeyboardEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	result := eventInit.MetaKey
	return codec.EncodeBoolean(cbCtx, result)
}

func KeyboardEvent_repeat[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var eventInit uievents.KeyboardEventInit
	eventInit, err = codec.RetrieveEventInit[uievents.KeyboardEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	result := eventInit.Repeat
	return codec.EncodeBoolean(cbCtx, result)
}

func KeyboardEvent_isComposing[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var eventInit uievents.KeyboardEventInit
	eventInit, err = codec.RetrieveEventInit[uievents.KeyboardEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	result := eventInit.IsComposing
	return codec.EncodeBoolean(cbCtx, result)
}
