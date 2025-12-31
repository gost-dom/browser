// This file is generated. Do not edit.

package fetch

import (
	fetch "github.com/gost-dom/browser/internal/fetch"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Body[T any] struct{}

func NewBody[T any](scriptHost js.ScriptEngine[T]) Body[T] {
	return Body[T]{}
}

func (wrapper Body[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Body[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("arrayBuffer", Body_arrayBuffer)
	jsClass.CreateOperation("blob", Body_blob)
	jsClass.CreateOperation("bytes", Body_bytes)
	jsClass.CreateOperation("formData", Body_formData)
	jsClass.CreateOperation("json", Body_json)
	jsClass.CreateOperation("text", Body_text)
	jsClass.CreateAttribute("body", Body_body, nil)
	jsClass.CreateAttribute("bodyUsed", Body_bodyUsed, nil)
}

func Body_arrayBuffer[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.Body_arrayBuffer: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Body_blob[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.Body_blob: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Body_bytes[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.Body_bytes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Body_formData[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.Body_formData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Body_text[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.Body_text: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Body_body[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[fetch.Body](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Body()
	return encodeReadableStream(cbCtx, result)
}

func Body_bodyUsed[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Body.Body_bodyUsed: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
