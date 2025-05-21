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

	registerJSClass("FormData", "", createFormData)
	registerJSClass("DOMParser", "", createDOMParserPrototype)

	for _, cls := range scripting.HtmlElements {
		if _, found := classes[cls]; !found {
			registerJSClass(cls, "HTMLElement", createIllegalConstructor)
		}
	}
}

func New(opts ...HostOption) *V8ScriptHost {
	config := hostOptions{}
	for _, opt := range opts {
		opt(&config)
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
	delete(ctx.host.contexts, ctx.v8ctx)
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
	log.Error(host.logger, "Rejected promise", "err", err)
}

func (host *V8ScriptHost) Logger() log.Logger { return host.logger }

func (host *V8ScriptHost) Close() {
	host.mu.Lock()
	defer host.mu.Unlock()
	undiposedContexts := slices.Collect(maps.Values(host.contexts))
	undisposedCount := len(undiposedContexts)

	if undisposedCount > 0 {
		log.Warn(host.logger, "Closing script host with undisposed contexts", "count", undisposedCount)
		for _, ctx := range undiposedContexts {
			ctx.Close()
		}
	}
	host.inspectorClient.Dispose()
	host.inspector.Dispose()
	host.iso.Dispose()
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
	errorCallback := func(err error) {
		if w != nil {
			w.DispatchEvent(event.NewErrorEvent(err))
		}
	}
	context.eventLoop = newEventLoop(context, errorCallback)
	host.inspector.ContextCreated(context.v8ctx)
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
	if w != nil {
		context.cacheNode(context.v8ctx.Global(), w)
	}
	host.addContext(context)

	return context
}
