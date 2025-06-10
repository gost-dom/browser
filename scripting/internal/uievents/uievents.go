package uievents

import (
	"github.com/gost-dom/browser/internal/uievents"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w UIEvent[T]) decodeMouseEventInit(
	cbCtx js.CallbackContext[T],
	v js.Value[T],
) (codec.EventInit, error) {
	return w.decodeUIEventInit(cbCtx, v)
}

func (w UIEvent[T]) decodePointerEventInit(
	cbCtx js.CallbackContext[T],
	v js.Value[T],
) (codec.EventInit, error) {
	return w.decodeMouseEventInit(cbCtx, v)
}

func (w UIEvent[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	type_ string,
) (js.Value[T], error) {
	return w.CreateInstanceEventInitDict(cbCtx, type_)
}

func (w UIEvent[T]) CreateInstanceEventInitDict(
	cbCtx js.CallbackContext[T],
	type_ string,
	options ...interface{}) (js.Value[T], error) {
	e := uievents.NewUIEvent(type_)
	return codec.EncodeConstrucedValue(cbCtx, e)
}

func (w UIEvent[T]) decodeUIEventInit(
	cbCtx js.CallbackContext[T],
	v js.Value[T],
) (codec.EventInit, error) {
	return codec.DecodeEventInit(cbCtx, v)
}

type MouseEvent[T any] struct {
	UIEvent[T]
}

type PointerEvent[T any] struct {
	MouseEvent[T]
}

func NewMouseEvent[T any](host js.ScriptEngine[T]) MouseEvent[T] {
	return MouseEvent[T]{*NewUIEvent(host)}
}

func NewPointerEvent[T any](host js.ScriptEngine[T]) PointerEvent[T] {
	return PointerEvent[T]{NewMouseEvent(host)}
}
