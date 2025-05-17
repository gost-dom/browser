package v8host_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	win := initWindow(t)
	assert.Equal(t, "http://example.com/foo/bar", win.MustEval(`
		const u = new URL("foo/bar", "http://example.com");
		u.href
	`))
}
