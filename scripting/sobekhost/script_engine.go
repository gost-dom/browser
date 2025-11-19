package sobekhost

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal"
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

func init() {
	defaultEngine = &scriptEngine{
		internal.DefaultInitializer[jsTypeParam](),
	}
}
