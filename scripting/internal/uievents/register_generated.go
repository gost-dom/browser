// This file is generated. Do not edit.

package uievents

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "UIEvent", "Event", NewUIEvent, UIEventConstructor)
	js.RegisterClass(e, "KeyboardEvent", "UIEvent", NewKeyboardEvent, KeyboardEventConstructor)
	js.RegisterClass(e, "MouseEvent", "UIEvent", NewMouseEvent, MouseEventConstructor)
	js.RegisterClass(e, "PointerEvent", "MouseEvent", NewPointerEvent, PointerEventConstructor)
}
