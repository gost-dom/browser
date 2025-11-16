// This file is generated. Do not edit.

package uievents

import (
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (wrapper KeyboardEvent[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w KeyboardEvent[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("getModifierState", w.getModifierState)
	jsClass.CreatePrototypeAttribute("key", w.key, nil)
	jsClass.CreatePrototypeAttribute("code", w.code, nil)
	jsClass.CreatePrototypeAttribute("location", w.location, nil)
	jsClass.CreatePrototypeAttribute("ctrlKey", w.ctrlKey, nil)
	jsClass.CreatePrototypeAttribute("shiftKey", w.shiftKey, nil)
	jsClass.CreatePrototypeAttribute("altKey", w.altKey, nil)
	jsClass.CreatePrototypeAttribute("metaKey", w.metaKey, nil)
	jsClass.CreatePrototypeAttribute("repeat", w.repeat, nil)
	jsClass.CreatePrototypeAttribute("isComposing", w.isComposing, nil)
}

func (w KeyboardEvent[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: KeyboardEvent.Constructor", js.ThisLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call return: KeyboardEvent.Constructor", js.LogAttr("res", res))
	}()
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	eventInitDict, found, errArg := js.ConsumeOptionalArg(cbCtx, "eventInitDict", w.decodeKeyboardEventInit)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceEventInitDict(cbCtx, type_, eventInitDict)
	}
	return w.CreateInstance(cbCtx, type_)
}

func (w KeyboardEvent[T]) getModifierState(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: KeyboardEvent.getModifierState", js.ThisLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call return: KeyboardEvent.getModifierState", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) code(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: KeyboardEvent.code", js.ThisLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call return: KeyboardEvent.code", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.code: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) location(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: KeyboardEvent.location", js.ThisLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call return: KeyboardEvent.location", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.location: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) ctrlKey(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: KeyboardEvent.ctrlKey", js.ThisLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call return: KeyboardEvent.ctrlKey", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.ctrlKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) shiftKey(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: KeyboardEvent.shiftKey", js.ThisLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call return: KeyboardEvent.shiftKey", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.shiftKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) altKey(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: KeyboardEvent.altKey", js.ThisLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call return: KeyboardEvent.altKey", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.altKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) metaKey(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: KeyboardEvent.metaKey", js.ThisLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call return: KeyboardEvent.metaKey", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.metaKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) repeat(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: KeyboardEvent.repeat", js.ThisLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call return: KeyboardEvent.repeat", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.repeat: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w KeyboardEvent[T]) isComposing(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: KeyboardEvent.isComposing", js.ThisLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call return: KeyboardEvent.isComposing", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "KeyboardEvent.isComposing: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
