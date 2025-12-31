// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeHTMLInputElement[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("checkValidity", HTMLInputElement_checkValidity)
	jsClass.CreateAttribute("name", HTMLInputElement_name, HTMLInputElement_setName)
	jsClass.CreateAttribute("type", HTMLInputElement_type, HTMLInputElement_setType)
	jsClass.CreateAttribute("value", HTMLInputElement_value, HTMLInputElement_setValue)
}

func HTMLInputElementConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func HTMLInputElement_checkValidity[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CheckValidity()
	return codec.EncodeBoolean(cbCtx, result)
}

func HTMLInputElement_name[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Name()
	return codec.EncodeString(cbCtx, result)
}

func HTMLInputElement_setName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetName(val)
	return nil, nil
}

func HTMLInputElement_type[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type()
	return codec.EncodeString(cbCtx, result)
}

func HTMLInputElement_setType[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetType(val)
	return nil, nil
}

func HTMLInputElement_value[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return codec.EncodeString(cbCtx, result)
}

func HTMLInputElement_setValue[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}
