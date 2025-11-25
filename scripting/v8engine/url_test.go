package v8engine_test

import (
	"testing"

	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	win := browsertest.InitWindow(t, nil)
	assert.Equal(t, "http://example.com/foo/bar", win.MustEval(`
		const u = new URL("foo/bar", "http://example.com");
		u.href
	`))
}

func TestURLSearchParamsFromIterable(t *testing.T) {
	win := browsertest.InitWindow(t, nil)
	usp, ok := win.MustEval(`
		{
			const fd = new FormData()
			fd.append("f", "foo")
			fd.append("b", "bar")
			fd.append("b", "baz")
			new URLSearchParams(fd)
		}
	`).(urlinterfaces.URLSearchParams)
	assert.True(t, ok, "Interface was of the expected type")
	got, _ := usp.Get("f")
	assert.Equal(t, "foo", got)
	assert.Equal(t, []string{"bar", "baz"}, usp.GetAll("b"))

	usp = win.MustEval(`
		{
			new URLSearchParams({
				f: "foo",
				b: "bar"
			})
		}
	`).(urlinterfaces.URLSearchParams)
	assert.Equal(t, []string{"foo"}, usp.GetAll("f"))
	assert.Equal(t, []string{"bar"}, usp.GetAll("b"))
}

func TestURLSearchParams(t *testing.T) {
	// This test reflect an implementation that doesn't follow the spec in that
	// query params aren't returned in the order they are specified.
	win := browsertest.InitWindow(t, nil)
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

	assert.ElementsMatch(t, []any{"f,foo", "b,bar"}, win.MustEval(`
		{
			const events = []
			for (const [k,v] of p2) {
				events.push(`+"`${k},${v}`"+`);
			}
			events;
		}
	`))

	assert.ElementsMatch(t, []any{"f,foo", "b,bar"}, win.MustEval(`
		{
			const events = []
			for (const [k,v] of p2.entries()) {
				events.push(`+"`${k},${v}`"+`);
			}
			events
		}
	`))

	assert.ElementsMatch(t, []any{"f", "b"}, win.MustEval(`
		{
			const events = []
			for (const k of p2.keys()) {
				events.push(k);
			}
			events
		}
	`))

	assert.ElementsMatch(t, []any{"foo", "bar"}, win.MustEval(`
		{
			const events = []
			for (const k of p2.values()) {
				events.push(k);
			}
			events
		}
	`))
}
