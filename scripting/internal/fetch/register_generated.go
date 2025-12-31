// This file is generated. Do not edit.

package fetch

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "Headers", "", InitializeHeaders, HeadersConstructor)
	js.RegisterClass(e, "Request", "", InitializeRequest, RequestConstructor)
	js.RegisterClass(e, "Response", "", InitializeResponse, ResponseConstructor)
}
