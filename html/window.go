package html

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	netURL "net/url"
	"strings"
	"time"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/url"
)

var ErrTooManyRedirects = errors.New("Too many redirects")
var ErrCancelled = errors.New("Cancelled")

type ScriptEngineOptions struct {
	HttpClient *http.Client
	Logger     *slog.Logger
}

// ScriptEngine represents a JavaScript interpreter with a specific global scope
// configuration.
//
// A script engine is used to create multiple script contexts, all with the same
// initial values in global scope.
//
// E.g., v8 has the root level of an "Isolate", on which you create "templates"
// for values that must exist in global scope; including inheritance
// hierarchies. Gost-DOM can reuse isolates between tests, reducing the amount
// of work necessary to recreate global scope.
//
// Multiple engines can be used, if client code has different scenarios
// requiring different APIs exposed to global scope.
type ScriptEngine interface {
	NewHost(ScriptEngineOptions) ScriptHost
}

// A script host is a cacheable initialized "host" created from the script
// engine. Each host can create a new [ScriptContext], representing a new global
// scope.
//
// Client code should call Close() when done using the host, allowing it to be
// returned to the cache.
type ScriptHost interface {
	NewContext(window Window) ScriptContext
	Close()
}

type Clock interface {
	// Deprecated: Call ProcessEvents() instead
	RunAll() error
	// Advance simulated time by a specific duration. Will run relevant
	// callbacks registered with setInterval or setTimeout
	Advance(time.Duration) error
	// ProcessEvents ensures that all immediate functions as well as
	// microtasks are executed.
	ProcessEvents(ctx context.Context) error
	// ProcessEventsWhile will process events in the event loop as long as a
	// callback function f returns true. This has two specific uses over the
	// simpler ProcessEvents
	//
	// - ProcessEvents will not stop when code uses setInterval as there is
	// always an event scheduled in the future
	// - ProcessEvents will stop if it doesn't know about a future event
	//
	// Example, clicking a button result in an HTTP request. The server will
	// eventually push a web socket message. Using ProcessEvents could fail due
	// to a race condition.
	//
	// For these types of tests, Go's synctest package is worth considering as
	// well.
	ProcessEventsWhile(ctx context.Context, f func() bool) error
}

// Describes a current browser context
type BrowsingContext interface {
	// Logger returns the currently configured logger for the window. Returns
	// nil if no instance is created.
	Logger() *slog.Logger
	Context() context.Context
	HTTPClient() http.Client
	LocationHREF() string
}

type ScriptContext interface {
	// Run a script, and convert the result to a Go type. Only use this if you
	// need the return value, otherwise call Run.
	//
	// If the evaluated JS value cannot be converted to a Go value, an error is
	// returned.
	//
	// Deprecated: Call Compile and then eval the compiled script
	Eval(script string) (any, error)
	// Run a script. This is should be used instead of eval when the return value
	// is not needed, as eval returns an error when the return value cannot be
	// converted to a go type.
	//
	// Deprecated: Call Compile and then run the compiled script
	Run(script string) error
	// Compile a script that can later be executed.
	Compile(script string) (Script, error)
	// Download a script from the specified url
	DownloadScript(url string) (Script, error)
	// Download an ECMAScript Module from the specified url
	DownloadModule(url string) (Script, error)
	Clock() Clock
	Close()
}

type Script interface {
	// Run a script, and convert the result to a Go type. Only use this if you
	// need the return value, otherwise call Run.
	//
	// If the evaluated JS value cannot be converted to a Go value, an error is
	// returned.
	Eval() (any, error)
	// Run a script. This is should be used instead of eval when the return value
	// is not needed, as eval returns an error when the return value cannot be
	// converted to a go type.
	Run() error
}

type Window interface {
	event.EventTarget
	BrowsingContext
	entity.ObjectIder
	entity.Components
	Document() HTMLDocument
	Close()
	Clock() Clock
	// Open an HTML document from an href. A URL will be opened, but a path will
	// be resolved based on the current location.
	Navigate(href string) error
	LoadHTML(string) error // TODO: Remove, for testing
	// Eval calls [ScriptContext.Eval]
	Eval(string) (any, error)
	// Run calls [ScriptContext.Run]
	Run(string) error
	ScriptContext() ScriptContext
	Location() Location
	History() *History
	ParseFragment(ownerDocument dom.Document, reader io.Reader) (dom.DocumentFragment, error)
	// unexported

	fetchRequest(req *http.Request) error
	resolveHref(string) *url.URL
	window() *window
}

