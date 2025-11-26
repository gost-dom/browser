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
	"github.com/gost-dom/browser/scripting/v8host"
)

type browserConfig struct {
	client http.Client
	logger *slog.Logger
	engine ScriptEngine
	ctx    context.Context

	ownsHost bool // TODO: Should always own host once WithScriptHost is removed
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

type staticHostEngine struct{ host html.ScriptHost }

// Deprecated: This will be removed in the next release
func (e staticHostEngine) NewHost(html.ScriptEngineOptions) html.ScriptHost {
	host := html.ScriptHost(e.host)
	return host
}

// WithScriptHost uses a specific script host.
//
// Deprecated: Prefer WithScriptEngine
func WithScriptHost(host html.ScriptHost) BrowserOption {
	return func(b *browserConfig) { b.engine = staticHostEngine{host} }
}

func WithScriptEngine(engine html.ScriptEngine) BrowserOption {
	return func(b *browserConfig) { b.engine = engine; b.ownsHost = true }
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
	ownsHost   bool
	mu         sync.Mutex
}

// New initialises a new [Browser]. Options can be one of
//
//   - [WithScriptEngine]
//   - [WithLogger]
//   - [WithHandler]
//   - [WithContext]
//
// Script engine defaults to V8. This will change in the future, but a migration
// path is not ready.
//
// Deprecated: This function WILL change behaviour. Previous behaviour was to
// default to use V8 if nothing was specified. New default will be to have not
// script engine
func New(options ...BrowserOption) *Browser {
	config := &browserConfig{client: NewHttpClient()}
	for _, o := range options {
		o(config)
	}
	engine := config.engine
	ownsHost := config.ownsHost
	if engine == nil {
		engine = v8host.DefaultEngine()
		ownsHost = true
	}
	b := &Browser{
		Client: config.client,
		Logger: config.logger,
		ScriptHost: engine.NewHost(
			html.ScriptEngineOptions{
				Logger:     config.logger,
				HttpClient: &config.client,
			}),
		ctx: config.ctx,

		ownsHost: ownsHost,
	}
	if config.ctx != nil {
		context.AfterFunc(config.ctx, b.Close)
	}
	return b
}

// NewWindow creates a new window. Panics if the browser has been closed
func (b *Browser) NewWindow() Window {
	b.mu.Lock()
	defer b.mu.Unlock()

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
	b.mu.Lock()
	defer b.mu.Unlock()

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

// NewFromHandler initialises a new [Browser] with with an [http.Handler]
//
// Deprecated: Prefer browser.New(browser.WithHandler(...)) instead.
func NewFromHandler(handler http.Handler) *Browser {
	return New(WithHandler(handler))
}

// Deprecated: NewBrowser should not be called. Call New instead.
//
// This method will selfdestruct in 10 commits
func NewBrowser() *Browser {
	return New()
}

// Deprecated: NewBrowserFromHandler should not be called, call, NewFromHandler instead.
//
// This method will selfdestruct in 10 commits
func NewBrowserFromHandler(handler http.Handler) *Browser {
	return NewFromHandler(handler)
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
	b.mu.Lock()
	defer b.mu.Unlock()

	b.logger().Info("Browser: Close()")
	for _, win := range b.windows {
		win.Close()
	}
	if b.ScriptHost != nil && b.ownsHost {
		b.ScriptHost.Close()
	}
	b.closed = true
}

func (b *Browser) Closed() bool { return b.closed }

func (b *Browser) logger() *slog.Logger {
	if b.Logger != nil {
		return b.Logger
	}
	return log.Default()
}
