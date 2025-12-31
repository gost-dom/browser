// This file is generated. Do not edit.

package streams

import (
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	streams "github.com/gost-dom/browser/internal/streams"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeReadableStream[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("cancel", ReadableStream_cancel)
	jsClass.CreateOperation("getReader", ReadableStream_getReader)
	jsClass.CreateOperation("pipeThrough", ReadableStream_pipeThrough)
	jsClass.CreateOperation("pipeTo", ReadableStream_pipeTo)
	jsClass.CreateOperation("tee", ReadableStream_tee)
	jsClass.CreateAttribute("locked", ReadableStream_locked, nil)
}

func ReadableStreamConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	underlyingSource, errArg1 := js.ConsumeArgument(cbCtx, "underlyingSource", codec.ZeroValue, decodeObject)
	strategy, errArg2 := js.ConsumeArgument(cbCtx, "strategy", nil, decodeQueuingStrategy)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	return CreateReadableStream(cbCtx, underlyingSource, strategy...)
}

func ReadableStream_cancel[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStream.ReadableStream_cancel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func ReadableStream_getReader[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[streams.ReadableStream](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	options, errArg1 := js.ConsumeArgument(cbCtx, "options", nil, decodeReadableStreamGetReaderOptions)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetReader(options...)
	return encodeReadableStreamReader(cbCtx, result)
}

func ReadableStream_pipeThrough[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStream.ReadableStream_pipeThrough: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func ReadableStream_pipeTo[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStream.ReadableStream_pipeTo: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func ReadableStream_tee[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStream.ReadableStream_tee: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func ReadableStream_locked[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStream.ReadableStream_locked: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
