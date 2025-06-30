// This file is generated. Do not edit.

package fetch

import (
	fetch "github.com/gost-dom/browser/internal/fetch"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Response[T any] struct {
	body *Body[T]
}

func NewResponse[T any](scriptHost js.ScriptEngine[T]) *Response[T] {
	return &Response[T]{NewBody(scriptHost)}
}

func (wrapper Response[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Response[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("clone", w.clone)
	jsClass.CreatePrototypeAttribute("type", w.type_, nil)
	jsClass.CreatePrototypeAttribute("url", w.url, nil)
	jsClass.CreatePrototypeAttribute("redirected", w.redirected, nil)
	jsClass.CreatePrototypeAttribute("status", w.status, nil)
	jsClass.CreatePrototypeAttribute("ok", w.ok, nil)
	jsClass.CreatePrototypeAttribute("statusText", w.statusText, nil)
	jsClass.CreatePrototypeAttribute("headers", w.headers, nil)
	w.body.installPrototype(jsClass)
}

func (w Response[T]) clone(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Response.clone")
	return codec.EncodeCallbackErrorf(cbCtx, "Response.clone: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Response[T]) type_(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Response.type_")
	return codec.EncodeCallbackErrorf(cbCtx, "Response.type_: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Response[T]) url(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Response.url")
	return codec.EncodeCallbackErrorf(cbCtx, "Response.url: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Response[T]) redirected(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Response.redirected")
	return codec.EncodeCallbackErrorf(cbCtx, "Response.redirected: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Response[T]) status(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Response.status")
	instance, err := js.As[*fetch.Response](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Status
	return codec.EncodeInt(cbCtx, result)
}

func (w Response[T]) ok(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Response.ok")
	return codec.EncodeCallbackErrorf(cbCtx, "Response.ok: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Response[T]) statusText(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Response.statusText")
	return codec.EncodeCallbackErrorf(cbCtx, "Response.statusText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Response[T]) headers(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Response.headers")
	return codec.EncodeCallbackErrorf(cbCtx, "Response.headers: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
