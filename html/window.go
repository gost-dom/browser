package html

import (
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

type ScriptHost interface {
	NewContext(window Window) ScriptContext
	Close()
}

type Clock interface {
	RunAll() error
	Advance(time.Duration) error
}

// Describes a current browser context
type BrowsingContext interface {
	// Logger returns the currently configured logger for the window. Returns
	// nil if no instance is created.
	Logger() *slog.Logger
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
	Document() dom.Document
	Close()
	Clock() Clock
	Navigate(string) error // TODO: Remove, perhaps? for testing
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
}

type window struct {
	event.EventTarget
	document            dom.Document
	history             *History
	scriptEngineFactory ScriptHost
	scriptContext       ScriptContext
	httpClient          http.Client
	baseLocation        string
	domParser           domParser
	logger              *slog.Logger
	deferredScripts     []*htmlScriptElement
}

func newWindow(windowOptions ...WindowOption) *window {
	var options WindowOptions
	for _, option := range windowOptions {
		option.Apply(&options)
	}
	win := &window{
		EventTarget:         event.NewEventTarget(),
		httpClient:          options.HttpClient,
		baseLocation:        options.BaseLocation,
		scriptEngineFactory: options.ScriptHost,
		logger:              options.Logger,
		history:             new(History),
	}
	if win.baseLocation == "" {
		win.baseLocation = "about:blank"
	}
	win.history.window = win
	win.history.pushLoad(win.baseLocation)
	win.domParser = domParser{}
	win.initScriptEngine()
	win.document = NewHTMLDocument(win)
	event.SetEventTargetSelf(win)
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
	if href == "" {
		return w.baseLocation
	}
	w.baseLocation = w.resolveHref(href).Href()
	return w.baseLocation
}

func (w *window) History() *History {
	return w.history
}

func (w *window) ParseFragment(
	ownerDocument dom.Document,
	reader io.Reader,
) (dom.DocumentFragment, error) {
	return w.domParser.ParseFragment(ownerDocument, reader)
}

// NewWindowReader will create a new window and load parse the HTML from the
// reader. If there is an error reading from the stream, or parsing the DOM, an
// error is returned.
//
// If this function returns without an error, the DOM will have been parsed and
// the DOMContentLoaded event has been dispached on the [dom.Document]
func NewWindowReader(reader io.Reader, windowOptions ...WindowOption) (Window, error) {
	window := newWindow(windowOptions...)
	err := window.parseReader(reader)
	return window, err
}

func (w *window) parseReader(reader io.Reader) error {
	err := w.domParser.ParseReader(w, &w.document, reader)
	for _, s := range w.deferredScripts {
		s.run()
	}
	if err == nil {
		w.document.DispatchEvent(event.New(dom.DocumentEventDOMContentLoaded, nil))
		// 'load' is emitted when css and images are loaded, not relevant yet, so
		// just emit it right await
		w.document.DispatchEvent(event.New(dom.DocumentEventLoad, nil))
	}
	if el, _ := w.document.QuerySelector("[autofocus]"); el != nil {
		if el, ok := el.(HTMLElement); ok {
			el.Focus()
		}
	}

	return err
}

func (w *window) HTTPClient() http.Client { return w.httpClient }

func (w *window) Document() dom.Document {
	return w.document
}

func (w *window) Navigate(href string) (err error) {
	log.Info(w.Logger(), "Window.navigate:", "href", href)
	defer func() {
		if err != nil {
			log.Warn(w.logger, "Window.navigate: Error response", "err", err.Error())
		}
	}()
	w.History().pushLoad(href)
	w.initScriptEngine()
	w.baseLocation = href
	if href == "about:blank" {
		w.document = NewHTMLDocument(w)
		return nil
	} else {
		return w.get(href)
	}
}

// reload is used internally to load a page into the browser, but without
// affecting the history
func (w *window) reload(href string) error {
	log.Debug(w.Logger(), "Window.reload:", "href", href)
	w.initScriptEngine()
	w.baseLocation = href
	if href == "about:blank" {
		w.document = NewHTMLDocument(w)
		return nil
	} else {
		return w.get(href)
	}
}

func (w *window) get(href string) error {
	if req, err := http.NewRequest("GET", href, nil); err == nil {
		return w.fetchRequest(req)
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
	return w.parseReader(resp.Body)
}

func (w *window) LoadHTML(html string) error {
	return w.parseReader(strings.NewReader(html))
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

func (w *window) Location() Location {
	var u *netURL.URL
	if w.baseLocation != "" {
		u, _ = netURL.Parse(w.baseLocation)
	} else {
		u = new(netURL.URL)
	}
	return newLocation(u)
}

func (w *window) Clock() Clock { return w.scriptContext.Clock() }

func (w *window) LocationHREF() string { return w.baseLocation }

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

func (w *window) Logger() log.Logger { return w.logger }

type WindowOptions struct {
	ScriptHost
	HttpClient   http.Client
	BaseLocation string
	Logger       *slog.Logger
}

type WindowOption interface {
	Apply(options *WindowOptions)
}

type WindowOptionFunc func(*WindowOptions)

func (f WindowOptionFunc) Apply(options *WindowOptions) { f(options) }

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
