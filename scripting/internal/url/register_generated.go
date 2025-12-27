// This file is generated. Do not edit.

package url

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "URL", "", NewURL, URLConstructor)
	js.RegisterClass(e, "URLSearchParams", "", NewURLSearchParams, URLSearchParamsConstructor)
}
