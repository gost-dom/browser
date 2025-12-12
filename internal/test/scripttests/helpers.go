package scripttests

import (
	"net/http"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
)

var initBrowser = browsertest.InitBrowser

func initWindow(
	t *testing.T,
	e html.ScriptEngine,
	h http.Handler,
	opts ...InitOption,
) htmltest.WindowHelper {
	var o option
	for _, opt := range opts {
		opt(&o)
	}
	logger := gosttest.NewTestLogger(t, o.logOptions...)
	ctx := o.ctx
	if ctx == nil {
		ctx = t.Context()
	}
	b := htmltest.NewBrowserHelper(t, browser.New(
		browser.WithContext(ctx),
		browser.WithLogger(logger),
		browser.WithHandler(h),
		browser.WithScriptEngine(e),
	))
	t.Cleanup(b.Close)
	if h == nil {
		return b.NewWindow()
	} else {
		return b.OpenWindow("https://example.com/index.html")
	}
}
