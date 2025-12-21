package html_test

import (
	"net/http"
	// . "github.com/gost-dom/browser/internal/gosthttp"
)

// func NewWindowFromHandler(handler http.Handler, loc string) html.Window {
// 	b := browser.New(browser.WithHandler(handler))
// 	win, err := b.Open(loc)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return win
// }
//
// func windowOptionsBaseUrl(u string) html.WindowOption {
// 	return html.WindowOptionFunc(func(o *html.WindowOptions) {
// 		o.BaseLocation = u
// 	})
// }
//
// func windowOptionHandler(h http.Handler) html.WindowOption {
// 	return html.WindowOptionFunc(func(o *html.WindowOptions) {
// 		o.HttpClient = NewHttpClientFromHandler(h)
// 	})
// }

// A simple http Handler that just calls ParseForm, making the form data
// available on the http request object. When combined with
// [HTTPRequestRecorder], test code can easily inspect the submitted form
// values.
var ParseFormHandler http.Handler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) { r.ParseForm() },
)
