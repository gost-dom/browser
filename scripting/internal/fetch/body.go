package fetch

import (
	"fmt"

	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/internal/promise"
	"github.com/gost-dom/browser/internal/streams"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func Body_json[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[fetch.Body](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodePromise(cbCtx, promise.ReadAll(instance), EncodeJSONBytes)
}

func EncodeJSONBytes[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	return scope.JSONParse(string(b))
}

// Body_text consumes the body and resolves to its contents as a UTF-8 string.
func Body_text[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[fetch.Body](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodePromise(cbCtx, promise.ReadAll(instance), encodeBytesAsString)
}

func encodeBytesAsString[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	return scope.NewString(string(b)), nil
}

// Body_bytes consumes the body and resolves to a Uint8Array of its contents.
func Body_bytes[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[fetch.Body](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodePromise(cbCtx, promise.ReadAll(instance), encodeBytesAsUint8Array)
}

func encodeBytesAsUint8Array[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	return scope.NewUint8Array(b), nil
}

// Body_arrayBuffer consumes the body and resolves to an ArrayBuffer of its
// contents. The buffer is obtained from a freshly created Uint8Array view.
func Body_arrayBuffer[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[fetch.Body](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodePromise(cbCtx, promise.ReadAll(instance), encodeBytesAsArrayBuffer)
}

func encodeBytesAsArrayBuffer[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	arr := scope.NewUint8Array(b)
	obj, ok := arr.AsObject()
	if !ok {
		// A freshly created Uint8Array is always an object; if that ever fails
		// we must not return the Uint8Array, as callers expect an ArrayBuffer.
		return nil, fmt.Errorf("gost-dom/fetch: arrayBuffer: Uint8Array is not an object")
	}
	buf, err := obj.Get("buffer")
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func encodeReadableStream[T any](
	cbCtx js.CallbackContext[T],
	body streams.ReadableStream,
) (js.Value[T], error) {
	return cbCtx.Constructor("ReadableStream").NewInstance(body)
}
