// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type htmlInputElementV8Wrapper[T any] struct{}

func newHTMLInputElementV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *htmlInputElementV8Wrapper[T] {
	return &htmlInputElementV8Wrapper[T]{}
}

func (wrapper htmlInputElementV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w htmlInputElementV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("checkValidity", w.checkValidity)
	jsClass.CreatePrototypeAttribute("type", w.type_, w.setType)
}

func (w htmlInputElementV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLInputElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlInputElementV8Wrapper[T]) checkValidity(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLInputElement.checkValidity")
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.CheckValidity()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w htmlInputElementV8Wrapper[T]) type_(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLInputElement.type_")
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Type()
	return codec.EncodeString(cbCtx, result)
}

func (w htmlInputElementV8Wrapper[T]) setType(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLInputElement.setType")
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetType(val)
	return cbCtx.ReturnWithValue(nil)
}
