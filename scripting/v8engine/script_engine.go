package v8engine

import (
	"runtime"
	"sync"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal"
	"github.com/gost-dom/v8go"
)

// MAX_POOL_SIZE sets a limit to the number of script hosts that will be pooled
// for reuse. By default Go will run as many tests in parallel as you have CPU
// cores, so there shouldn't be a reason for a larger pool, but this provides a
// bit of flexibility.
var MAX_POOL_SIZE = runtime.NumCPU() * 2

var _ html.ScriptEngine = &v8ScriptEngine{}

type v8ScriptEngine struct {
	m          sync.Mutex
	isolates   []*V8ScriptHost
	configurer *internal.ScriptEngineConfigurer[jsTypeParam]
}

// engineOwnedHost embeds a V8ScriptHost but adds a new Close() method that
// allows the host to be returned to the pool.
type engineOwnedHost struct {
	*V8ScriptHost

	engine *v8ScriptEngine
}

func (h engineOwnedHost) Close() {
	h.V8ScriptHost.Close()
	if !h.engine.add(h.V8ScriptHost) {
		h.Dispose()
	}
}

func (e *v8ScriptEngine) NewHost(options html.ScriptEngineOptions) html.ScriptHost {
	return engineOwnedHost{e.newHost(options), e}
}

func (e *v8ScriptEngine) newHost(options html.ScriptEngineOptions) *V8ScriptHost {
	host, found := e.tryGet()
	if found {
		host.disposed = false
		host.logger = options.Logger
		host.httpClient = options.HttpClient
	} else {
		host = e.createHost(options)
	}
	host.inspectorClient = v8go.NewInspectorClient(consoleAPIMessageFunc(host.consoleAPIMessage))
	host.inspector = v8go.NewInspector(host.iso, host.inspectorClient)
	return host
}

func (e *v8ScriptEngine) createHost(config html.ScriptEngineOptions) *V8ScriptHost {
	host := &V8ScriptHost{
		mu:         new(sync.Mutex),
		iso:        v8go.NewIsolate(),
		httpClient: config.HttpClient,
		logger:     config.Logger,
		globals:    globals{make(map[string]v8Class)},
		contexts:   make(map[*v8go.Context]*V8ScriptContext),
	}
	host.iso.SetPromiseRejectedCallback(host.promiseRejected)
	host.windowTemplate = v8go.NewObjectTemplate(host.iso)
	host.iterator = newV8Iterator(host)
	host.windowTemplate.SetInternalFieldCount(1)

	e.configurer.Configure(host)
	return host
}

func (e *v8ScriptEngine) tryGet() (iso *V8ScriptHost, found bool) {
	e.m.Lock()
	defer e.m.Unlock()

	l := len(e.isolates)
	if l == 0 {
		return nil, false
	}

	iso = e.isolates[l-1]
	e.isolates = e.isolates[0 : l-1]
	return iso, true
}

func (e *v8ScriptEngine) add(iso *V8ScriptHost) bool {
	e.m.Lock()
	defer e.m.Unlock()

	if len(e.isolates) >= MAX_POOL_SIZE {
		return false
	} else {
		e.isolates = append(e.isolates, iso)
		return true
	}
}

var defaultEngine *v8ScriptEngine

func DefaultEngine() html.ScriptEngine {
	return defaultEngine
}

func init() {
	defaultEngine = new(v8ScriptEngine)
	defaultEngine.configurer = internal.DefaultInitializer[jsTypeParam]()
}
