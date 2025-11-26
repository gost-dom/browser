package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

func testCustomEvent(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e)
	assert.Equal(t, "foo", win.MustEval(`
		const e = new CustomEvent("foo", { detail: { f: "foo", b: "bar" }})
		const d = e.detail
		d.f
	`))
}
