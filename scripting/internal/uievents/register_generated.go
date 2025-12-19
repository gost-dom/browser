// This file is generated. Do not edit.

package uievents

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "UIEvent", "Event", NewUIEvent)
	js.RegisterClass(reg, "KeyboardEvent", "UIEvent", NewKeyboardEvent)
	js.RegisterClass(reg, "MouseEvent", "UIEvent", NewMouseEvent)
	js.RegisterClass(reg, "PointerEvent", "MouseEvent", NewPointerEvent)
}
