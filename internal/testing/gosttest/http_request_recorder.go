package gosttest

import (
	"net/http"
	"testing"
)

// HttpRequestFormRecorder is a very simple http.Handler that just records the
// incoming requests, and returns a 200 status.
//
// The recorder automatically calls [http/Request.ParseForm] to make form data
// availeble.
type HttpRequestFormRecorder struct {
	T        testing.TB
	Requests []*http.Request
}

func (rec *HttpRequestFormRecorder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rec.Requests = append(rec.Requests, r)
}

// Single asserts that a single request was made
func (r HttpRequestFormRecorder) Single() *http.Request {
	r.T.Helper()
	if len(r.Requests) != 1 {
		r.T.Errorf("Expected single recorded request. Got: %d", len(r.Requests))
	}
	return r.Requests[0]
}