type window struct {
	event.EventTarget
	entity.Entity
	document            HTMLDocument
	history             *History
	scriptEngineFactory ScriptHost
	scriptContext       ScriptContext
	httpClient          http.Client
	// baseLocation        string
	logger          *slog.Logger
	deferredScripts []*htmlScriptElement
	context         context.Context
}

func newWindow(windowOptions ...WindowOption) *window {
	var options WindowOptions
	for _, option := range windowOptions {
		option.Apply(&options)
	}
	ctx := options.Context
	if ctx == nil {
		ctx = context.Background()
	}
	baseLocation := options.BaseLocation
	win := &window{
		EventTarget: event.NewEventTarget(),
		httpClient:  options.HttpClient,
		// baseLocation:        options.BaseLocation,
		scriptEngineFactory: options.ScriptHost,
		logger:              options.Logger,
		history:             new(History),
		context:             ctx,
	}
	if baseLocation == "" {
		baseLocation = "about:blank"
	}
	win.history.window = win
	win.history.pushLoad(baseLocation)
	url, err := url.NewUrl(baseLocation)
	if err != nil {
		win.Logger().
			Warn("newWindow: Error parsing location", slog.String("location", options.BaseLocation), log.ErrAttr(err))
	}
	win.document = NewHTMLDocument(win, newLocation(url))
	event.SetEventTargetSelf(win)
	win.initScriptEngine()
	return win
}

func (w *window) deferScript(e *htmlScriptElement) {
	w.deferredScripts = append(w.deferredScripts, e)
}

func (w *window) checkRedirect(req *http.Request, via []*http.Request) error {
	if len(via) > 9 {
		return ErrTooManyRedirects
	}
	w.history.ReplaceState(EMPTY_STATE, req.URL.String())
	return nil
}

func NewWindow(windowOptions ...WindowOption) Window {
	return newWindow(windowOptions...)
}

func OpenWindowFromLocation(location string, windowOptions ...WindowOption) (Window, error) {
	var options WindowOptions
	for _, option := range windowOptions {
		option.Apply(&options)
	}
	if options.BaseLocation != "" {
		u, err := netURL.Parse(options.BaseLocation)
		if err == nil {
			location = u.JoinPath(location).String()
		}
	} else {
		options.BaseLocation = location
	}
	result := newWindow(options)
	return result, result.get(location)
}

func (w *window) HTTPClient() http.Client  { return w.httpClient }
func (w *window) Document() HTMLDocument   { return w.document }
func (w *window) History() *History        { return w.history }
func (w *window) Context() context.Context { return w.context }

func (w *window) initScriptEngine() {
	w.EventTarget.RemoveAll()
	factory := w.scriptEngineFactory
	engine := w.scriptContext
	if engine != nil {
		engine.Close()
	}
	if factory != nil {
		w.scriptContext = factory.NewContext(w)
	}
}

func (w *window) setBaseLocation(href string) string {
	loc := w.document.location()
	if href != "" && loc != nil {
		url := w.resolveHref(href)
		loc.set(url)
	}
	return w.document.Location().Href()
}

func (w *window) ParseFragment(
	ownerDocument dom.Document,
	reader io.Reader,
) (dom.DocumentFragment, error) {
	return dom.ParseFragment(ownerDocument, reader)
}

// NewWindowReader will create a new window and load parse the HTML from the
// reader. If there is an error reading from the stream, or parsing the DOM, an
// error is returned.
//
// If this function returns without an error, the DOM will have been parsed and
// the DOMContentLoaded event has been dispached on the [dom.Document]
//
// Experimental: This function will likely be removed in favour of other ways of
// creating an initialized window
func NewWindowReader(
	reader io.Reader,
	url *url.URL,
	windowOptions ...WindowOption,
) (Window, error) {
	window := newWindow(windowOptions...)
	err := window.parseReader(reader, url)
	return window, err
}

