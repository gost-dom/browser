package browsertest

import (
	"net/http"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/gost-dom/browser/v8browser"
)

// InitBrowser creates a browser with a V8 engine and a default set of options.
//
// This browser will be configured to log to the t instance. As a consequence,
// uncaught JavaScript errors will result in a test error.
func InitBrowser(t testing.TB, handler http.Handler) *browser.Browser {
	b := v8browser.New(
		browser.WithHandler(handler),
		browser.WithLogger(gosttest.NewTestLogger(t)),
	)
	t.Cleanup(b.Close)
	return b
}

// InitBrowser creates a browser and an empty window with a V8 engine and a
// default set of options.
//
// See also: [InitBrowser]
func InitWindow(t testing.TB) htmltest.WindowHelper {
	b := InitBrowser(t, nil)
	return htmltest.NewWindowHelper(t, b.NewWindow())
}
