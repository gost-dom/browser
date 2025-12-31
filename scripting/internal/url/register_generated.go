// This file is generated. Do not edit.

package url

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "URL", "", InitializeURL, URLConstructor)
	js.RegisterClass(e, "URLSearchParams", "", InitializeURLSearchParams, URLSearchParamsConstructor)
}
