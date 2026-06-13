package fetch

import (
	"fmt"

	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/internal/promise"
	"github.com/gost-dom/browser/internal/streams"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

// encodeBodyPromise reads the entire [fetch.Body] of the current instance and
// returns a JavaScript Promise that resolves to the body contents, encoded by
// encoder. It is the shared implementation behind the Body consumer methods
// (json, text, bytes, arrayBuffer).
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

// Body_json consumes the body and resolves to the value parsed from its JSON
// contents.
func Body_json[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return encodeBodyPromise(cbCtx, EncodeJSONBytes)
}

// EncodeJSONBytes parses b as JSON and returns the resulting JavaScript value.
func EncodeJSONBytes[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	return scope.JSONParse(string(b))
}

// Body_text consumes the body and resolves to its contents as a UTF-8 string.
func Body_text[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return encodeBodyPromise(cbCtx, encodeBytesAsString)
}

// encodeBytesAsString encodes b as a JavaScript string.
func encodeBytesAsString[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	return scope.NewString(string(b)), nil
}

// Body_bytes consumes the body and resolves to a Uint8Array of its contents.
func Body_bytes[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return encodeBodyPromise(cbCtx, encodeBytesAsUint8Array)
}

// encodeBytesAsUint8Array encodes b as a JavaScript Uint8Array.
func encodeBytesAsUint8Array[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	return scope.NewUint8Array(b), nil
}

// Body_arrayBuffer consumes the body and resolves to an ArrayBuffer of its
// contents. The buffer is obtained from a freshly created Uint8Array view.
func Body_arrayBuffer[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return encodeBodyPromise(cbCtx, encodeBytesAsArrayBuffer)
}

// encodeBytesAsArrayBuffer encodes b as a JavaScript ArrayBuffer, returning the
// buffer that backs a freshly created Uint8Array view of b.
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

// encodeReadableStream wraps body in a JavaScript ReadableStream instance.
func encodeReadableStream[T any](
	cbCtx js.CallbackContext[T],
	body streams.ReadableStream,
) (js.Value[T], error) {
	return cbCtx.Constructor("ReadableStream").NewInstance(body)
}
