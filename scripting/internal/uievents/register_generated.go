// This file is generated. Do not edit.

package uievents

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureWindowRealm[T any](e js.ScriptEngine[T]) {
	InitializeUIEvent(js.CreateClass(e, "UIEvent", "Event", UIEventConstructor))
	InitializeKeyboardEvent(js.CreateClass(e, "KeyboardEvent", "UIEvent", KeyboardEventConstructor))
	InitializeMouseEvent(js.CreateClass(e, "MouseEvent", "UIEvent", MouseEventConstructor))
	InitializePointerEvent(js.CreateClass(e, "PointerEvent", "MouseEvent", PointerEventConstructor))
}

func ConfigureDedicatedWorkerGlobalScopeRealm[T any](e js.ScriptEngine[T]) {}
