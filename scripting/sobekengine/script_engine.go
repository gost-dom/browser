package sobekengine

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/cache"
	"github.com/gost-dom/browser/scripting/internal"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type scriptEngine struct {
	initializer *internal.ScriptEngineConfigurer[jsTypeParam]
	cache       cache.Cache
}

func (e *scriptEngine) NewHost(opts html.ScriptEngineOptions) html.ScriptHost {
	res := scriptHost{
		logger:      opts.Logger,
		httpClient:  opts.HttpClient,
		initializer: e.initializer,
		cache:       &e.cache,
		clock:       opts.Clock,
	}
	return &res
}

var defaultEngine *scriptEngine

func DefaultEngine() html.ScriptEngine {
	return defaultEngine
}

func newEngine(configurators ...js.Configurer[jsTypeParam]) *scriptEngine {
	return &scriptEngine{
		initializer: internal.NewScriptEngineConfigurer(configurators...),
	}
}

func init() {
	configurer := internal.CreateWindowsConfigurer[jsTypeParam]()
	defaultEngine = newEngine(configurer)
}
