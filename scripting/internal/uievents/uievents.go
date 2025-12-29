package uievents

import (
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/uievents"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w UIEvent[T]) decodeMouseEventInit(s js.Scope[T], v js.Value[T]) (codec.EventInit, error) {
	return w.decodeUIEventInit(s, v)
}

func (w UIEvent[T]) decodePointerEventInit(s js.Scope[T], v js.Value[T]) (codec.EventInit, error) {
	return w.decodeMouseEventInit(s, v)
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
	return codec.EncodeConstructedValue(cbCtx, e)
}

func (w UIEvent[T]) decodeUIEventInit(s js.Scope[T], v js.Value[T]) (codec.EventInit, error) {
	return codec.DecodeEventInit(s, v)
}

func (w UIEvent[T]) decodeKeyboardEventInit(s js.Scope[T], v js.Value[T]) (codec.EventInit, error) {
	return w.decodeUIEventInit(s, v)
}

type MouseEvent[T any] struct {
	UIEvent[T]
}

type PointerEvent[T any] struct {
	MouseEvent[T]
}

type KeyboardEvent[T any] struct {
	UIEvent[T]
}

func (e KeyboardEvent[T]) key(cbctx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[*event.Event](cbctx.Instance())
	if err != nil {
		return nil, err
	}
	eventInit, ok := instance.Data.(uievents.KeyboardEventInit)
	if !ok {
		return nil, cbctx.NewTypeError("Object is not a KeyboardEvent")
	}
	return codec.EncodeString(cbctx, eventInit.Key)
}

func NewMouseEvent[T any](host js.ScriptEngine[T]) MouseEvent[T] {
	return MouseEvent[T]{*NewUIEvent(host)}
}

func NewPointerEvent[T any](host js.ScriptEngine[T]) PointerEvent[T] {
	return PointerEvent[T]{NewMouseEvent(host)}
}

func NewKeyboardEvent[T any](host js.ScriptEngine[T]) KeyboardEvent[T] {
	return KeyboardEvent[T]{*NewUIEvent(host)}
}
