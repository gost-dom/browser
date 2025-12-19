// This file is generated. Do not edit.

package xhr

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "FormData", "", NewFormData)
	js.RegisterClass(reg, "XMLHttpRequestEventTarget", "EventTarget", NewXMLHttpRequestEventTarget)
	js.RegisterClass(reg, "XMLHttpRequest", "XMLHttpRequestEventTarget", NewXMLHttpRequest)
}
