package internal

import (
	"github.com/gost-dom/browser/scripting/internal/dom"
	"github.com/gost-dom/browser/scripting/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/fileapi"
	"github.com/gost-dom/browser/scripting/internal/html"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/browser/scripting/internal/mathml"
	"github.com/gost-dom/browser/scripting/internal/streams"
	"github.com/gost-dom/browser/scripting/internal/svg"
	"github.com/gost-dom/browser/scripting/internal/uievents"
	"github.com/gost-dom/browser/scripting/internal/url"
	"github.com/gost-dom/browser/scripting/internal/xhr"
)

type ScriptEngineConfigurer[T any] struct {
	initializers []js.Configurer[T]
}

func NewScriptEngineConfigurer[T any](i ...js.Configurer[T]) *ScriptEngineConfigurer[T] {
	return &ScriptEngineConfigurer[T]{i}
}

func (c *ScriptEngineConfigurer[T]) AddConfigurer(configurer js.Configurer[T]) {
	c.initializers = append(c.initializers, configurer)
}

func (c *ScriptEngineConfigurer[T]) AddConfigurerFunc(f func(js.ScriptEngine[T])) {
	c.initializers = append(c.initializers, js.ConfigurerFunc[T](f))
}

func configureNatives[T any](e js.ScriptEngine[T]) {
	e.SetUnhandledPromiseRejectionHandler(
		js.ErrorHandlerFunc[T](handleUnhandledPromiseRejection[T]),
	)
	configureConsole(e)

}

func (c *ScriptEngineConfigurer[T]) Configure(e js.ScriptEngine[T]) {
	for _, i := range c.initializers {
		i.Configure(e)
	}
}

func handleUnhandledPromiseRejection[T any](scope js.Scope[T], err error) {
	dom.HandleJSCallbackError(scope, "promiseRejected", err)
}

func CreateWindowsConfigurer[T any]() *ScriptEngineConfigurer[T] {
	result := NewScriptEngineConfigurer[T]()
	result.AddConfigurerFunc(configureNatives)

	result.AddConfigurerFunc(dom.ConfigureScriptEngine)
	result.AddConfigurerFunc(fetch.ConfigureScriptEngine)
	result.AddConfigurerFunc(html.ConfigureScriptEngine)
	result.AddConfigurerFunc(svg.ConfigureScriptEngine)
	result.AddConfigurerFunc(mathml.ConfigureScriptEngine)
	result.AddConfigurerFunc(xhr.ConfigureScriptEngine)
	result.AddConfigurerFunc(url.Bootstrap)
	result.AddConfigurerFunc(uievents.Bootstrap)
	result.AddConfigurerFunc(streams.ConfigureWindowRealm)
	result.AddConfigurerFunc(fileapi.ConfigureWindowRealm)
	result.AddConfigurerFunc(InstallPolyfills)

	return result
}
