package uievents

import (
	"github.com/gost-dom/browser/internal/uievents"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w UIEventV8Wrapper[T]) decodeMouseEventInit(
	cbCtx js.CallbackContext[T],
	v js.Value[T],
) (codec.EventInit, error) {
	return w.decodeUIEventInit(cbCtx, v)
}

func (w UIEventV8Wrapper[T]) decodePointerEventInit(
	cbCtx js.CallbackContext[T],
	v js.Value[T],
) (codec.EventInit, error) {
	return w.decodeMouseEventInit(cbCtx, v)
}

func (w UIEventV8Wrapper[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	type_ string,
) (js.Value[T], error) {
	return w.CreateInstanceEventInitDict(cbCtx, type_)
}

func (w UIEventV8Wrapper[T]) CreateInstanceEventInitDict(
	cbCtx js.CallbackContext[T],
	type_ string,
	options ...interface{}) (js.Value[T], error) {
	e := uievents.NewUIEvent(type_)
	return codec.EncodeConstrucedValue(cbCtx, e)
}

func (w UIEventV8Wrapper[T]) decodeUIEventInit(
	cbCtx js.CallbackContext[T],
	v js.Value[T],
) (codec.EventInit, error) {
	return codec.DecodeEventInit(cbCtx, v)
}

type MouseEventV8Wrapper[T any] struct {
	UIEventV8Wrapper[T]
}

type PointerEventV8Wrapper[T any] struct {
	MouseEventV8Wrapper[T]
}

func NewMouseEventV8Wrapper[T any](host js.ScriptEngine[T]) MouseEventV8Wrapper[T] {
	return MouseEventV8Wrapper[T]{*NewUIEventV8Wrapper(host)}
}

func NewPointerEventV8Wrapper[T any](host js.ScriptEngine[T]) PointerEventV8Wrapper[T] {
	return PointerEventV8Wrapper[T]{NewMouseEventV8Wrapper(host)}
}
