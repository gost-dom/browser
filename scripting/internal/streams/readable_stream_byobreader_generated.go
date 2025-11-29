// This file is generated. Do not edit.

package streams

import (
	log "github.com/gost-dom/browser/internal/log"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"log/slog"
)

type ReadableStreamBYOBReader[T any] struct{}

func NewReadableStreamBYOBReader[T any](scriptHost js.ScriptEngine[T]) *ReadableStreamBYOBReader[T] {
	return &ReadableStreamBYOBReader[T]{}
}

func (wrapper ReadableStreamBYOBReader[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w ReadableStreamBYOBReader[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("read", w.read)
	jsClass.CreatePrototypeMethod("releaseLock", w.releaseLock)
}

func (w ReadableStreamBYOBReader[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "ReadableStreamBYOBReader"), slog.String("Method", "Constructor"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	stream, errArg1 := js.ConsumeArgument(cbCtx, "stream", nil, w.decodeReadableStream)
	if errArg1 != nil {
		return nil, errArg1
	}
	return w.CreateInstance(cbCtx, stream)
}

func (w ReadableStreamBYOBReader[T]) read(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "ReadableStreamBYOBReader"), slog.String("Method", "read"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStreamBYOBReader.read: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ReadableStreamBYOBReader[T]) releaseLock(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "ReadableStreamBYOBReader"), slog.String("Method", "releaseLock"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "ReadableStreamBYOBReader.releaseLock: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
