// This file is generated. Do not edit.

package xhr

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "FormData", "", NewFormDataV8Wrapper)
	js.RegisterClass(reg, "XMLHttpRequest", "XMLHttpRequestEventTarget", NewXMLHttpRequestV8Wrapper)
	js.RegisterClass(reg, "XMLHttpRequestEventTarget", "EventTarget", NewXMLHttpRequestEventTargetV8Wrapper)
}
