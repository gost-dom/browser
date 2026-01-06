// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeDocument[T any](jsClass js.Class[T]) {
	jsClass.CreateAttribute("location", Document_location, nil, js.LegacyUnforgeable())
	jsClass.CreateAttribute("body", Document_body, Document_setBody)
	jsClass.CreateAttribute("head", Document_head, nil)
}

func Document_location[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Location()
	return codec.EncodeEntity(cbCtx, result)
}

func Document_body[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Body()
	return codec.EncodeEntity(cbCtx, result)
}

func Document_setBody[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLDocument](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeHTMLElement)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	return nil, instance.SetBody(val)
}

func Document_head[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Head()
	return codec.EncodeEntity(cbCtx, result)
}
