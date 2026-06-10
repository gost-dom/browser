// Package browseroptions contains more "advanced" options to pass to the browser.
//
// These were extracted to a separate package to avoid adding unnecessary dependencies to the root browser package.
package browseroptions

import (
	"net/http"
	"time"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/internal/fetch"
)

// FetchRoundtripOptions describes properties for an individual fetch request.
type FetchRoundtripOptions = fetch.RoundtripOptions
type InitFetchRoundtripOptionsFunc = fetch.InitRoundTripOptionsFunc

func FetchRequestOptions(f InitFetchRoundtripOptionsFunc) browser.BrowserOption {
	return browser.WithComponentType[fetch.InitRoundTripOptionsFunc](
		func(r *http.Request, o *fetch.RoundtripOptions) {
			f(r, (*FetchRoundtripOptions)(o))
		},
	)
}

func FetchDelay(d time.Duration) browser.BrowserOption {
	return browser.WithComponentType[fetch.InitRoundTripOptionsFunc](
		func(r *http.Request, o *fetch.RoundtripOptions) {
			o.Delay = d
		},
	)
}

// SetDefaultFetchDelay sets the
//
// Note: This is a global default, and should only ever be set in a init
// function; and preferably not at all. If the delay affects the outcome of a
// test, the test should explicitly control that delay.
func SetDefaultFetchDelay(d time.Duration) { fetch.SetDefaultDelay(d) }
