package v8engine

import (
	"errors"
	"fmt"
	"maps"
	"net/http"
	"slices"
	"sync"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

type globals struct {
	namedGlobals map[string]v8Class
}

type hostOptions struct {
	httpClient *http.Client
	logger     log.Logger
}

type HostOption func(o *hostOptions)

func WithLogger(logger log.Logger) HostOption {
	return func(o *hostOptions) {
		if logger != nil {
			o.logger = logger
		}
	}
}

func WithHTTPClient(client *http.Client) HostOption {
	return func(o *hostOptions) {
		if client != nil {
			o.httpClient = client
		}
	}
}

type V8ScriptHost struct {
	logger          log.Logger
	mu              *sync.Mutex
	iso             *v8go.Isolate
	inspector       *v8go.Inspector
	inspectorClient *v8go.InspectorClient
	windowTemplate  *v8go.ObjectTemplate
	globals         globals
	scripts         [][2]string
	httpClient      *http.Client
	contexts        map[*v8go.Context]*V8ScriptContext
	disposed        bool
	iterator        v8Iterator
}

// consoleAPIMessageFunc represents a function that can receive javascript
// console messages and implements the [v8go.consoleAPIMessageFunc] interface.
//
// This type is a simple solution to avoid exporting the consoleAPIMessage
// function.
type consoleAPIMessageFunc func(message v8go.ConsoleAPIMessage)

func (f consoleAPIMessageFunc) ConsoleAPIMessage(message v8go.ConsoleAPIMessage) {
	f(message)
}

func (host *V8ScriptHost) consoleAPIMessage(message v8go.ConsoleAPIMessage) {
	switch message.ErrorLevel {
	case v8go.ErrorLevelDebug:
		host.Logger().Debug(message.Message)
	case v8go.ErrorLevelInfo:
	case v8go.ErrorLevelLog:
		host.Logger().Info(message.Message)
	case v8go.ErrorLevelWarning:
		host.Logger().Warn(message.Message)
	case v8go.ErrorLevelError:
		host.Logger().Error(message.Message)
	}
}

func (host *V8ScriptHost) deleteContext(ctx *V8ScriptContext) {
	host.mu.Lock()
	defer host.mu.Unlock()
	delete(host.contexts, ctx.v8ctx)
}

func (host *V8ScriptHost) promiseRejected(msg v8go.PromiseRejectMessage) {
	if msg.Event != v8go.PromiseRejectWithNoHandler {
		return
	}
	ctx := host.mustGetContext(msg.Context)
	var err error
	if msg.Value == nil {
		err = errors.New("unhandled promise rejection: no error value")
	} else {
		if exc, excErr := msg.Value.AsException(); excErr == nil {
			err = fmt.Errorf("unhandled promise rejection: %w", exc)
		} else {
			err = fmt.Errorf("unhandled promise rejection: %v", msg.Value)
		}
	}

	js.HandleJSCallbackError(newV8Scope(ctx), "Promise", err)
}

func (host *V8ScriptHost) Logger() log.Logger {
	if host.logger == nil {
		return log.Default()
	}
	return host.logger
}

// New returns a new initialized V8ScriptHost
//
// Deprecated: Obtain V8ScriptHost instances from a script engine which will
// handle caching of unused isolates
func New(opts ...HostOption) *V8ScriptHost {
	config := hostOptions{}
	for _, opt := range opts {
		opt(&config)
	}

	return defaultEngine.newHost(html.ScriptEngineOptions{
		Logger:     config.logger,
		HttpClient: config.httpClient,
	})
}

// Close informs that client code is done using this script host. The host will
// be placed into a pool;
func (host *V8ScriptHost) Close() {
	host.setDisposed()

	undisposedContexts := host.undisposedContexts()
	undisposedCount := len(undisposedContexts)

	if undisposedCount > 0 {
		host.Logger().Warn(
			"Closing script host with undisposed contexts",
			"count",
			undisposedCount,
		)
		for _, ctx := range undisposedContexts {
			ctx.Close()
		}
	}

	host.inspectorClient.Dispose()
	host.inspector.Dispose()
}

func (host *V8ScriptHost) Dispose() {
	host.iso.Dispose()
}

func (host *V8ScriptHost) undisposedContexts() []*V8ScriptContext {
	host.mu.Lock()
	defer host.mu.Unlock()
	return slices.Collect(maps.Values(host.contexts))
}

func (host *V8ScriptHost) setDisposed() {
	host.mu.Lock()
	defer host.mu.Unlock()
	host.disposed = true
}

// NewContext creates a new script context using w as the global window object.
func (host *V8ScriptHost) NewContext(w html.Window) html.ScriptContext {
	host.assertUndisposed()
	v8ctx := v8go.NewContext(host.iso, host.windowTemplate)
	context := &V8ScriptContext{
		host:     host,
		clock:    clock.New(clock.WithLogger(w.Logger())),
		v8ctx:    v8ctx,
		window:   w,
		v8nodes:  make(map[entity.ObjectId]jsValue),
		resolver: moduleResolver{host: host},
	}
	if _, err := context.runScript("Object.setPrototypeOf(globalThis, globalThis.Window.prototype)"); err != nil {
		panic(err)
	}
	host.addContext(context)
	if err := context.initializeGlobals(); err != nil {
		panic(err)
	}
	host.inspector.ContextCreated(context.v8ctx)
	global := newV8Object(context, context.v8ctx.Global())
	context.addDisposer(global.(js.Disposable))
	context.cacheEntity(global, w)

	return context
}

func (host *V8ScriptHost) assertUndisposed() {
	host.mu.Lock()
	defer host.mu.Unlock()

	if host.disposed {
		panic("gost-dom/v8engine: V8ScriptHost.NewContext: script host disposed")
	}
}

func (host *V8ScriptHost) CreateClass(
	name string,
	extends js.Class[jsTypeParam],
	callback js.FunctionCallback[jsTypeParam],
) js.Class[jsTypeParam] {
	ft := wrapV8Callback(host, callback.WithLog(name, "Constructor"))
	result := newV8Class(host, name, ft)
	result.inst.SetInternalFieldCount(1)
	if extends != nil {
		ft.Inherit(extends.(v8Class).ft)
	}
	host.windowTemplate.Set(name, ft)
	host.globals.namedGlobals[name] = result
	return result
}

// CreateGlobalObject implements [js/ScriptEngine.CreateGlobalObject]
func (h *V8ScriptHost) CreateGlobalObject(name string) js.GlobalObject[jsTypeParam] {
	tmpl := v8go.NewObjectTemplate(h.iso)
	result := newV8GlobalObject(h, tmpl)
	h.windowTemplate.Set(name, tmpl)
	return result
}

func (host *V8ScriptHost) CreateFunction(
	name string,
	callback js.FunctionCallback[jsTypeParam],
) {
	ft := wrapV8Callback(host, callback.WithLog("", name))
	host.windowTemplate.Set(name, ft)
}

func (host *V8ScriptHost) RunScript(script, src string) {
	host.scripts = append(host.scripts, [2]string{script, src})
}
