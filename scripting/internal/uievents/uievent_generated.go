// This file is generated. Do not edit.

package uievents

import (
	"errors"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type UIEventV8Wrapper[T any] struct{}

func NewUIEventV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *UIEventV8Wrapper[T] {
	return &UIEventV8Wrapper[T]{}
}

func (wrapper UIEventV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w UIEventV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("view", w.view, nil)
	jsClass.CreatePrototypeAttribute("detail", w.detail, nil)
}

func (w UIEventV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: UIEvent.Constructor")
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

func (w UIEventV8Wrapper[T]) view(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: UIEvent.view")
	return nil, errors.New("UIEvent.view: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w UIEventV8Wrapper[T]) detail(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: UIEvent.detail")
	return nil, errors.New("UIEvent.detail: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
