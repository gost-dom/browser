package internal

import (
	"github.com/gost-dom/browser/scripting/internal/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type ScriptEngineConfigurer[T any] struct {
	initializers []js.Configurer[T]
}

func NewScriptEngineConfigurer[T any](i []js.Configurer[T]) *ScriptEngineConfigurer[T] {
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
	Configure(e)
	html.Initialize(e)
	Bootstrap(e)
	InstallPolyfills(e)
}

func (c *ScriptEngineConfigurer[T]) Configure(host js.ScriptEngine[T]) {
	for _, i := range c.initializers {
		i.Configure(host)
	}
}
