// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeHTMLOrSVGElement[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("focus", HTMLOrSVGElement_focus)
	jsClass.CreateOperation("blur", HTMLOrSVGElement_blur)
	jsClass.CreateAttribute("dataset", HTMLOrSVGElement_dataset, nil)
	jsClass.CreateAttribute("nonce", HTMLOrSVGElement_nonce, HTMLOrSVGElement_setNonce)
	jsClass.CreateAttribute("autofocus", HTMLOrSVGElement_autofocus, HTMLOrSVGElement_setAutofocus)
	jsClass.CreateAttribute("tabIndex", HTMLOrSVGElement_tabIndex, HTMLOrSVGElement_setTabIndex)
}

func HTMLOrSVGElement_focus[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	instance.Focus()
	return nil, nil
}

func HTMLOrSVGElement_blur[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Blur()
	return nil, nil
}

func HTMLOrSVGElement_dataset[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Dataset()
	return codec.EncodeEntity(cbCtx, result)
}

func HTMLOrSVGElement_nonce[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Nonce()
	return codec.EncodeString(cbCtx, result)
}

func HTMLOrSVGElement_setNonce[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetNonce(val)
	return nil, nil
}

func HTMLOrSVGElement_autofocus[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Autofocus()
	return codec.EncodeBoolean(cbCtx, result)
}

func HTMLOrSVGElement_setAutofocus[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeBoolean)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetAutofocus(val)
	return nil, nil
}

func HTMLOrSVGElement_tabIndex[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.TabIndex()
	return codec.EncodeInt(cbCtx, result)
}

func HTMLOrSVGElement_setTabIndex[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeInt)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTabIndex(val)
	return nil, nil
}
