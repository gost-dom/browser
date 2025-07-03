package browser

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/gosthttp"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/v8host"
)

type browserConfig struct {
	client http.Client
	logger *slog.Logger
	host   html.ScriptHost
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

func WithScriptHost(host html.ScriptHost) BrowserOption {
	return func(b *browserConfig) { b.host = host }
}

func WithContext(ctx context.Context) BrowserOption {
	return func(b *browserConfig) { b.ctx = ctx }
}

// Pretty stupid right now, but should _probably_ allow handling multiple
// windows/tabs. This used to be the case for _some_ identity providers, but I'm
// not sure if that even work anymore because of browser security.
type Browser struct {
	Client     http.Client
	ScriptHost ScriptHost
	Logger     log.Logger
	ctx        context.Context
	windows    []Window
	closed     bool
	ownsHost   bool
}

// NewWindow creates a new window. Panics if the browser has been closed
func (b *Browser) NewWindow() Window {
	if b.closed {
		panic("gost-dom/browser: browser closed")
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
	// log.Debug("Browser: OpenWindow", "URL", location)
	resp, err := b.Client.Get(location)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Non-ok Response: %d", resp.StatusCode)
	}
	window, err = html.NewWindowReader(resp.Body, b.createOptions(location))
	b.windows = append(b.windows, window)
	return
}

// NewFromHandler initialises a new [Browser] with with an [http.Handler]
//
// Deprecated: Prefer browser.New(browser.WithHandler(...)) instead.
func NewFromHandler(handler http.Handler) *Browser {
	return New(WithHandler(handler))
}

// New initialises a new [Browser] with the default script engine.
func New(options ...BrowserOption) *Browser {
	config := &browserConfig{client: NewHttpClient()}
	for _, o := range options {
		o(config)
	}
	b := &Browser{
		Client: config.client,
		Logger: config.logger,
		ScriptHost: v8host.New(v8host.WithLogger(config.logger),
			v8host.WithHTTPClient(&config.client),
		),
		ctx: config.ctx,
	}
	if b.ScriptHost == nil {
		b.ownsHost = true
		b.ScriptHost = v8host.New(v8host.WithLogger(config.logger))
	}
	if config.ctx != nil {
		context.AfterFunc(config.ctx, b.Close)
	}
	return b
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

func (b *Browser) Close() {
	log.Debug(b.Logger, "Browser: Close()")
	for _, win := range b.windows {
		win.Close()
	}
	if b.ScriptHost != nil && b.ownsHost {
		b.ScriptHost.Close()
	}
	b.closed = true
}

func (b *Browser) Closed() bool { return b.closed }
