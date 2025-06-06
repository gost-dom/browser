// This file is generated. Do not edit.

package v8host

import (
	"errors"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (wrapper mouseEventV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w mouseEventV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("getModifierState", w.getModifierState)
	jsClass.CreatePrototypeAttribute("screenX", w.screenX, nil)
	jsClass.CreatePrototypeAttribute("screenY", w.screenY, nil)
	jsClass.CreatePrototypeAttribute("clientX", w.clientX, nil)
	jsClass.CreatePrototypeAttribute("clientY", w.clientY, nil)
	jsClass.CreatePrototypeAttribute("layerX", w.layerX, nil)
	jsClass.CreatePrototypeAttribute("layerY", w.layerY, nil)
	jsClass.CreatePrototypeAttribute("relatedTarget", w.relatedTarget, nil)
}

func (w mouseEventV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.Constructor")
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	eventInitDict, found, errArg := consumeOptionalArg(cbCtx, "eventInitDict", w.decodeMouseEventInit)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceEventInitDict(cbCtx, type_, eventInitDict)
	}
	return w.CreateInstance(cbCtx, type_)
}

func (w mouseEventV8Wrapper[T]) getModifierState(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.getModifierState")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper[T]) screenX(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.screenX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.screenX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper[T]) screenY(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.screenY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.screenY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper[T]) clientX(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.clientX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.clientX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper[T]) clientY(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.clientY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.clientY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper[T]) layerX(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.layerX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.layerX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper[T]) layerY(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.layerY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.layerY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper[T]) relatedTarget(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.relatedTarget")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.relatedTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

type uIEventV8Wrapper[T any] struct{}

func newUIEventV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *uIEventV8Wrapper[T] {
	return &uIEventV8Wrapper[T]{}
}

func (wrapper uIEventV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w uIEventV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("view", w.view, nil)
	jsClass.CreatePrototypeAttribute("detail", w.detail, nil)
}

func (w uIEventV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: UIEvent.Constructor")
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	eventInitDict, found, errArg := consumeOptionalArg(cbCtx, "eventInitDict", w.decodeUIEventInit)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceEventInitDict(cbCtx, type_, eventInitDict)
	}
	return w.CreateInstance(cbCtx, type_)
}

func (w uIEventV8Wrapper[T]) view(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: UIEvent.view")
	return cbCtx.ReturnWithError(errors.New("UIEvent.view: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w uIEventV8Wrapper[T]) detail(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: UIEvent.detail")
	return cbCtx.ReturnWithError(errors.New("UIEvent.detail: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
