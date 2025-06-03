package v8host

import (
	"errors"
	"fmt"
	"maps"
	"runtime"
	"slices"
	"sync"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

type jsScriptEngineInitializer = func(js.ScriptEngine[jsTypeParam])

// MAX_POOL_SIZE sets a limit to the number of script hosts that will be pooled
// for reuse. By default Go will run as many tests in parallel as you have CPU
// cores, so there shouldn't be a reason for a larger pool, but this provides a
// bit of flexibility.
var MAX_POOL_SIZE = runtime.NumCPU() * 2

// A pool of unused V8ScriptHost instances. When a host is disposed calling
// [V8ScriptHost.Close], it will be added to the pool.
type scriptHostPool struct {
	m        sync.Mutex
	isolates []*V8ScriptHost
}

func (pool *scriptHostPool) releaseAll() []*V8ScriptHost {
	pool.m.Lock()
	defer pool.m.Unlock()

	res := pool.isolates
	pool.isolates = nil
	return res
}

func (pool *scriptHostPool) add(iso *V8ScriptHost) bool {
	pool.m.Lock()
	defer pool.m.Unlock()

	if len(pool.isolates) >= MAX_POOL_SIZE {
		return false
	} else {
		pool.isolates = append(pool.isolates, iso)
		return true
	}
}

func (pool *scriptHostPool) tryGet() (iso *V8ScriptHost, found bool) {
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

var pool = &scriptHostPool{}

type globalInstall struct {
	name        string
	constructor *v8go.FunctionTemplate
}

type globals struct {
	namedGlobals map[string]v8Class
}

type hostOptions struct {
	logger log.Logger
}

type HostOption func(o *hostOptions)

func WithLogger(logger log.Logger) HostOption { return func(o *hostOptions) { o.logger = logger } }

type V8ScriptHost struct {
	logger          log.Logger
	mu              *sync.Mutex
	iso             *v8go.Isolate
	inspector       *v8go.Inspector
	inspectorClient *v8go.InspectorClient
	windowTemplate  *v8go.ObjectTemplate
	globals         globals
	contexts        map[*v8go.Context]*V8ScriptContext
	disposed        bool
	iterator        v8Iterator
}

type jsConstructorFactory1 = func(*V8ScriptHost) *v8go.FunctionTemplate
type jsConstructorFactory = func(*V8ScriptHost, *v8go.FunctionTemplate) *v8go.FunctionTemplate

type class struct {
	globalIdentifier string
	constructor      jsConstructorFactory
	subClasses       []class
}

// createGlobals returns an ordered list of constructors to be created in global
// scope. They must be installed in "order", as base classes must be installed
// before subclasses
func createGlobals(host *V8ScriptHost) []globalInstall {
	result := make([]globalInstall, 0)
	var iter func(class classSpec) *v8go.FunctionTemplate
	uniqueNames := make(map[string]*v8go.FunctionTemplate)
	iter = func(class classSpec) *v8go.FunctionTemplate {
		if constructor, found := uniqueNames[class.name]; found {
			return constructor
		}
		var superClassConstructor *v8go.FunctionTemplate
		if class.superClassName != "" {
			superClassSpec, found := classes[class.superClassName]
			if !found {
				panic(fmt.Sprintf(
					"Missing super class spec. Class: %s. Super: %s",
					class.name, class.superClassName,
				))
			}
			superClassConstructor = iter(superClassSpec)
		}
		constructor := class.factory(host, superClassConstructor)
		uniqueNames[class.name] = constructor
		result = append(result, globalInstall{class.name, constructor})
		return constructor
	}
	for _, class := range classes {
		iter(class)
	}
	host.iterator = newV8Iterator(host)
	return result
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
		log.Debug(host.logger, message.Message)
	case v8go.ErrorLevelInfo:
	case v8go.ErrorLevelLog:
		log.Info(host.logger, message.Message)
	case v8go.ErrorLevelWarning:
		log.Warn(host.logger, message.Message)
	case v8go.ErrorLevelError:
		log.Error(host.logger, message.Message)
	}
}

type classSpec struct {
	name           string
	superClassName string
	factory        jsConstructorFactory
}

var classes map[string]classSpec = make(map[string]classSpec)
var initializers []js.Configurator[jsTypeParam]

func registerJSClass(
	className string,
	superClassName string,
	constructorFactory jsConstructorFactory1,
) {
	spec := classSpec{
		className, superClassName, func(host *V8ScriptHost, extends *v8go.FunctionTemplate) *v8go.FunctionTemplate {
			res := constructorFactory(host)
			if extends != nil {
				res.Inherit(extends)
			}
			return res
		},
	}
	if _, ok := classes[className]; ok {
		panic("Same class added twice: " + className)
	}
	if superClassName == "" {
		classes[className] = spec
		return
	}
	parent, parentFound := classes[superClassName]
	for parentFound {
		if parent.superClassName == className {
			panic("Recursive class parents" + className)
		}
		parent, parentFound = classes[parent.superClassName]
	}
	classes[className] = spec
}

func init() {
	registerJSClass("File", "", createCustomEvent)
	registerJSClass("CustomEvent", "Event", createCustomEvent)
	registerJSClass("Location", "", createLocationPrototype)
	// registerJSClass("EventTarget", "", createEventTarget)
	registerJSClass("XMLHttpRequestEventTarget", "EventTarget", createIllegalConstructor)

	registerJSClass("HTMLDocument", "Document", createHTMLDocumentPrototype)
	registerJSClass("DocumentFragment", "Node", createDocumentFragmentPrototype)
	registerJSClass("ShadowRoot", "DocumentFragment", createShadowRootPrototype)
	registerJSClass("Attr", "Node", createAttr)

	// registerJSClass("DOMParser", "", createDOMParserPrototype)

	for _, cls := range scripting.HtmlElements {
		if _, found := classes[cls]; !found {
			registerJSClass(cls, "HTMLElement", createIllegalConstructor)
		}
	}
}

func createHostInstance(config hostOptions) *V8ScriptHost {
	var host *V8ScriptHost
	res, hostReused := pool.tryGet()

	if hostReused {
		host = res
		host.disposed = false
		host.logger = config.logger
	} else {
		host = &V8ScriptHost{
			mu:     new(sync.Mutex),
			iso:    v8go.NewIsolate(),
			logger: config.logger,
		}
	}

	host.inspectorClient = v8go.NewInspectorClient(consoleAPIMessageFunc(host.consoleAPIMessage))
	host.inspector = v8go.NewInspector(host.iso, host.inspectorClient)

	if !hostReused {
		host.iso.SetPromiseRejectedCallback(host.promiseRejected)
		globalInstalls := createGlobals(host)
		host.globals = globals{make(map[string]v8Class)}
		var window *v8go.FunctionTemplate
		for _, globalInstall := range globalInstalls {
			host.globals.namedGlobals[globalInstall.name] = newV8Class(
				host,
				globalInstall.constructor,
			)
			if globalInstall.name == "Window" {
				window = globalInstall.constructor
			}
		}
		if window == nil {
			panic("No window was created. " + constants.BUG_ISSUE_URL)
		}
		host.windowTemplate = window.InstanceTemplate()
		host.contexts = make(map[*v8go.Context]*V8ScriptContext)
		installGlobals(window, host, globalInstalls)
		installEventLoopGlobals(host, host.windowTemplate)
		for _, i := range initializers {
			i.Configure(host)
		}
	}
	return host
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

	log.Error(host.logger, "Rejected promise", log.ErrAttr(err))
	UnhandledError(&v8Scope{ctx}, err)
}

func (host *V8ScriptHost) Logger() log.Logger {
	if host.logger == nil {
		return log.Default()
	}
	return host.logger
}

func New(opts ...HostOption) *V8ScriptHost {
	config := hostOptions{}
	for _, opt := range opts {
		opt(&config)
	}

	res := createHostInstance(config)
	res.logger = config.logger
	return res
}

// Close informs that client code is done using this script host. The host will
// be placed into a pool;
func (host *V8ScriptHost) Close() {
	host.setDisposed()

	undisposedContexts := host.undisposedContexts()
	undisposedCount := len(undisposedContexts)

	if undisposedCount > 0 {
		log.Warn(
			host.logger,
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

	if !pool.add(host) {
		host.iso.Close()
	}
}

// Dispose all pooled isolates
func Shutdown() {
	for _, host := range pool.releaseAll() {
		host.iso.Dispose()
	}
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
// Calling with a nil value for w is allowed, but not supported; and any attempt
// to access the DOM will result in a runtime error.
func (host *V8ScriptHost) NewContext(w html.Window) html.ScriptContext {
	v8ctx := v8go.NewContext(host.iso, host.windowTemplate)
	// TODO: The possibility to use nil is primarily for testing support
	context := &V8ScriptContext{
		host:    host,
		clock:   clock.New(),
		v8ctx:   v8ctx,
		window:  w,
		v8nodes: make(map[entity.ObjectId]jsValue),
	}
	context.global = newV8Object(context, v8ctx.Global())
	host.addContext(context)
	host.inspector.ContextCreated(context.v8ctx)
	if w != nil {
		global := newV8Object(context, context.v8ctx.Global())
		global.SetNativeValue(w)
		context.addDisposer(global.(js.Disposable))
		context.cacheEntity(global, w)
	}
	err := installPolyfills(context)
	if err != nil {
		// TODO: Handle
		panic(
			fmt.Sprintf(
				"Error installing polyfills. Should not be possible on a passing build of Gost-DOM.\n  Please file an issue if this is a release version of Gost-DOM: %s\n  Error: %v",
				constants.BUG_ISSUE_URL,
				err,
			),
		)
	}

	return context
}

func (host *V8ScriptHost) CreateClass(
	name string,
	extends js.Class[jsTypeParam],
	callback js.FunctionCallback[jsTypeParam],
) js.Class[jsTypeParam] {
	ft := wrapV8Callback(host, callback)
	ft.InstanceTemplate().SetInternalFieldCount(1)
	if extends != nil {
		ft.Inherit(extends.(v8Class).ft)
	}
	result := newV8Class(host, ft)
	host.windowTemplate.Set(name, ft)
	host.globals.namedGlobals[name] = result
	return result
}

func UnhandledError(scope js.Scope[jsTypeParam], err error) {
	scope.Window().DispatchEvent(event.NewErrorEvent(err))
}
