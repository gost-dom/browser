package uievents

import (
	"errors"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/uievents"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func decodeInto[T, U any](
	scope js.Scope[T],
	target *U,
	opts js.Object[T],
	key string,
	decoder func(js.Scope[T], js.Value[T]) (U, error),
) error {
	val, err := opts.Get(key)
	if err != nil {
		return err
	}
	if js.IsUndefined(val) {
		return nil
	}
	decoded, err := decoder(scope, val)
	if err == nil {
		*target = decoded
	}
	return err
}

func decodeKeyboardEventInit[T any](
	scope js.Scope[T],
	options js.Object[T],
	init *uievents.KeyboardEventInit,
) error {
	return errors.Join(
		decodeInto(scope, &init.Key, options, "key", codec.DecodeString),
	)
}

func decodePointerEventInit[T any](
	_ js.Scope[T],
	options js.Object[T],
	init *uievents.PointerEventInit,
) error {
	return nil
}

func decodeMouseEventInit[T any](
	_ js.Scope[T],
	options js.Object[T],
	init *uievents.MouseEventInit,
) error {
	return nil
}

func decodeUIEventInit[T any](
	_ js.Scope[T],
	options js.Object[T],
	init *uievents.UIEventInit,
) error {
	return nil
}

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

func (w KeyboardEvent[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	type_ string,
) (js.Value[T], error) {
	return w.CreateInstanceEventInitDict(cbCtx, type_)
}

func (w KeyboardEvent[T]) CreateInstanceEventInitDict(
	cbCtx js.CallbackContext[T],
	type_ string,
	options ...any) (js.Value[T], error) {
	e := uievents.NewKeyboardEvent(type_, uievents.KeyboardEventInit{})
	return codec.EncodeConstructedValue(cbCtx, e)
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
