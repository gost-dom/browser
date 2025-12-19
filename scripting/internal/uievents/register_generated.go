// This file is generated. Do not edit.

package uievents

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "UIEvent", "Event", NewUIEvent)
	js.RegisterClass(e, "KeyboardEvent", "UIEvent", NewKeyboardEvent)
	js.RegisterClass(e, "MouseEvent", "UIEvent", NewMouseEvent)
	js.RegisterClass(e, "PointerEvent", "MouseEvent", NewPointerEvent)
}
