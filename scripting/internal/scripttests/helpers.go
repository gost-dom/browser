package scripttests

import (
	"net/http"
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
)

var initBrowser = browsertest.InitBrowser
var WithLogOption = browsertest.WithLogOption
var WithMinLogLevel = browsertest.WithMinLogLevel

func initWindow(
	t *testing.T, e html.ScriptEngine, h http.Handler, opts ...browsertest.InitOption,
) htmltest.WindowHelper {
	b := htmltest.NewBrowserHelper(t, initBrowser(t, h, e, opts...))
	t.Cleanup(b.Close)
	return b.NewWindow()
}

func openWindow(
	t *testing.T, e html.ScriptEngine, h http.Handler, url string, opts ...browsertest.InitOption,
) htmltest.WindowHelper {
	b := htmltest.NewBrowserHelper(t, initBrowser(t, h, e, opts...))
	t.Cleanup(b.Close)
	return b.OpenWindow(url)
}
