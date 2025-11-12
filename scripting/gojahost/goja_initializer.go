package gojahost

import (
	"github.com/gost-dom/browser/scripting/internal"
	"github.com/gost-dom/browser/scripting/internal/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type ScriptEngineConfigurer struct {
	initializers []js.Configurator[jsTypeParam]
}

func (c *ScriptEngineConfigurer) AddConfigurator(configurer js.Configurator[jsTypeParam]) {
	c.initializers = append(c.initializers, configurer)
}

var factory = new(ScriptEngineConfigurer)

func init() {
	var classRegistrations = js.NewClassBuilder[jsTypeParam]()
	internal.Bootstrap(classRegistrations)

	js.AddConfigurator(factory, internal.Configure)
	js.AddConfigurator(factory, html.Initialize)
	js.AddConfigurator(factory, classRegistrations.CreateGlobals)
	js.AddConfigurator(factory, internal.InstallPolyfills)
}
