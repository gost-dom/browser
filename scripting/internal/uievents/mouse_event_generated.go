// This file is generated. Do not edit.

package uievents

import (
	"errors"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (wrapper MouseEventV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w MouseEventV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("getModifierState", w.getModifierState)
	jsClass.CreatePrototypeAttribute("screenX", w.screenX, nil)
	jsClass.CreatePrototypeAttribute("screenY", w.screenY, nil)
	jsClass.CreatePrototypeAttribute("clientX", w.clientX, nil)
	jsClass.CreatePrototypeAttribute("clientY", w.clientY, nil)
	jsClass.CreatePrototypeAttribute("layerX", w.layerX, nil)
	jsClass.CreatePrototypeAttribute("layerY", w.layerY, nil)
	jsClass.CreatePrototypeAttribute("relatedTarget", w.relatedTarget, nil)
}

func (w MouseEventV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.Constructor")
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

func (w MouseEventV8Wrapper[T]) getModifierState(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.getModifierState")
	return nil, errors.New("MouseEvent.getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEventV8Wrapper[T]) screenX(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.screenX")
	return nil, errors.New("MouseEvent.screenX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEventV8Wrapper[T]) screenY(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.screenY")
	return nil, errors.New("MouseEvent.screenY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEventV8Wrapper[T]) clientX(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.clientX")
	return nil, errors.New("MouseEvent.clientX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEventV8Wrapper[T]) clientY(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.clientY")
	return nil, errors.New("MouseEvent.clientY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEventV8Wrapper[T]) layerX(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.layerX")
	return nil, errors.New("MouseEvent.layerX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEventV8Wrapper[T]) layerY(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.layerY")
	return nil, errors.New("MouseEvent.layerY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w MouseEventV8Wrapper[T]) relatedTarget(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.relatedTarget")
	return nil, errors.New("MouseEvent.relatedTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
