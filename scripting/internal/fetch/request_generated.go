// This file is generated. Do not edit.

package fetch

import (
	fetch "github.com/gost-dom/browser/internal/fetch"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Request[T any] struct {
	body *Body[T]
}

func NewRequest[T any](scriptHost js.ScriptEngine[T]) *Request[T] {
	return &Request[T]{NewBody(scriptHost)}
}

func (wrapper Request[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Request[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("clone", Request_clone)
	jsClass.CreateAttribute("method", Request_method, nil)
	jsClass.CreateAttribute("url", Request_url, nil)
	jsClass.CreateAttribute("headers", Request_headers, nil)
	jsClass.CreateAttribute("destination", Request_destination, nil)
	jsClass.CreateAttribute("referrer", Request_referrer, nil)
	jsClass.CreateAttribute("referrerPolicy", Request_referrerPolicy, nil)
	jsClass.CreateAttribute("mode", Request_mode, nil)
	jsClass.CreateAttribute("credentials", Request_credentials, nil)
	jsClass.CreateAttribute("cache", Request_cache, nil)
	jsClass.CreateAttribute("redirect", Request_redirect, nil)
	jsClass.CreateAttribute("integrity", Request_integrity, nil)
	jsClass.CreateAttribute("keepalive", Request_keepalive, nil)
	jsClass.CreateAttribute("isReloadNavigation", Request_isReloadNavigation, nil)
	jsClass.CreateAttribute("isHistoryNavigation", Request_isHistoryNavigation, nil)
	jsClass.CreateAttribute("signal", Request_signal, nil)
	jsClass.CreateAttribute("duplex", Request_duplex, nil)
	w.body.installPrototype(jsClass)
}

func RequestConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	input, errArg1 := js.ConsumeArgument(cbCtx, "input", nil, decodeRequestInfo)
	init, errArg2 := js.ConsumeArgument(cbCtx, "init", nil, decodeRequestInit)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	return CreateRequest(cbCtx, input, init...)
}

func Request_clone[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_clone: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_method[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_method: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_url[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*fetch.Request](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.URL()
	return codec.EncodeString(cbCtx, result)
}

func Request_headers[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*fetch.Request](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := &instance.Headers
	return encodeHeaders(cbCtx, result)
}

func Request_destination[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_destination: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_referrer[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_referrer: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_referrerPolicy[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_referrerPolicy: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_mode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_mode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_credentials[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_credentials: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_cache[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_cache: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_redirect[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_redirect: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_integrity[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_integrity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_keepalive[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_keepalive: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_isReloadNavigation[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_isReloadNavigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_isHistoryNavigation[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_isHistoryNavigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_signal[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_signal: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Request_duplex[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.Request_duplex: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
