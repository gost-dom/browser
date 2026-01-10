// This file is generated. Do not edit.

package internal

import (
	dom "github.com/gost-dom/browser/scripting/internal/dom"
	fetch "github.com/gost-dom/browser/scripting/internal/fetch"
	html "github.com/gost-dom/browser/scripting/internal/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
	streams "github.com/gost-dom/browser/scripting/internal/streams"
	url "github.com/gost-dom/browser/scripting/internal/url"
	xhr "github.com/gost-dom/browser/scripting/internal/xhr"
)

func ConfigureWindowRealm[T any](e js.ScriptEngine[T]) {
	dom.ConfigureWindowRealm(e)
	html.ConfigureWindowRealm(e)
	fetch.ConfigureWindowRealm(e)
	streams.ConfigureWindowRealm(e)
	url.ConfigureWindowRealm(e)
	xhr.ConfigureWindowRealm(e)
}

func ConfigureDedicatedWorkerGlobalScopeRealm[T any](e js.ScriptEngine[T]) {
	dom.ConfigureDedicatedWorkerGlobalScopeRealm(e)
	html.ConfigureDedicatedWorkerGlobalScopeRealm(e)
	fetch.ConfigureDedicatedWorkerGlobalScopeRealm(e)
	streams.ConfigureDedicatedWorkerGlobalScopeRealm(e)
	url.ConfigureDedicatedWorkerGlobalScopeRealm(e)
	xhr.ConfigureDedicatedWorkerGlobalScopeRealm(e)
}
