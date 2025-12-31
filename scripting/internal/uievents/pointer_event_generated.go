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

func InitializePointerEvent[T any](jsClass js.Class[T]) {
	jsClass.CreateAttribute("width", PointerEvent_width, nil)
	jsClass.CreateAttribute("height", PointerEvent_height, nil)
	jsClass.CreateAttribute("pressure", PointerEvent_pressure, nil)
	jsClass.CreateAttribute("tangentialPressure", PointerEvent_tangentialPressure, nil)
}

func PointerEventConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	type_, errType := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	options, errOpts := js.ConsumeArgument(cbCtx, "options", codec.ZeroValue, codec.DecodeJsObject)
	err = gosterror.First(errType, errOpts)
	if err != nil {
		return nil, err
	}
	var data uievents.PointerEventInit
	e := event.Event{Type: type_}
	if options != nil {
		err = errors.Join(
			codec.DecodeEvent(cbCtx, options, &e),
			decodePointerEventInit(cbCtx, options, &data))
		if err != nil {
			return nil, err
		}
	}
	e.Data = data
	return codec.EncodeConstructedValue(cbCtx, &e)
}

func PointerEvent_width[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.PointerEvent_width: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func PointerEvent_height[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.PointerEvent_height: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func PointerEvent_pressure[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.PointerEvent_pressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func PointerEvent_tangentialPressure[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.PointerEvent_tangentialPressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
