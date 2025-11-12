package internal

import (
	"github.com/gost-dom/browser/scripting/internal/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type ScriptEngineConfigurer[T any] struct {
	initializers []js.Configurator[T]
}

func (c *ScriptEngineConfigurer[T]) AddConfigurator(configurer js.Configurator[T]) {
	c.initializers = append(c.initializers, configurer)
}

func DefaultInitializer[T any]() *ScriptEngineConfigurer[T] {
	factory := &ScriptEngineConfigurer[T]{}
	var classRegistrations = js.NewClassBuilder[T]()
	Bootstrap(classRegistrations)

	js.AddConfigurator(factory, Configure)
	js.AddConfigurator(factory, html.Initialize)
	js.AddConfigurator(factory, classRegistrations.CreateGlobals)
	js.AddConfigurator(factory, InstallPolyfills)
	return factory
}

func (c *ScriptEngineConfigurer[T]) Configure(host js.ScriptEngine[T]) {
	for _, i := range c.initializers {
		i.Configure(host)
	}
}
