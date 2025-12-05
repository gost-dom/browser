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
	jsClass.CreatePrototypeMethod("clone", w.clone)
	jsClass.CreatePrototypeAttribute("method", w.method, nil)
	jsClass.CreatePrototypeAttribute("url", w.url, nil)
	jsClass.CreatePrototypeAttribute("headers", w.headers, nil)
	jsClass.CreatePrototypeAttribute("destination", w.destination, nil)
	jsClass.CreatePrototypeAttribute("referrer", w.referrer, nil)
	jsClass.CreatePrototypeAttribute("referrerPolicy", w.referrerPolicy, nil)
	jsClass.CreatePrototypeAttribute("mode", w.mode, nil)
	jsClass.CreatePrototypeAttribute("credentials", w.credentials, nil)
	jsClass.CreatePrototypeAttribute("cache", w.cache, nil)
	jsClass.CreatePrototypeAttribute("redirect", w.redirect, nil)
	jsClass.CreatePrototypeAttribute("integrity", w.integrity, nil)
	jsClass.CreatePrototypeAttribute("keepalive", w.keepalive, nil)
	jsClass.CreatePrototypeAttribute("isReloadNavigation", w.isReloadNavigation, nil)
	jsClass.CreatePrototypeAttribute("isHistoryNavigation", w.isHistoryNavigation, nil)
	jsClass.CreatePrototypeAttribute("signal", w.signal, nil)
	jsClass.CreatePrototypeAttribute("duplex", w.duplex, nil)
	w.body.installPrototype(jsClass)
}

func (w Request[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	input, errArg1 := js.ConsumeArgument(cbCtx, "input", nil, w.decodeRequestInfo)
	init, errArg2 := js.ConsumeArgument(cbCtx, "init", nil, w.decodeRequestInit)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	return w.CreateInstance(cbCtx, input, init...)
}

func (w Request[T]) clone(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.clone: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) method(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.method: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) url(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*fetch.Request](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.URL()
	return codec.EncodeString(cbCtx, result)
}

func (w Request[T]) headers(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*fetch.Request](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := &instance.Headers
	return w.toHeaders(cbCtx, result)
}

func (w Request[T]) destination(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.destination: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) referrer(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.referrer: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) referrerPolicy(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.referrerPolicy: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) mode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.mode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) credentials(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.credentials: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) cache(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.cache: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) redirect(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.redirect: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) integrity(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.integrity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) keepalive(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.keepalive: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) isReloadNavigation(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.isReloadNavigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) isHistoryNavigation(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.isHistoryNavigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) signal(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.signal: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Request[T]) duplex(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Request.duplex: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
