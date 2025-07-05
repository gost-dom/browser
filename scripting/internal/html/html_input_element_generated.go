// This file is generated. Do not edit.

package html

import (
	"errors"

	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLInputElement[T any] struct{}

func NewHTMLInputElement[T any](scriptHost js.ScriptEngine[T]) *HTMLInputElement[T] {
	return &HTMLInputElement[T]{}
}

func (wrapper HTMLInputElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLInputElement[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("checkValidity", w.checkValidity)
	jsClass.CreatePrototypeAttribute("name", w.name, w.setName)
	jsClass.CreatePrototypeAttribute("type", w.type_, w.setType)
	jsClass.CreatePrototypeAttribute("value", w.value, w.setValue)
}

func (w HTMLInputElement[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: HTMLInputElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLInputElement[T]) checkValidity(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: HTMLInputElement.checkValidity")
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CheckValidity()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w HTMLInputElement[T]) name(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: HTMLInputElement.name")
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Name()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLInputElement[T]) setName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: HTMLInputElement.setName")
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetName(val)
	return nil, nil
}

func (w HTMLInputElement[T]) type_(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: HTMLInputElement.type_")
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLInputElement[T]) setType(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: HTMLInputElement.setType")
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetType(val)
	return nil, nil
}

func (w HTMLInputElement[T]) value(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: HTMLInputElement.value")
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLInputElement[T]) setValue(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: HTMLInputElement.setValue")
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}
