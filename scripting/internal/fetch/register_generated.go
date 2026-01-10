// This file is generated. Do not edit.

package fetch

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureWindowRealm[T any](e js.ScriptEngine[T]) {
	InitializeHeaders(js.CreateClass(e, "Headers", "", HeadersConstructor))
	InitializeRequest(js.CreateClass(e, "Request", "", RequestConstructor))
	InitializeResponse(js.CreateClass(e, "Response", "", ResponseConstructor))
	window, ok := e.Class("Window")
	if !ok {
		panic("gost-dom/fetch: Window: class not registered")
	}
	InitializeWindowOrWorkerGlobalScope(window)
}

func ConfigureDedicatedWorkerGlobalScopeRealm[T any](e js.ScriptEngine[T]) {
	InitializeHeaders(js.CreateClass(e, "Headers", "", HeadersConstructor))
	InitializeRequest(js.CreateClass(e, "Request", "", RequestConstructor))
	InitializeResponse(js.CreateClass(e, "Response", "", ResponseConstructor))
	workerGlobalScope, ok := e.Class("WorkerGlobalScope")
	if !ok {
		panic("gost-dom/fetch: WorkerGlobalScope: class not registered")
	}
	InitializeWindowOrWorkerGlobalScope(workerGlobalScope)
}
