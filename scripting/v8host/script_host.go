package v8host

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"sync"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting"
	"github.com/gost-dom/v8go"
)

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

func (pool *scriptHostPool) add(iso *V8ScriptHost) {
	pool.m.Lock()
	defer pool.m.Unlock()

	pool.isolates = append(pool.isolates, iso)
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

// disposable represents a resource that needs cleanup when a context is closed.
// E.g., cgo handles that need to be released.
type disposable interface{ dispose() }

type globalInstall struct {
	name        string
	constructor *v8go.FunctionTemplate
}

type globals struct {
	namedGlobals map[string]*v8go.FunctionTemplate
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
}

type jsConstructorFactory = func(*V8ScriptHost) *v8go.FunctionTemplate

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
				panic(
					"Missing super class spec. Class: " + class.name + ". Super: " + class.superClassName,
				)
			}
			superClassConstructor = iter(superClassSpec)
		}
		constructor := class.factory(host)
		if superClassConstructor != nil {
			constructor.Inherit(superClassConstructor)
		}
		uniqueNames[class.name] = constructor
		result = append(result, globalInstall{class.name, constructor})
		return constructor
	}
	for _, class := range classes {
		iter(class)
	}
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

func registerJSClass(
	className string,
	superClassName string,
	constructorFactory jsConstructorFactory,
) {
	spec := classSpec{
		className, superClassName, constructorFactory,
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
	registerJSClass("NamedNodeMap", "", createNamedNodeMap)
	registerJSClass("Location", "", createLocationPrototype)
	registerJSClass("NodeList", "", createNodeList)
	registerJSClass("EventTarget", "", createEventTarget)
	registerJSClass("XMLHttpRequestEventTarget", "EventTarget", createIllegalConstructor)

	registerJSClass("Document", "Node", createDocumentPrototype)
	registerJSClass("HTMLDocument", "Document", createHTMLDocumentPrototype)
	registerJSClass("DocumentFragment", "Node", createDocumentFragmentPrototype)
	registerJSClass("ShadowRoot", "DocumentFragment", createShadowRootPrototype)
	registerJSClass("Attr", "Node", createAttr)

	registerJSClass("DOMParser", "", createDOMParserPrototype)

	for _, cls := range scripting.HtmlElements {
		if _, found := classes[cls]; !found {
			registerJSClass(cls, "HTMLElement", createIllegalConstructor)
		}
	}
}

func createHostInstance(config hostOptions) *V8ScriptHost {
	if res, ok := pool.tryGet(); ok {
		return res
	}

	host := &V8ScriptHost{
		mu:     new(sync.Mutex),
		iso:    v8go.NewIsolate(),
		logger: config.logger,
	}
	host.inspectorClient = v8go.NewInspectorClient(consoleAPIMessageFunc(host.consoleAPIMessage))
	host.inspector = v8go.NewInspector(host.iso, host.inspectorClient)

	host.iso.SetPromiseRejectedCallback(host.promiseRejected)

	globalInstalls := createGlobals(host)
	host.globals = globals{make(map[string]*v8go.FunctionTemplate)}
	for _, globalInstall := range globalInstalls {
		host.globals.namedGlobals[globalInstall.name] = globalInstall.constructor
	}
	constructors := host.globals.namedGlobals
	window := constructors["Window"]
	host.windowTemplate = window.InstanceTemplate()
	host.contexts = make(map[*v8go.Context]*V8ScriptContext)
	installGlobals(window, host, globalInstalls)
	installEventLoopGlobals(host, host.windowTemplate)
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

	ctx.eventLoop.errorCb(err)
	log.Error(host.logger, "Rejected promise", log.ErrAttr(err))
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

	pool.add(host)
}

// Dispose all pooled isolates
func Shutdown() {
	for _, host := range pool.releaseAll() {
		host.inspectorClient.Dispose()
		host.inspector.Dispose()
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
	// TODO: The possibility to use nil is primarily for testing support
	context := &V8ScriptContext{
		host:    host,
		clock:   clock.New(),
		v8ctx:   v8go.NewContext(host.iso, host.windowTemplate),
		window:  w,
		v8nodes: make(map[entity.ObjectId]*v8go.Value),
	}
	host.addContext(context)
	errorCallback := func(err error) {
		if w != nil {
			w.DispatchEvent(event.NewErrorEvent(err))
		}
	}
	context.eventLoop = newEventLoop(context, errorCallback)
	host.inspector.ContextCreated(context.v8ctx)
	if w != nil {
		context.cacheNode(context.v8ctx.Global(), w)
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

// <<<<<<< HEAD
// =======
//
// func (ctx *V8ScriptContext) Clock() html.Clock { return ctx.clock }
//
// func (host *V8ScriptHost) addContext(ctx *V8ScriptContext) {
// 	host.mu.Lock()
// 	defer host.mu.Unlock()
// 	host.contexts[ctx.v8ctx] = ctx
// }
//
// func (ctx *V8ScriptContext) Close() {
// 	if ctx.disposed {
// 		panic("Context already disposed")
// 	}
// 	ctx.disposed = true
// 	ctx.host.inspector.ContextDestroyed(ctx.v8ctx)
// 	log.Debug(ctx.host.logger,
// 		"ScriptContext: Dispose")
// 	for _, dispose := range ctx.disposers {
// 		dispose.dispose()
// 	}
// 	delete(ctx.host.contexts, ctx.v8ctx)
// 	ctx.v8ctx.Close()
// }
//
// func (ctx *V8ScriptContext) addDisposer(disposer disposable) {
// 	ctx.disposers = append(ctx.disposers, disposer)
// }
//
// func (ctx *V8ScriptContext) runScript(script string) (res *v8.Value, err error) {
// 	res, err = ctx.v8ctx.RunScript(script, "")
// 	ctx.eventLoop.tick()
// 	return
// }
//
// func (ctx *V8ScriptContext) Run(script string) error {
// 	_, err := ctx.runScript(script)
// 	return err
// }
//
// func (ctx *V8ScriptContext) Eval(script string) (interface{}, error) {
// 	result, err := ctx.runScript(script)
// 	if err == nil {
// 		return v8ValueToGoValue(result)
// 	}
// 	return nil, err
// }
//
// func (ctx *V8ScriptContext) EvalCore(script string) (any, error) {
// 	return ctx.runScript(script)
// }
//
// func (ctx *V8ScriptContext) RunFunction(script string, arguments ...any) (res any, err error) {
// 	var (
// 		v  *v8.Value
// 		f  *v8.Function
// 		ok bool
// 	)
// 	if v, err = ctx.runScript(script); err == nil {
// 		f, err = v.AsFunction()
// 	}
// 	if err == nil {
// 		args := make([]v8.Valuer, len(arguments))
// 		for i, a := range arguments {
// 			if args[i], ok = a.(v8.Valuer); !ok {
// 				err = fmt.Errorf("V8ScriptContext.RunFunction: Arguments is not a V8 value: %d", i)
// 			}
// 		}
// 		return f.Call(ctx.v8ctx.Global(), args...)
// 	}
// 	return
// }
//
// func (ctx *V8ScriptContext) Export(val any) (any, error) {
// 	if res, ok := val.(*v8.Value); ok {
// 		return v8ValueToGoValue(res)
// 	} else {
// 		return nil, errors.New("V8ScriptContext.Export: value not a V8 value")
// 	}
// }
//
// func v8ValueToGoValue(result *v8go.Value) (interface{}, error) {
// 	if result == nil {
// 		return nil, nil
// 	}
// 	if result.IsBoolean() {
// 		return result.Boolean(), nil
// 	}
// 	if result.IsInt32() {
// 		return result.Int32(), nil
// 	}
// 	if result.IsString() {
// 		return result.String(), nil
// 	}
// 	if result.IsNull() {
// 		return nil, nil
// 	}
// 	if result.IsUndefined() {
// 		return nil, nil
// 	}
// 	if result.IsArray() {
// 		obj, _ := result.AsObject()
// 		length, err := obj.Get("length")
// 		l := length.Uint32()
// 		errs := make([]error, l+1)
// 		errs[0] = err
// 		result := make([]any, l)
// 		for i := uint32(0); i < l; i++ {
// 			val, err := obj.GetIdx(i)
// 			if err == nil {
// 				result[i], err = v8ValueToGoValue(val)
// 			}
// 			errs[i+1] = err
// 		}
// 		return result, errors.Join(errs...)
// 	}
// 	return nil, fmt.Errorf("Value not yet supported: %v", *result)
// }
// >>>>>>> a12e9dc (Creating a pool of isolates)
