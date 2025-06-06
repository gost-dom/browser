// This file is generated. Do not edit.

package url

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "URL", "", NewURLV8Wrapper)
	js.RegisterClass(reg, "URLSearchParams", "", NewURLSearchParamsV8Wrapper)
}
