package uievents

import (
	"github.com/gost-dom/browser/internal/uievents"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func decodePointerEventInit[T any](
	_ js.Scope[T], _ js.Object[T], _ *uievents.PointerEventInit,
) error {
	return nil
}

func decodeMouseEventInit[T any](_ js.Scope[T], _ js.Object[T], _ *uievents.MouseEventInit) error {
	return nil
}

func decodeUIEventInit[T any](_ js.Scope[T], _ js.Object[T], _ *uievents.UIEventInit) error {
	return nil
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
