// This file is generated. Do not edit.

package uievents

import (
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (wrapper MouseEvent[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w MouseEvent[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("getModifierState", w.getModifierState)
	jsClass.CreatePrototypeAttribute("screenX", w.screenX, nil)
	jsClass.CreatePrototypeAttribute("screenY", w.screenY, nil)
	jsClass.CreatePrototypeAttribute("clientX", w.clientX, nil)
	jsClass.CreatePrototypeAttribute("clientY", w.clientY, nil)
	jsClass.CreatePrototypeAttribute("layerX", w.layerX, nil)
	jsClass.CreatePrototypeAttribute("layerY", w.layerY, nil)
	jsClass.CreatePrototypeAttribute("relatedTarget", w.relatedTarget, nil)
}

func (w MouseEvent[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: MouseEvent.Constructor", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	eventInitDict, found, errArg := js.ConsumeOptionalArg(cbCtx, "eventInitDict", w.decodeMouseEventInit)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceEventInitDict(cbCtx, type_, eventInitDict)
	}
	return w.CreateInstance(cbCtx, type_)
}

func (w MouseEvent[T]) getModifierState(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: MouseEvent.getModifierState", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) screenX(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: MouseEvent.screenX", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.screenX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) screenY(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: MouseEvent.screenY", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.screenY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) clientX(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: MouseEvent.clientX", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.clientX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) clientY(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: MouseEvent.clientY", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.clientY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) layerX(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: MouseEvent.layerX", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.layerX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) layerY(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: MouseEvent.layerY", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.layerY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEvent[T]) relatedTarget(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: MouseEvent.relatedTarget", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "MouseEvent.relatedTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
