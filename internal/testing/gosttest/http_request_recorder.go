package gosttest

import (
	"net/http"
	"testing"
)

// HTTPRequestRecorder is an HTTPHandler middleware that keeps a record of all
// incoming request objects.
type HTTPRequestRecorder struct {
	T        testing.TB
	Handler  http.Handler
	Requests []*http.Request
}

func NewHTTPRequestRecorder(t testing.TB, handler http.Handler) *HTTPRequestRecorder {
	return &HTTPRequestRecorder{T: t, Handler: handler}
}

func (rec *HTTPRequestRecorder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec.Requests = append(rec.Requests, r)
	rec.Handler.ServeHTTP(w, r)
}

// URLs return all URL strings recorded
func (r HTTPRequestRecorder) URLs() []string {
	res := make([]string, len(r.Requests))
	for i, req := range r.Requests {
		res[i] = req.URL.String()
	}
	return res
}

// Clear deletes all recorded Requests.
func (r *HTTPRequestRecorder) Clear() { r.Requests = nil }

func (r *HTTPRequestRecorder) Single() *http.Request {
	r.T.Helper()
	if l := len(r.Requests); l != 1 {
		r.T.Errorf("HTTPRequestRecorder: expected single request. No of requests: %d", l)
		return nil
	}
	return r.Requests[0]
}
