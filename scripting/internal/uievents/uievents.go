package uievents

import (
	"github.com/gost-dom/browser/internal/uievents"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func decodeMouseEventInit[T any](s js.Scope[T], v js.Value[T]) (codec.EventInit, error) {
	return decodeUIEventInit(s, v)
}

func decodePointerEventInit[T any](s js.Scope[T], v js.Value[T]) (codec.EventInit, error) {
	return decodeMouseEventInit(s, v)
}

func CreateUIEvent[T any](cbCtx js.CallbackContext[T], type_ string) (js.Value[T], error) {
	return CreateUIEventEventInitDict(cbCtx, type_)
}

func CreateKeyboardEvent[T any](cbCtx js.CallbackContext[T], type_ string) (js.Value[T], error) {
	return CreateUIEvent(cbCtx, type_)
}

func CreateKeyboardEventEventInitDict[T any](
	cbCtx js.CallbackContext[T],
	type_ string,
	options ...interface{},
) (js.Value[T], error) {
	return CreateUIEventEventInitDict(cbCtx, type_, options...)
}

func CreateMouseEvent[T any](cbCtx js.CallbackContext[T], type_ string) (js.Value[T], error) {
	return CreateUIEvent(cbCtx, type_)
}

func CreateMouseEventEventInitDict[T any](
	cbCtx js.CallbackContext[T],
	type_ string,
	options ...interface{},
) (js.Value[T], error) {
	return CreateUIEventEventInitDict(cbCtx, type_, options...)
}
func CreatePointerEvent[T any](cbCtx js.CallbackContext[T], type_ string) (js.Value[T], error) {
	return CreateUIEvent(cbCtx, type_)
}

func CreatePointerEventEventInitDict[T any](
	cbCtx js.CallbackContext[T],
	type_ string,
	options ...interface{},
) (js.Value[T], error) {
	return CreateUIEventEventInitDict(cbCtx, type_, options...)
}

func CreateUIEventEventInitDict[T any](
	cbCtx js.CallbackContext[T],
	type_ string,
	options ...interface{}) (js.Value[T], error) {
	e := uievents.NewUIEvent(type_)
	return codec.EncodeConstrucedValue(cbCtx, e)
}

func decodeUIEventInit[T any](s js.Scope[T], v js.Value[T]) (codec.EventInit, error) {
	return codec.DecodeEventInit(s, v)
}

func decodeKeyboardEventInit[T any](s js.Scope[T], v js.Value[T]) (codec.EventInit, error) {
	return decodeUIEventInit(s, v)
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

func NewMouseEvent[T any](host js.ScriptEngine[T]) MouseEvent[T] {
	return MouseEvent[T]{*NewUIEvent(host)}
}

func NewPointerEvent[T any](host js.ScriptEngine[T]) PointerEvent[T] {
	return PointerEvent[T]{NewMouseEvent(host)}
}

func NewKeyboardEvent[T any](host js.ScriptEngine[T]) KeyboardEvent[T] {
	return KeyboardEvent[T]{*NewUIEvent(host)}
}
