// This file is generated. Do not edit.

package uievents

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "UIEvent", "Event", InitializeUIEvent, UIEventConstructor)
	js.RegisterClass(e, "KeyboardEvent", "UIEvent", InitializeKeyboardEvent, KeyboardEventConstructor)
	js.RegisterClass(e, "MouseEvent", "UIEvent", InitializeMouseEvent, MouseEventConstructor)
	js.RegisterClass(e, "PointerEvent", "MouseEvent", InitializePointerEvent, PointerEventConstructor)
}
