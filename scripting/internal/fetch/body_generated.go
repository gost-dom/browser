// This file is generated. Do not edit.

package fetch

import (
	fetch "github.com/gost-dom/browser/internal/fetch"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Body[T any] struct{}

func NewBody[T any](scriptHost js.ScriptEngine[T]) *Body[T] {
	return &Body[T]{}
}

func (wrapper Body[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Body[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("arrayBuffer", w.arrayBuffer)
	jsClass.CreatePrototypeMethod("blob", w.blob)
	jsClass.CreatePrototypeMethod("bytes", w.bytes)
	jsClass.CreatePrototypeMethod("formData", w.formData)
	jsClass.CreatePrototypeMethod("json", w.json)
	jsClass.CreatePrototypeMethod("text", w.text)
	jsClass.CreatePrototypeAttribute("body", w.body, nil)
	jsClass.CreatePrototypeAttribute("bodyUsed", w.bodyUsed, nil)
}

func (w Body[T]) arrayBuffer(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.arrayBuffer: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Body[T]) blob(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.blob: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Body[T]) bytes(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.bytes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Body[T]) formData(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.formData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Body[T]) text(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.text: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Body[T]) body(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[fetch.Body](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Body()
	return w.toReadableStream(cbCtx, result)
}

func (w Body[T]) bodyUsed(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.bodyUsed: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
