package browsertest

import (
	"log/slog"
	"net/http"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
)

type windowOption func(html.Window)

type option struct {
	logOptions    []gosttest.HandlerOption
	windowOptions []windowOption
}

type InitOption func(*option)

func withWindowOption(wo windowOption) InitOption {
	return func(o *option) {
		o.windowOptions = append(o.windowOptions, wo)
	}
}

func WithLogOption(lo gosttest.HandlerOption) InitOption {
	return func(o *option) { o.logOptions = append(o.logOptions, lo) }
}

func WithMinLogLevel(lvl slog.Level) InitOption {
	return WithLogOption(gosttest.MinLogLevel(lvl))
}

func withHtml(h string) windowOption {
	return func(w html.Window) {
		w.LoadHTML(h)
	}
}

func WithHtml(html string) InitOption {
	return withWindowOption(windowOption(withHtml(html)))
}

func asOptions(o option) InitOption {
	return func(other *option) {
		*other = o
	}
}

// InitBrowser creates a browser with a script engine and a default set of
// options. If no engine is passed, V8 will be used.
//
// This browser will be configured to log to the t instance. As a consequence,
// uncaught JavaScript errors will result in a test error.
func InitBrowser(
	t testing.TB,
	handler http.Handler,
	engine html.ScriptEngine,
	opts ...InitOption,
) htmltest.BrowserHelper {
	var o option
	for _, opt := range opts {
		opt(&o)
	}
	logger := gosttest.NewTestLogger(t, o.logOptions...)
	b := htmltest.NewBrowserHelper(t, browser.New(
		browser.WithScriptEngine(engine),
		browser.WithHandler(handler),
		browser.WithLogger(logger),
	))
	t.Cleanup(b.Close)
	return b
}

// InitWindow creates a browser and an empty window with a script engine and a
// default set of options.
//
// See also: [InitBrowser]
func InitWindow(t testing.TB, engine html.ScriptEngine, opts ...InitOption) htmltest.WindowHelper {
	var o option
	for _, opt := range opts {
		opt(&o)
	}
	b := InitBrowser(t, nil, engine, asOptions(o))
	win := b.NewWindow()
	for _, wo := range o.windowOptions {
		wo(win)
	}
	return htmltest.NewWindowHelper(t, win)
}
