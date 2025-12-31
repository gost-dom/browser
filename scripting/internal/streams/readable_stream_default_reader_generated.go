// This file is generated. Do not edit.

package streams

import (
	streams "github.com/gost-dom/browser/internal/streams"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type ReadableStreamDefaultReader[T any] struct{}

func NewReadableStreamDefaultReader[T any](scriptHost js.ScriptEngine[T]) ReadableStreamDefaultReader[T] {
	return ReadableStreamDefaultReader[T]{}
}

func (wrapper ReadableStreamDefaultReader[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w ReadableStreamDefaultReader[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("read", ReadableStreamDefaultReader_read)
	jsClass.CreateOperation("releaseLock", ReadableStreamDefaultReader_releaseLock)
}

func ReadableStreamDefaultReaderConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	stream, errArg1 := js.ConsumeArgument(cbCtx, "stream", nil, decodeReadableStream)
	if errArg1 != nil {
		return nil, errArg1
	}
	return CreateReadableStreamDefaultReader(cbCtx, stream)
}

func ReadableStreamDefaultReader_read[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[streams.ReadableStreamDefaultReader](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Read()
	return encodePromiseReadableStreamReadResult(cbCtx, result)
}

func ReadableStreamDefaultReader_releaseLock[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStreamDefaultReader.ReadableStreamDefaultReader_releaseLock: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
