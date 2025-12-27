// This file is generated. Do not edit.

package xhr

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "FormData", "", NewFormData, FormDataConstructor)
	js.RegisterClass(e, "XMLHttpRequestEventTarget", "EventTarget", NewXMLHttpRequestEventTarget, XMLHttpRequestEventTargetConstructor)
	js.RegisterClass(e, "XMLHttpRequest", "XMLHttpRequestEventTarget", NewXMLHttpRequest, XMLHttpRequestConstructor)
}
