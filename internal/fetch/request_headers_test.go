package fetch_test

import (
	"net/http"
	"testing"

	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/types"
	"github.com/stretchr/testify/assert"
)

// TestFetchForwardsHeaders verifies that headers set on a fetch Request are
// forwarded onto the request sent by the http.Client to the round tripper.
func TestFetchForwardsHeaders(t *testing.T) {
	recorder := gosttest.NewHTTPRequestRecorder(t,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) }),
	)
	bc, cancel := gosttest.NewBrowsingContextFromT(t, recorder, gosttest.WithTimeoutMS(100))
	defer cancel()

	f := fetch.New(bc)
	res, err := f.Fetch(f.NewRequest("url",
		fetch.WithMethod("POST"),
		fetch.WithHeaders([][2]types.ByteString{
			{"Content-Type", "application/json"},
			{"X-Example-Header", "example-value"},
		}),
	))
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, 200, res.Status)
	req := recorder.Single()
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"),
		"Content-Type header forwarded to the outgoing request")
	assert.Equal(t, "example-value", req.Header.Get("X-Example-Header"),
		"custom header forwarded to the outgoing request")
}

// TestFetchDoesNotForwardCookieHeaders verifies that cookie-related headers are
// not forwarded onto the outgoing request. They are forbidden request headers;
// cookies on real requests are managed by the http.Client's cookie jar, not by
// copying header values.
func TestFetchDoesNotForwardCookieHeaders(t *testing.T) {
	recorder := gosttest.NewHTTPRequestRecorder(t,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) }),
	)
	bc, cancel := gosttest.NewBrowsingContextFromT(t, recorder, gosttest.WithTimeoutMS(100))
	defer cancel()

	f := fetch.New(bc)
	_, err := f.Fetch(f.NewRequest("url",
		fetch.WithHeaders([][2]types.ByteString{
			{"Cookie", "session=example"},
			{"Set-Cookie", "session=example"},
		}),
	))
	if !assert.NoError(t, err) {
		return
	}
	req := recorder.Single()
	assert.Empty(t, req.Header.Values("Cookie"),
		"Cookie is a forbidden request header and must not be forwarded")
	assert.Empty(t, req.Header.Values("Set-Cookie"),
		"Set-Cookie is a forbidden request header and must not be forwarded")
}
