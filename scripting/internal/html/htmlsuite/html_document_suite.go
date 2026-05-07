package htmlsuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

func testHTMLDocument(t *testing.T, e html.ScriptEngine) {
	t.Parallel()

	win := browsertest.InitWindow(t, e)
	ok := win.MustEval(`
		const e = document.createElement('textarea')
		e instanceof HTMLTextAreaElement
	`)
	assert.Equal(t, ok, true)
}
