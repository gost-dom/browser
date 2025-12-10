package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/stretchr/testify/assert"
)

func testErrorHandling(t *testing.T, e html.ScriptEngine) {
	win := initWindow(t, e, nil)
	result := win.MustEval(`
		let err
		try {
			const newDoc = new Document()
			document.documentElement.append(newDoc)
		} catch(e) {
			err = e
		}
		err instanceof DOMException
	`)

	assert.True(t, result.(bool))
	isError := win.MustEval(`err instanceof Error`)
	assert.True(t, isError.(bool), "err instanceof Error")

	expectedErr := win.Document().Append(dom.NewDocument(win))
	assert.Equal(t, expectedErr.Error(), win.MustEval("err.message"))
}
