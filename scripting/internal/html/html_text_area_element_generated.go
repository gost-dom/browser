// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeHTMLTextAreaElement[T any](jsClass js.Class[T]) {
	jsClass.CreateAttribute("value", HTMLTextAreaElement_value, HTMLTextAreaElement_setValue)
}

func HTMLTextAreaElement_value[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLTextAreaElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return codec.EncodeString(cbCtx, result)
}

func HTMLTextAreaElement_setValue[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLTextAreaElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}
