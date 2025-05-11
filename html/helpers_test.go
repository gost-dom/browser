package html_test

import (
	"net/http"
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/gosthttp"
	. "github.com/gost-dom/browser/internal/gosthttp"
)

func ParseHtmlString(s string) (res dom.Document) {
	html.NewDOMParser().ParseReader(nil, &res, strings.NewReader(s))
	return
}

func NewWindowFromHandler(handler http.Handler) html.Window {
	return html.NewWindow(
		html.WindowOptions{HttpClient: gosthttp.NewHttpClientFromHandler(handler)},
	)
}
func windowOptionHandler(h http.Handler) html.WindowOption {
	return html.WindowOptionFunc(func(o *html.WindowOptions) {
		o.HttpClient = NewHttpClientFromHandler(h)
	})
}

// A simple http Handler that just calls ParseForm, making the form data
// available on the http request object. When combined with
// [HTTPRequestRecorder], test code can easily inspect the submitted form
// values.
var ParseFormHandler http.Handler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
	})
