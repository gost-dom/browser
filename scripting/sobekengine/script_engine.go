package sobekengine

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type scriptEngine struct {
	initializer *internal.ScriptEngineConfigurer[jsTypeParam]
}

func (e scriptEngine) NewHost(opts html.ScriptEngineOptions) html.ScriptHost {
	res := scriptHost{
		Logger:      opts.Logger,
		HttpClient:  opts.HttpClient,
		initializer: e.initializer,
	}
	return &res
}

var defaultEngine *scriptEngine

func DefaultEngine() html.ScriptEngine {
	return defaultEngine
}

func newEngine(configurators ...js.Configurer[jsTypeParam]) *scriptEngine {
	return &scriptEngine{internal.NewScriptEngineConfigurer(configurators...)}
}

func init() {
	configurer := internal.CreateWindowsConfigurer[jsTypeParam]()
	defaultEngine = newEngine(configurer)
}
