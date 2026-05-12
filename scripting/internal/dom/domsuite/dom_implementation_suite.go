package domsuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

func testDomImplementation(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e)

	t.Run("createHTMLDocument", func(t *testing.T) {
		win.MustRun(`
			const domImplementation = document.implementation
			const htmldoc = domImplementation.createHTMLDocument()
		`)
		assert.Equal(
			t,
			true,
			win.MustEval(`htmldoc instanceof HTMLDocument`),
			"Document instanceof HTMLDocument",
		)

		assert.Equal(
			t,
			"<html><head></head><body></body></html>",
			win.MustEval(`htmldoc.documentElement.outerHTML`),
			"Document element HTML",
		)

		assert.Equal(
			t,
			nil,
			win.MustEval(`domImplementation.createHTMLDocument().querySelector("title")`),
		)

		assert.Equal(
			t,
			"Test Doc",
			win.MustEval(
				`domImplementation.createHTMLDocument("Test Doc").querySelector("title").textContent`,
			),
		)
	})
}
