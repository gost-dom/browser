package browsertest

import (
	"fmt"
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

// By default, log entries of Error log level cause the test to fail. This
// option will suppress that behaviour; for when verifying explicit error
// scenarios.
func WithIgnoreErrorLogs() InitOption {
	return WithLogOption(gosttest.AllowErrors())
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
	if handler == nil {
		handler = http.HandlerFunc(dummyHttpServer)
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

func dummyHttpServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<body><h1>Dummy page</h1></body>")
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
