package fetch_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/fetch"
	"github.com/stretchr/testify/assert"
)

func TestResponseOk(t *testing.T) {
	for _, tc := range []struct {
		status int
		want   bool
	}{{199, false}, {200, true}, {204, true}, {299, true}, {300, false}, {404, false}, {500, false}} {
		assert.Equalf(t, tc.want, (fetch.Response{Status: tc.status}).Ok(),
			"Ok() for status %d", tc.status)
	}
}

func TestResponseStatusText(t *testing.T) {
	// With no originating http.Response, StatusText falls back to the standard
	// reason phrase for the status code.
	assert.Equal(t, "OK", (fetch.Response{Status: 200}).StatusText())
	assert.Equal(t, "Not Found", (fetch.Response{Status: 404}).StatusText())
}
