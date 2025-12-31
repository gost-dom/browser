// This file is generated. Do not edit.

package fetch

import (
	fetch "github.com/gost-dom/browser/internal/fetch"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Response[T any] struct {
	body Body[T]
}

func NewResponse[T any](scriptHost js.ScriptEngine[T]) Response[T] {
	return Response[T]{NewBody(scriptHost)}
}

func (wrapper Response[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Response[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("clone", Response_clone)
	jsClass.CreateAttribute("type", Response_type, nil)
	jsClass.CreateAttribute("url", Response_url, nil)
	jsClass.CreateAttribute("redirected", Response_redirected, nil)
	jsClass.CreateAttribute("status", Response_status, nil)
	jsClass.CreateAttribute("ok", Response_ok, nil)
	jsClass.CreateAttribute("statusText", Response_statusText, nil)
	jsClass.CreateAttribute("headers", Response_headers, nil)
	w.body.installPrototype(jsClass)
}

func Response_clone[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Response.Response_clone: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Response_type[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Response.Response_type: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Response_url[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Response.Response_url: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Response_redirected[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Response.Response_redirected: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Response_status[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*fetch.Response](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Status
	return codec.EncodeInt(cbCtx, result)
}

func Response_ok[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Response.Response_ok: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Response_statusText[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Response.Response_statusText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Response_headers[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*fetch.Response](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := &instance.Headers
	return encodeHeaders(cbCtx, result)
}
