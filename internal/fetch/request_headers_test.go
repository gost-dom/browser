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
// forwarded onto the outgoing HTTP request.
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
			{"Content-Type", "application/json+protobuf"},
			{"X-Goog-Api-Key", "secret-key"},
		}),
	))
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, 200, res.Status)
	req := recorder.Requests[0]
	assert.Equal(t, "application/json+protobuf", req.Header.Get("Content-Type"),
		"Content-Type header forwarded to the outgoing request")
	assert.Equal(t, "secret-key", req.Header.Get("X-Goog-Api-Key"),
		"custom header forwarded to the outgoing request")
}
