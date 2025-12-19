package internal

import (
	"github.com/gost-dom/browser/scripting/internal/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type ScriptEngineConfigurer[T any] struct {
	initializers []js.Configurator[T]
}

func NewScriptEngineConfigurer[T any](i []js.Configurator[T]) *ScriptEngineConfigurer[T] {
	return &ScriptEngineConfigurer[T]{i}
}

func (c *ScriptEngineConfigurer[T]) AddConfigurator(configurer js.Configurator[T]) {
	c.initializers = append(c.initializers, configurer)
}

func DefaultInitializer[T any](e js.ScriptEngine[T]) {
	var classRegistrations = js.NewClassBuilder[T]()
	Bootstrap(classRegistrations)

	Configure(e)
	html.Initialize(e)
	classRegistrations.CreateGlobals(e)
	InstallPolyfills(e)
}

func (c *ScriptEngineConfigurer[T]) Configure(host js.ScriptEngine[T]) {
	for _, i := range c.initializers {
		i.Configure(host)
	}
}