func (w *window) parseReader(reader io.Reader, u *url.URL) error {
	l := w.document.location()
	if u != nil {
		l = newLocation(u)
	}
	w.document = NewEmptyHtmlDocument(w, l)
	err := dom.ParseDocument(w.document, reader)
	for _, s := range w.deferredScripts {
		s.run()
	}
	if err == nil {
		w.document.DispatchEvent(&event.Event{Type: dom.DocumentEventDOMContentLoaded})
		// 'load' is emitted when css and images are loaded, not relevant yet, so
		// just emit it right await
		w.document.DispatchEvent(&event.Event{Type: dom.DocumentEventLoad})
	}
	if el, _ := w.document.QuerySelector("[autofocus]"); el != nil {
		if el, ok := el.(HTMLElement); ok {
			el.Focus()
		}
	}

	return err
}

func (w *window) Navigate(href string) (err error) {
	w.Logger().Info("Window.navigate", "href", href)
	if href != "about:blank" {
		href = w.resolveHref(href).String()
	}
	defer func() {
		if err != nil {
			w.Logger().Warn("Window.navigate: Error response", log.ErrAttr(err))
		}
	}()
	w.History().pushLoad(href)
	w.initScriptEngine()
	return w.get(href)
}

// reload is used internally to load a page into the browser, but without
// affecting the history
func (w *window) reload(href string) error {
	w.Logger().Debug("Window.reload:", "href", href)
	w.initScriptEngine()
	return w.get(href)
}

func (w *window) get(href string) error {
	url := url.ParseURL(href)
	if href == "about:blank" {
		w.document = NewHTMLDocument(w, newLocation(url))
		return nil
	} else if req, err := http.NewRequest("GET", href, nil); err == nil {
		err := w.fetchRequest(req)
		return err
	} else {
		return err
	}
}

// fetchRequest handles "User agent" requests. I.e., navigation by clicking a
// link, submitting a form, of calling a JS API that makes the browser itself
// navigate to a new URL.
func (w *window) fetchRequest(req *http.Request) error {
	// Create a copy of the client, and set the CheckRedirect to a function that
	// updates the window location to reflect the new URL.
	client := w.httpClient
	client.CheckRedirect = w.checkRedirect
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Non-ok Response")
	}
	u := url.ParseURL(resp.Request.URL.String())
	return w.parseReader(resp.Body, u)
}

func (w *window) LoadHTML(html string) error {
	return w.parseReader(strings.NewReader(html), nil)
}

func (w *window) Run(script string) error {
	if w.scriptContext != nil {
		return w.scriptContext.Run(script)
	}
	return errors.New("Script engine not initialised")
}

func (w *window) Eval(script string) (any, error) {
	if w.scriptContext != nil {
		return w.scriptContext.Eval(script)
	}
	return nil, errors.New("Script engine not initialised")
}

func (w *window) ScriptContext() ScriptContext { return w.scriptContext }

func (w *window) Location() Location { return w.document.Location() }

func (w *window) Clock() Clock { return w.scriptContext.Clock() }

func (w *window) LocationHREF() string {

	return w.Location().Href()
}

func (w *window) Close() {
	if w.scriptContext != nil {
		w.scriptContext.Close()
	}
}

func (w *window) ObjectId() entity.ObjectId { return -1 }

// resolveHref takes an href from a <a> tag, or action from a <form> tag and
// resolves an absolute URL that must be requested.
func (w *window) resolveHref(href string) *url.URL {
	r, err := url.NewUrlBase(href, w.Location().Href())
	if err != nil {
		panic(err)
	}
	return r
}

func (w *window) Logger() log.Logger {
	if w.logger != nil {
		return w.logger
	}
	return log.Default()
}

func (w *window) window() *window { return w }

type WindowOptions struct {
	ScriptHost
	HttpClient   http.Client
	BaseLocation string
	Logger       *slog.Logger
	Context      context.Context
}

type WindowOption interface {
	Apply(options *WindowOptions)
}

type WindowOptionFunc func(*WindowOptions)

func (f WindowOptionFunc) Apply(options *WindowOptions) { f(options) }

func WindowOptionLogger(l *slog.Logger) WindowOptionFunc {
	return func(options *WindowOptions) { options.Logger = l }
}

func WindowOptionLocation(location string) WindowOptionFunc {
	return func(options *WindowOptions) { options.BaseLocation = location }
}

func WindowOptionHost(host ScriptHost) WindowOptionFunc {
	return func(options *WindowOptions) { options.ScriptHost = host }
}

func WindowOptionHTTPClient(client http.Client) WindowOptionFunc {
	return func(options *WindowOptions) { options.HttpClient = client }
}

func (o WindowOptions) Apply(options *WindowOptions) {
	*options = o
}
