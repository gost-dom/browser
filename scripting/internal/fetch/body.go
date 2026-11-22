package fetch

import (
	"fmt"

	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/internal/promise"
	"github.com/gost-dom/browser/internal/streams"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

// encodeBodyPromise is the shared implementation behind the Body consumer
// methods (json, text, bytes, arrayBuffer): it reads the entire [fetch.Body] of
// the current instance and resolves a Promise to the contents encoded by encoder.
func encodeBodyPromise[T any](
	cbCtx js.CallbackContext[T],
	encoder codec.Encoder[T, []byte],
) (js.Value[T], error) {
	instance, err := js.As[fetch.Body](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodePromise(cbCtx, promise.ReadAll(instance), encoder)
}

func Body_json[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return encodeBodyPromise(cbCtx, EncodeJSONBytes)
}

func EncodeJSONBytes[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	return scope.JSONParse(string(b))
}

func Body_text[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return encodeBodyPromise(cbCtx, encodeBytesAsString)
}

func encodeBytesAsString[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	return scope.NewString(string(b)), nil
}

func Body_bytes[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return encodeBodyPromise(cbCtx, encodeBytesAsUint8Array)
}

func encodeBytesAsUint8Array[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	return scope.NewUint8Array(b), nil
}

func Body_arrayBuffer[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return encodeBodyPromise(cbCtx, encodeBytesAsArrayBuffer)
}

func encodeBytesAsArrayBuffer[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	arr := scope.NewUint8Array(b)
	obj, ok := arr.AsObject()
	if !ok {
		// A freshly created Uint8Array is always an object; if that ever fails
		// we must not return the Uint8Array, as callers expect an ArrayBuffer.
		return nil, fmt.Errorf("gost-dom/fetch: arrayBuffer: Uint8Array is not an object")
	}
	return obj.Get("buffer")
}

func encodeReadableStream[T any](
	cbCtx js.CallbackContext[T],
	body streams.ReadableStream,
) (js.Value[T], error) {
	return cbCtx.Constructor("ReadableStream").NewInstance(body)
}
