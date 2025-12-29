// This file is generated. Do not edit.

package uievents

import (
	event "github.com/gost-dom/browser/dom/event"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	uievents "github.com/gost-dom/browser/internal/uievents"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (wrapper MouseEvent[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w MouseEvent[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("getModifierState", w.getModifierState)
	jsClass.CreateAttribute("screenX", w.screenX, nil)
	jsClass.CreateAttribute("screenY", w.screenY, nil)
	jsClass.CreateAttribute("clientX", w.clientX, nil)
	jsClass.CreateAttribute("clientY", w.clientY, nil)
	jsClass.CreateAttribute("layerX", w.layerX, nil)
	jsClass.CreateAttribute("layerY", w.layerY, nil)
	jsClass.CreateAttribute("relatedTarget", w.relatedTarget, nil)
}

func (w MouseEvent[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	type_, errType := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	options, errOpts := js.ConsumeArgument(cbCtx, "options", codec.ZeroValue, codec.DecodeJsObject)
	err = gosterror.First(errType, errOpts)
	if err != nil {
		return nil, err
	}
	var data uievents.MouseEventInit
	e := event.Event{Type: type_}
	if options != nil {
		err = codec.DecodeEvent(cbCtx, options, &e)
		if err != nil {
			return nil, err
		}
		err = decodeMouseEventInit(cbCtx, options, &data)
		if err != nil {
			return nil, err
		}
	}
	e.Data = data
	return codec.EncodeConstructedValue(cbCtx, &e)
}

func (w MouseEvent[T]) getModifierState(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) screenX(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.screenX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) screenY(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.screenY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) clientX(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.clientX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) clientY(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.clientY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) layerX(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.layerX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) layerY(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.layerY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) relatedTarget(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.relatedTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
