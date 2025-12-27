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

func (wrapper PointerEvent[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w PointerEvent[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateAttribute("width", w.width, nil)
	jsClass.CreateAttribute("height", w.height, nil)
	jsClass.CreateAttribute("pressure", w.pressure, nil)
	jsClass.CreateAttribute("tangentialPressure", w.tangentialPressure, nil)
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

func (w PointerEvent[T]) width(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.width: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w PointerEvent[T]) height(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.height: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w PointerEvent[T]) pressure(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.pressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w PointerEvent[T]) tangentialPressure(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.tangentialPressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
