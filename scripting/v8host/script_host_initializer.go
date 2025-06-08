package v8host

import (
	"sync"

	"github.com/gost-dom/browser/scripting/internal"
	"github.com/gost-dom/browser/scripting/internal/html"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

type ScriptEngineConfigurer struct {
	initializers []js.Configurator[jsTypeParam]
}

func (c *ScriptEngineConfigurer) AddConfigurator(configurer js.Configurator[jsTypeParam]) {
	c.initializers = append(c.initializers, configurer)
}

func (c *ScriptEngineConfigurer) createHost(config hostOptions) *V8ScriptHost {

	host := &V8ScriptHost{
		mu:       new(sync.Mutex),
		iso:      v8go.NewIsolate(),
		logger:   config.logger,
		globals:  globals{make(map[string]v8Class)},
		contexts: make(map[*v8go.Context]*V8ScriptContext),
	}
	host.iso.SetPromiseRejectedCallback(host.promiseRejected)
	host.windowTemplate = v8go.NewObjectTemplate(host.iso)
	host.iterator = newV8Iterator(host)
	host.windowTemplate.SetInternalFieldCount(1)
	for _, i := range factory.initializers {
		i.Configure(host)
	}
	return host
}

var factory = new(ScriptEngineConfigurer)

func init() {
	var classRegistrations = js.NewClassBuilder[jsTypeParam]()
	internal.Bootstrap(classRegistrations)

	js.AddConfigurator(factory, internal.Configure)
	js.AddConfigurator(factory, html.Initialize)
	js.AddConfigurator(factory, classRegistrations.CreateGlobals)
}
