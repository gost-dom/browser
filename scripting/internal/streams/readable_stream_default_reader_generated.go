// This file is generated. Do not edit.

package streams

import (
	streams "github.com/gost-dom/browser/internal/streams"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type ReadableStreamDefaultReader[T any] struct{}

func NewReadableStreamDefaultReader[T any](scriptHost js.ScriptEngine[T]) *ReadableStreamDefaultReader[T] {
	return &ReadableStreamDefaultReader[T]{}
}

func (wrapper ReadableStreamDefaultReader[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w ReadableStreamDefaultReader[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("read", w.read)
	jsClass.CreatePrototypeMethod("releaseLock", w.releaseLock)
}

func (w ReadableStreamDefaultReader[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ReadableStreamDefaultReader.Constructor")
	stream, errArg1 := js.ConsumeArgument(cbCtx, "stream", nil, w.decodeReadableStream)
	if errArg1 != nil {
		return nil, errArg1
	}
	return w.CreateInstance(cbCtx, stream)
}

func (w ReadableStreamDefaultReader[T]) read(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ReadableStreamDefaultReader.read")
	instance, err := js.As[streams.ReadableStreamDefaultReader](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Read()
	return w.toPromiseReadableStreamReadResult(cbCtx, result)
}

func (w ReadableStreamDefaultReader[T]) releaseLock(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ReadableStreamDefaultReader.releaseLock")
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStreamDefaultReader.releaseLock: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
