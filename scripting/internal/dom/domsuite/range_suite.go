package domsuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

func testRange(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e, browsertest.WithHtml(
		`<body><h1>Hello</h1></body>`))

	win.MustRun(`const range = document.createRange()`)
	doc := win.Document()
	assert.EqualValues(t, 0, win.MustEval(`range.startOffset`))
	assert.EqualValues(t, 0, win.MustEval(`range.endOffset`))
	assert.Equal(t, doc, win.MustEval(`range.startContainer`))
	assert.Equal(t, doc, win.MustEval(`range.endContainer`))
}
