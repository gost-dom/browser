package v8host

import (
	"runtime"
	"sync"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/v8go"
)

// MAX_POOL_SIZE sets a limit to the number of script hosts that will be pooled
// for reuse. By default Go will run as many tests in parallel as you have CPU
// cores, so there shouldn't be a reason for a larger pool, but this provides a
// bit of flexibility.
var MAX_POOL_SIZE = runtime.NumCPU() * 2

type v8ScriptEngine struct {
	m        sync.Mutex
	isolates []*V8ScriptHost
}

type engineOwnedHost struct {
	*V8ScriptHost
	engine *v8ScriptEngine
}

func (h engineOwnedHost) Close() {
	h.Close()
	if !h.engine.add(h.V8ScriptHost) {
		h.Dispose()
	}
}

func (e *v8ScriptEngine) NewHost(options html.ScriptEngineOptions) html.ScriptHost {
	return e.newHost(options)
}

func (e *v8ScriptEngine) newHost(options html.ScriptEngineOptions) *V8ScriptHost {
	host, found := e.tryGet()
	if found {
		host.disposed = false
		host.logger = options.Logger
		host.httpClient = options.HttpClient
	} else {
		host = factory.createHost(hostOptions{
			logger:     options.Logger,
			httpClient: options.HttpClient,
		})
	}
	host.inspectorClient = v8go.NewInspectorClient(consoleAPIMessageFunc(host.consoleAPIMessage))
	host.inspector = v8go.NewInspector(host.iso, host.inspectorClient)
	return host
}

func NewEngine() html.ScriptEngine {
	return &v8ScriptEngine{}
}

func (pool *v8ScriptEngine) tryGet() (iso *V8ScriptHost, found bool) {
	pool.m.Lock()
	defer pool.m.Unlock()

	l := len(pool.isolates)
	if l == 0 {
		return nil, false
	}

	iso = pool.isolates[l-1]
	pool.isolates = pool.isolates[0 : l-1]
	return iso, true
}

func (pool *v8ScriptEngine) add(iso *V8ScriptHost) bool {
	pool.m.Lock()
	defer pool.m.Unlock()

	if len(pool.isolates) >= MAX_POOL_SIZE {
		return false
	} else {
		pool.isolates = append(pool.isolates, iso)
		return true
	}
}

var DefaultEngine v8ScriptEngine
