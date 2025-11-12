package v8host

import "github.com/gost-dom/browser/html"

type v8ScriptEngine struct{}

func (e v8ScriptEngine) NewHost(options html.ScriptEngineOptions) html.ScriptHost {
	return New(
		WithLogger(options.Logger),
		WithHTTPClient(options.HttpClient),
	)
}

func NewEngine() html.ScriptEngine {
	return v8ScriptEngine{}
}
