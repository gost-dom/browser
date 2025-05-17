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

func TestURLSearchParams(t *testing.T) {
	win := initWindow(t)
	assert.Equal(t, "value", win.MustEval(`
		{
			const p = new URLSearchParams()
			p.append("key", "value")
			p.get("key")
		}
	`))

	win.MustRun(`const p2 = new URLSearchParams("?f=foo&b=bar")`)
	assert.Equal(t, "foo", win.MustEval(`p2.get("f")`))
	assert.Equal(t, "bar", win.MustEval(`p2.get("b")`))
	assert.Nil(t, win.MustEval(`p2.get("baz")`), "Reading non-existing value")

	assert.Equal(t, true, win.MustEval(`
		const x = p2.urlSearchParams
		const y = p2.urlSearchParams
		x === y
	`))
}
