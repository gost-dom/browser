package dom

import (
	"errors"
	"io"
	"net/http"
	netURL "net/url"
	"strings"
)

type ScriptEngineFactory interface {
	NewScriptEngine(window Window) ScriptEngine
	Dispose()
}

type ScriptEngine interface {
	// Run a script, and convert the result to a Go type. This will result in an
	// error if the returned value cannot be represented as a Go type.
	Eval(script string) (any, error)
	// Run a script, ignoring any returned value
	Run(script string) error
	Dispose()
}

type DOMParser interface {
	// Parses a HTML or XML from an [io.Reader] instance. The parsed nodes will
	// have reference the window, e.g. letting events bubble to the window itself.
	// The document pointer will be replaced by the created document.
	//
	// The document is updated using a pointer rather than returned as a value, as
	// parseing process can e.g. execute script tags that require the document to
	// be set on the window _before_ the script is executed.
	ParseReader(window Window, document *Document, reader io.Reader) error
}

// TODO: Remove
type domParser struct{}

func (p domParser) ParseReader(window Window, document *Document, reader io.Reader) error {
	*document = NewDocument(window)
	return parseIntoDocument(window, *document, reader)
}

func NewDOMParser() DOMParser { return domParser{} }

type Window interface {
	EventTarget
	Document() Document
	Dispose()
	// TODO: Remove, for testing
	LoadHTML(string) error
	Eval(string) (any, error)
	Run(string) error
	SetScriptRunner(ScriptEngine)
	Location() Location
	NewXmlHttpRequest() XmlHttpRequest
}

type window struct {
	eventTarget
	document            Document
	scriptEngineFactory ScriptEngineFactory
	scriptEngine        ScriptEngine
	httpClient          http.Client
	baseLocation        string
	domParser           DOMParser
}

func newWindow(windowOptions ...WindowOption) *window {
	var options WindowOptions
	for _, option := range windowOptions {
		option.Apply(&options)
	}
	result := &window{
		eventTarget:         newEventTarget(),
		httpClient:          options.HttpClient,
		baseLocation:        options.BaseLocation,
		scriptEngineFactory: options.ScriptEngineFactory,
		domParser:           options.DOMParser,
	}
	if result.domParser == nil {
		result.domParser = domParser{}
	}
	result.initScriptEngine()
	result.document = NewDocument(result)
	return result
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
	resp, err := result.httpClient.Get(location)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Non-ok Response")
	}
	err = result.parseReader(resp.Body)
	return result, err
}

func (w *window) initScriptEngine() {
	factory := w.scriptEngineFactory
	engine := w.scriptEngine
	if engine != nil {
		engine.Dispose()
	}
	if factory != nil {
		w.scriptEngine = factory.NewScriptEngine(w)
	}
}

func NewWindowReader(reader io.Reader, windowOptions ...WindowOption) (Window, error) {
	window := newWindow(windowOptions...)
	err := window.parseReader(reader)
	return window, err
}

func (w *window) parseReader(reader io.Reader) error {
	err := w.domParser.ParseReader(w, &w.document, reader)
	if err == nil {
		w.document.DispatchEvent(NewCustomEvent(DocumentEventDOMContentLoaded))
		// 'load' is emitted when css and images are loaded, not relevant yet, so
		// just emit it right await
		w.document.DispatchEvent(NewCustomEvent(DocumentEventLoad))
	}
	return err
}

type WindowOptions struct {
	ScriptEngineFactory
	HttpClient   http.Client
	BaseLocation string
	DOMParser    DOMParser
}

type WindowOption interface {
	Apply(options *WindowOptions)
}

type WindowOptionFunc func(*WindowOptions)

func (f WindowOptionFunc) Apply(options *WindowOptions) { f(options) }

func WindowOptionLocation(location string) WindowOptionFunc {
	return func(options *WindowOptions) {
		options.BaseLocation = location
	}
}

func (o WindowOptions) Apply(options *WindowOptions) {
	*options = o
}

func (w *window) Document() Document {
	return w.document
}

func (w *window) LoadHTML(html string) error {
	return w.parseReader(strings.NewReader(html))
}

func (w *window) Run(script string) error {
	if w.scriptEngine != nil {
		return w.scriptEngine.Run(script)
	}
	return errors.New("Script engine not initialised")
}

func (w *window) Eval(script string) (any, error) {
	if w.scriptEngine != nil {
		return w.scriptEngine.Eval(script)
	}
	return nil, errors.New("Script engine not initialised")
}

func (w *window) SetScriptRunner(r ScriptEngine) {
	w.scriptEngine = r
}

func (w *window) Location() Location {
	var u *netURL.URL
	if w.baseLocation != "" {
		u, _ = netURL.Parse(w.baseLocation)
	} else {
		u = new(netURL.URL)
	}
	return NewLocationFromNetURL(u)
}

func (w *window) Dispose() {
	if w.scriptEngine != nil {
		w.scriptEngine.Dispose()
	}
}

func (w *window) NewXmlHttpRequest() XmlHttpRequest {
	return NewXmlHttpRequest(w.httpClient)
}
