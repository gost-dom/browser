// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type htmlInputElementV8Wrapper[T any] struct {
	handleReffedObject[html.HTMLInputElement, T]
}

func newHTMLInputElementV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *htmlInputElementV8Wrapper[T] {
	return &htmlInputElementV8Wrapper[T]{newHandleReffedObject[html.HTMLInputElement, T](scriptHost)}
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
	return w.toBoolean(cbCtx, result)
}

func (w htmlInputElementV8Wrapper[T]) type_(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLInputElement.type_")
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Type()
	return w.toString_(cbCtx, result)
}

func (w htmlInputElementV8Wrapper[T]) setType(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLInputElement.setType")
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetType(val)
	return cbCtx.ReturnWithValue(nil)
}
