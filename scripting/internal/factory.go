package internal

import (
	"github.com/gost-dom/browser/scripting/internal/dom"
	"github.com/gost-dom/browser/scripting/internal/fetch"
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

func DefaultInitializer[T any](e js.ScriptEngine[T]) {
	e.SetUnhandledPromiseRejectionHandler(
		js.ErrorHandlerFunc[T](handleUnhandledPromiseRejection[T]),
	)
	dom.ConfigureScriptEngine(e)
	fetch.ConfigureScriptEngine(e)
	configureConsole(e)

	html.ConfigureScriptEngine(e)
	svg.ConfigureScriptEngine(e)
	mathml.ConfigureScriptEngine(e)
	xhr.ConfigureScriptEngine(e)
	url.Bootstrap(e)
	uievents.Bootstrap(e)
	streams.Bootstrap(e)

	js.RegisterClass(e, "File", "", dom.NewEvent)

	InstallPolyfills(e)
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
	result.AddConfigurerFunc(DefaultInitializer)
	return result
}
