// This file is generated. Do not edit.

package uievents

import (
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type UIEvent[T any] struct{}

func NewUIEvent[T any](scriptHost js.ScriptEngine[T]) *UIEvent[T] {
	return &UIEvent[T]{}
}

func (wrapper UIEvent[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w UIEvent[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("view", w.view, nil)
	jsClass.CreatePrototypeAttribute("detail", w.detail, nil)
}

func (w UIEvent[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: UIEvent.Constructor")
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	eventInitDict, found, errArg := js.ConsumeOptionalArg(cbCtx, "eventInitDict", w.decodeUIEventInit)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceEventInitDict(cbCtx, type_, eventInitDict)
	}
	return w.CreateInstance(cbCtx, type_)
}

func (w UIEvent[T]) view(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: UIEvent.view")
	return codec.EncodeCallbackErrorf(cbCtx, "UIEvent.view: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w UIEvent[T]) detail(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: UIEvent.detail")
	return codec.EncodeCallbackErrorf(cbCtx, "UIEvent.detail: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
