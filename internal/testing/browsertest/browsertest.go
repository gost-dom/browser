package browsertest

import (
	"net/http"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/gost-dom/browser/scripting/v8engine"
	"github.com/gost-dom/browser/v8browser"
)

// InitBrowser creates a browser with a script engine and a default set of
// options. If no engine is passed, V8 will be used.
//
// This browser will be configured to log to the t instance. As a consequence,
// uncaught JavaScript errors will result in a test error.
func InitBrowser(t testing.TB, handler http.Handler, engine html.ScriptEngine) *browser.Browser {
	if engine == nil {
		engine = v8engine.DefaultEngine()
	}
	b := v8browser.New(
		browser.WithHandler(handler),
		browser.WithLogger(gosttest.NewTestLogger(t)),
	)
	t.Cleanup(b.Close)
	return b
}

// InitBrowser creates a browser and an empty window with a script engine and a
// default set of options.
//
// See also: [InitBrowser]
func InitWindow(t testing.TB, engine html.ScriptEngine) htmltest.WindowHelper {
	b := InitBrowser(t, nil, engine)
	return htmltest.NewWindowHelper(t, b.NewWindow())
}
