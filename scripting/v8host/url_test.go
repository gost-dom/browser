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
	assert.Nil(t, win.MustEval(`p2.get("baz")`))

	assert.Equal(t, true, win.MustEval(`
		const x = p2.urlSearchParams
		const y = p2.urlSearchParams
		x === y
	`))

	assert.Equal(t, `f,foo,b,bar`, win.MustEval(`
		{
			const events = []
			console.log("constructor", p2.__proto__.constructor.name)
			for (const [k,v] of p2) {
				events.push(k);
				events.push(v)
			}

			// p2.forEach((k, v) => { events.push(k); events.push(v) })
			events.join(",")
		}
	`))
}
