// This file is generated. Do not edit.

package uievents

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "UIEvent", "Event", UIEvent[T]{}.Initialize, UIEventConstructor)
	js.RegisterClass(e, "KeyboardEvent", "UIEvent", KeyboardEvent[T]{}.Initialize, KeyboardEventConstructor)
	js.RegisterClass(e, "MouseEvent", "UIEvent", MouseEvent[T]{}.Initialize, MouseEventConstructor)
	js.RegisterClass(e, "PointerEvent", "MouseEvent", PointerEvent[T]{}.Initialize, PointerEventConstructor)
}
