// This file is generated. Do not edit.

package streams

import (
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type ReadableStreamBYOBReader[T any] struct{}

func NewReadableStreamBYOBReader[T any](scriptHost js.ScriptEngine[T]) *ReadableStreamBYOBReader[T] {
	return &ReadableStreamBYOBReader[T]{}
}

func (wrapper ReadableStreamBYOBReader[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w ReadableStreamBYOBReader[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("read", w.read)
	jsClass.CreateOperation("releaseLock", w.releaseLock)
}

func ReadableStreamBYOBReaderConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	stream, errArg1 := js.ConsumeArgument(cbCtx, "stream", nil, decodeReadableStream)
	if errArg1 != nil {
		return nil, errArg1
	}
	return CreateReadableStreamBYOBReader(cbCtx, stream)
}

func (w ReadableStreamBYOBReader[T]) read(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStreamBYOBReader.read: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ReadableStreamBYOBReader[T]) releaseLock(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStreamBYOBReader.releaseLock: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
