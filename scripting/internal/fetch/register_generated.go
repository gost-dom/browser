// This file is generated. Do not edit.

package fetch

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "Body", "", NewBody)
	js.RegisterClass(reg, "Headers", "", NewHeaders)
	js.RegisterClass(reg, "Request", "", NewRequest)
	js.RegisterClass(reg, "Response", "", NewResponse)
}
