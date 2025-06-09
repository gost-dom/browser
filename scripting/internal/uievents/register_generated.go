// This file is generated. Do not edit.

package uievents

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "PointerEvent", "MouseEvent", NewPointerEventV8Wrapper)
	js.RegisterClass(reg, "MouseEvent", "UIEvent", NewMouseEventV8Wrapper)
	js.RegisterClass(reg, "UIEvent", "Event", NewUIEventV8Wrapper)
}
