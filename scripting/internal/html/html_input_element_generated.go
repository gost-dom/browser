// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
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
	jsClass.CreateOperation("checkValidity", w.checkValidity)
	jsClass.CreateAttribute("name", w.name, w.setName)
	jsClass.CreateAttribute("type", w.type_, w.setType)
	jsClass.CreateAttribute("value", w.value, w.setValue)
}

func HTMLInputElementConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLInputElement[T]) checkValidity(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CheckValidity()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w HTMLInputElement[T]) name(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Name()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLInputElement[T]) setName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetName(val)
	return nil, nil
}

func (w HTMLInputElement[T]) type_(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLInputElement[T]) setType(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetType(val)
	return nil, nil
}

func (w HTMLInputElement[T]) value(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLInputElement[T]) setValue(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}
