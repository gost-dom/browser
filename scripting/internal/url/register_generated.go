// This file is generated. Do not edit.

package url

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureWindowRealm[T any](e js.ScriptEngine[T]) {
	InitializeURL(js.CreateClass(e, "URL", "", URLConstructor))
	InitializeURLSearchParams(js.CreateClass(e, "URLSearchParams", "", URLSearchParamsConstructor))
}

func ConfigureDedicatedWorkerGlobalScopeRealm[T any](e js.ScriptEngine[T]) {
	InitializeURL(js.CreateClass(e, "URL", "", URLConstructor))
	InitializeURLSearchParams(js.CreateClass(e, "URLSearchParams", "", URLSearchParamsConstructor))
}
