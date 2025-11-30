package browser

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/gosthttp"
	"github.com/gost-dom/browser/internal/log"
)

type browserConfig struct {
	client http.Client
	logger *slog.Logger
	engine ScriptEngine
	ctx    context.Context
}

type BrowserOption func(*browserConfig)

func WithLogger(l *slog.Logger) BrowserOption { return func(b *browserConfig) { b.logger = l } }

// WithHandler configures the browser's [http.Client] to use an
// [http.Roundtripper] that bypasses the TCP stack and calls directly into the
// specified handler as a normal function call.
//
// Note: There is a current limitation that NO requests from the browser will be
// sent when using this. So sites will not work if they
//   - Depend on content from CDN
//   - Depend on an external service, e.g., an identity provider.
//
// That is a limitation that was the result of prioritising more important, and
// higher risk features.
func WithHandler(h http.Handler) BrowserOption {
	return func(b *browserConfig) { b.client = NewHttpClientFromHandler(h) }
}

func WithScriptEngine(engine html.ScriptEngine) BrowserOption {
	return func(b *browserConfig) { b.engine = engine }
}

// WithContext passes a [context.Context] than can trigger cancellation, e.g.:
//
//   - Close any open HTTP connections and disconnect from the server.
//   - Release resources, and reuse script hosts.
//
// See also: [Browser.Close]
func WithContext(ctx context.Context) BrowserOption {
	return func(b *browserConfig) { b.ctx = ctx }
}

// Browser contains an initialized browser with a script engine. Create new
// windows by calling [Browser.Open].
//
// Browser values should be closed by calling [Browser.Close], or passing a
// context:
//
//	func TestBrowserWithClose(t *testing.T) {
//		handler := NewRootHttpHandler()
//		b := browser.New(browser.WithHandler(handler))
//		t.Cleanup(func() { b.Close() })
//
//		win, err := b.Open("http://example.com")
//		// ...
//	}
//
// Passing a context:
//
//	func TestBrowserWithContext(t *testing.T) {
//		handler := NewRootHttpHandler()
//		b := browser.New(
//			browser.WithHandler(handler),
//			browser.WithContext(t.Context(),
//		)
//
//		win, err := b.Open("http://example.com")
//		// ...
//	}
type Browser struct {
	Client     http.Client
	ScriptHost ScriptHost
	Logger     log.Logger
	ctx        context.Context
	windows    []Window
	closed     bool

	// closeLock protects the windows slice and closed field. When creating a
	// new window, browser should panic if closed - as it may have a reference
	// to a script host that can be recycled (maybe the script host should be
	// made less likely to a use-after-free issue).
	// Also, the window list is protected, as closing needs to close each
	// window, releasing script contexts. So e.g., creating a new window in
	// parallel with closing could result in an orphaned window, holding a
	// script context.
	// TODO: There are multiple places creating new windows. This makes it easy
	// to create race conditions, as multiple places need to be lock-aware;
	// making it easy to overlook when working with code. Create a more
	// resilient design less likely to break
	closeLock sync.Mutex
}

// New initialises a new [Browser]. Options can be one of
//
//   - [WithScriptEngine]
//   - [WithLogger]
//   - [WithHandler]
//   - [WithContext]
func New(options ...BrowserOption) *Browser {
	config := &browserConfig{client: NewHttpClient()}
	for _, o := range options {
		o(config)
	}
	engine := config.engine
	var host html.ScriptHost
	if engine != nil {
		host = engine.NewHost(
			html.ScriptEngineOptions{
				Logger:     config.logger,
				HttpClient: &config.client,
			})
	}
	b := &Browser{
		Client:     config.client,
		Logger:     config.logger,
		ScriptHost: host,
		ctx:        config.ctx,
	}
	if config.ctx != nil {
		context.AfterFunc(config.ctx, b.Close)
	}
	return b
}

// NewWindow creates a new window. Panics if the browser has been closed
func (b *Browser) NewWindow() Window {
	b.closeLock.Lock()
	defer b.closeLock.Unlock()

	if b.closed {
		panic("gost-dom/browser: NewWindow(): browser closed")
	}
	window := html.NewWindow(b.createOptions(""))
	b.windows = append(b.windows, window)
	return window

}

// Open will open a new [html.Window], loading the specified location. If the
// server does not respons with a 200 status code, an error is returned.
//
// See [html.NewWindowReader] about the return value, and when the window
// returns.
func (b *Browser) Open(location string) (window Window, err error) {
	b.closeLock.Lock()
	defer b.closeLock.Unlock()

	if b.Closed() {
		panic("Open window on a closed browser")
	}

	// log.Debug("Browser: OpenWindow", "URL", location)
	resp, err := b.Client.Get(location)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Non-ok Response: %d", resp.StatusCode)
	}
	respLocation := resp.Request.URL.String()
	window, err = html.NewWindowReader(resp.Body, b.createOptions(respLocation))
	b.windows = append(b.windows, window)
	return
}

func (b *Browser) createOptions(location string) WindowOptions {
	return WindowOptions{
		ScriptHost:   b.ScriptHost,
		HttpClient:   b.Client,
		BaseLocation: location,
		Logger:       b.Logger,
		Context:      b.ctx,
	}
}

// Close "closes" a browser, releasing resources. This will close any
// initialized script hosts and contexts. Has two purposes.
//
//   - Reuse a template engine, reducing engine initialization overhead.
//   - Release memory for non-Go engines, e.g., V8
//
// The relevance depends mostly on the script engine. For a pure Go engine,
// resources would be garbage collections. And the ability to reuse a
// preconfigured engine depends on engine capabilities.
//
// Note: If a browser is initialized by passing a [context.Context] to the
// [WithContext] option, it will be closed if the context is cancelled.
func (b *Browser) Close() {
	b.closeLock.Lock()
	defer b.closeLock.Unlock()
	if b.Closed() {
		return
	}

	b.logger().Info("Browser: Close()")

	for _, win := range b.windows {
		win.Close()
	}
	if b.ScriptHost != nil {
		b.ScriptHost.Close()
	}
	b.windows = nil
	b.ScriptHost = nil
	b.closed = true
}

func (b *Browser) Closed() bool { return b.closed }

func (b *Browser) logger() *slog.Logger {
	if b.Logger != nil {
		return b.Logger
	}
	return log.Default()
}
