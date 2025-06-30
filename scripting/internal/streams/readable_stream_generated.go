// This file is generated. Do not edit.

package streams

import (
	"errors"
	streams "github.com/gost-dom/browser/internal/streams"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type ReadableStream[T any] struct{}

func NewReadableStream[T any](scriptHost js.ScriptEngine[T]) *ReadableStream[T] {
	return &ReadableStream[T]{}
}

func (wrapper ReadableStream[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w ReadableStream[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("cancel", w.cancel)
	jsClass.CreatePrototypeMethod("getReader", w.getReader)
	jsClass.CreatePrototypeMethod("pipeThrough", w.pipeThrough)
	jsClass.CreatePrototypeMethod("pipeTo", w.pipeTo)
	jsClass.CreatePrototypeMethod("tee", w.tee)
	jsClass.CreatePrototypeAttribute("locked", w.locked, nil)
}

func (w ReadableStream[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ReadableStream.Constructor")
	underlyingSource, errArg1 := js.ConsumeArgument(cbCtx, "underlyingSource", codec.ZeroValue, w.decodeObject)
	strategy, errArg2 := js.ConsumeArgument(cbCtx, "strategy", nil, w.decodeQueuingStrategy)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	return w.CreateInstance(cbCtx, underlyingSource, strategy...)
}

func (w ReadableStream[T]) cancel(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ReadableStream.cancel")
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStream.cancel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ReadableStream[T]) getReader(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ReadableStream.getReader")
	instance, errInst := js.As[streams.ReadableStream](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	options, errArg1 := js.ConsumeArgument(cbCtx, "options", nil, w.decodeReadableStreamGetReaderOptions)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetReader(options...)
	return w.toReadableStreamReader(cbCtx, result)
}

func (w ReadableStream[T]) pipeThrough(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ReadableStream.pipeThrough")
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStream.pipeThrough: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ReadableStream[T]) pipeTo(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ReadableStream.pipeTo")
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStream.pipeTo: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ReadableStream[T]) tee(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ReadableStream.tee")
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStream.tee: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ReadableStream[T]) locked(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ReadableStream.locked")
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStream.locked: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
