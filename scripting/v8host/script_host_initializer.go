package v8host

import (
	"github.com/gost-dom/browser/scripting/internal"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type ScriptEngineConfigurer struct{}

func (c ScriptEngineConfigurer) Register(configurer js.Configurator[jsTypeParam]) {
	initializers = append(initializers, configurer)
}

func init() { js.Register(ScriptEngineConfigurer{}, internal.Configure) }
