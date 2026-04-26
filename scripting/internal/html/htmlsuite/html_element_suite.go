package htmlsuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

func testHTMLElement(t *testing.T, e html.ScriptEngine) {
	t.Parallel()

	win := browsertest.InitWindow(t, e)
	assert.Equal(t, nil, win.MustEval(`
		const div = document.createElement("div")
		div.className
	`))
	assert.Equal(t, "foo bar", win.MustEval(`
		div.setAttribute("class", "foo bar")
		div.className
	`))
	win.MustRun("div.className = 'baz'")
	assert.Equal(t, "baz", win.MustEval(`div.getAttribute("class")`))
}
