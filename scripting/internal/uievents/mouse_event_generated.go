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

func InitializeMouseEvent[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("getModifierState", MouseEvent_getModifierState)
	jsClass.CreateAttribute("screenX", MouseEvent_screenX, nil)
	jsClass.CreateAttribute("screenY", MouseEvent_screenY, nil)
	jsClass.CreateAttribute("clientX", MouseEvent_clientX, nil)
	jsClass.CreateAttribute("clientY", MouseEvent_clientY, nil)
	jsClass.CreateAttribute("layerX", MouseEvent_layerX, nil)
	jsClass.CreateAttribute("layerY", MouseEvent_layerY, nil)
	jsClass.CreateAttribute("relatedTarget", MouseEvent_relatedTarget, nil)
}

func MouseEventConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	type_, errType := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	options, errOpts := js.ConsumeArgument(cbCtx, "options", codec.ZeroValue, codec.DecodeJsObject)
	err = gosterror.First(errType, errOpts)
	if err != nil {
		return nil, err
	}
	var data uievents.MouseEventInit
	e := event.Event{Type: type_}
	if options != nil {
		err = errors.Join(
			codec.DecodeEvent(cbCtx, options, &e),
			decodeMouseEventInit(cbCtx, options, &data))
		if err != nil {
			return nil, err
		}
	}
	e.Data = data
	return codec.EncodeConstructedValue(cbCtx, &e)
}

func MouseEvent_getModifierState[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.MouseEvent_getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func MouseEvent_screenX[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.MouseEvent_screenX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func MouseEvent_screenY[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.MouseEvent_screenY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func MouseEvent_clientX[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.MouseEvent_clientX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func MouseEvent_clientY[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.MouseEvent_clientY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func MouseEvent_layerX[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.MouseEvent_layerX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func MouseEvent_layerY[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.MouseEvent_layerY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func MouseEvent_relatedTarget[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.MouseEvent_relatedTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
