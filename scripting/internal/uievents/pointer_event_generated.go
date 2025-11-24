// This file is generated. Do not edit.

package uievents

import (
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (wrapper PointerEvent[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w PointerEvent[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("width", w.width, nil)
	jsClass.CreatePrototypeAttribute("height", w.height, nil)
	jsClass.CreatePrototypeAttribute("pressure", w.pressure, nil)
	jsClass.CreatePrototypeAttribute("tangentialPressure", w.tangentialPressure, nil)
}

func (w PointerEvent[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: PointerEvent.Constructor - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: PointerEvent.Constructor", js.LogAttr("res", res))
	}()
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	eventInitDict, found, errArg := js.ConsumeOptionalArg(cbCtx, "eventInitDict", w.decodePointerEventInit)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceEventInitDict(cbCtx, type_, eventInitDict)
	}
	return w.CreateInstance(cbCtx, type_)
}

func (w PointerEvent[T]) width(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: PointerEvent.width - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: PointerEvent.width", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.width: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w PointerEvent[T]) height(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: PointerEvent.height - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: PointerEvent.height", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.height: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w PointerEvent[T]) pressure(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: PointerEvent.pressure - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: PointerEvent.pressure", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.pressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w PointerEvent[T]) tangentialPressure(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: PointerEvent.tangentialPressure - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: PointerEvent.tangentialPressure", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "PointerEvent.tangentialPressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
