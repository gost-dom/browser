// This file is generated. Do not edit.

package fetch

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "Headers", "", NewHeaders, HeadersConstructor)
	js.RegisterClass(e, "Request", "", NewRequest, RequestConstructor)
	js.RegisterClass(e, "Response", "", NewResponse, ResponseConstructor)
}
