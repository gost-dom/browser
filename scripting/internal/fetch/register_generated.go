// This file is generated. Do not edit.

package fetch

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureWindowRealm[T any](e js.ScriptEngine[T]) {
	InitializeHeaders(js.CreateClass(e, "Headers", "", HeadersConstructor))
	InitializeRequest(js.CreateClass(e, "Request", "", RequestConstructor))
	InitializeResponse(js.CreateClass(e, "Response", "", ResponseConstructor))
	InitializeWindowOrWorkerGlobalScope(js.MustGetClass(e, "Window"))
}

func ConfigureDedicatedWorkerGlobalScopeRealm[T any](e js.ScriptEngine[T]) {
	InitializeHeaders(js.CreateClass(e, "Headers", "", HeadersConstructor))
	InitializeRequest(js.CreateClass(e, "Request", "", RequestConstructor))
	InitializeResponse(js.CreateClass(e, "Response", "", ResponseConstructor))
	InitializeWindowOrWorkerGlobalScope(js.MustGetClass(e, "WorkerGlobalScope"))
}
