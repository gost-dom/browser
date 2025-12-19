// This file is generated. Do not edit.

package fetch

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "Body", "", NewBody)
	js.RegisterClass(e, "Headers", "", NewHeaders)
	js.RegisterClass(e, "Request", "", NewRequest)
	js.RegisterClass(e, "Response", "", NewResponse)
}
