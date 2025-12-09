package scripttests

import (
	"testing"

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
		console.log("ERR", err?.message)
		console.log("ERR PROTOTYPE", Object.getPrototypeOf(err).constructor.name)
		err instanceof DOMException
	`)
	assert.True(t, result.(bool))
	isError := win.MustEval(`err instanceof Error`)
	assert.True(t, isError.(bool), "err instanceof Error")
}
