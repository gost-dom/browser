// This file is generated. Do not edit.

package xhr

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	InitializeFormData(js.CreateClass(e, "FormData", "", FormDataConstructor))
	InitializeXMLHttpRequestEventTarget(js.CreateClass(e, "XMLHttpRequestEventTarget", "EventTarget", nil))
	InitializeXMLHttpRequest(js.CreateClass(e, "XMLHttpRequest", "XMLHttpRequestEventTarget", XMLHttpRequestConstructor))
}
