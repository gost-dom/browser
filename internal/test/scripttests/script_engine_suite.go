package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/stretchr/testify/assert"
)

func testScriptEngineBehaviour(t *testing.T, e html.ScriptEngine) {
	t.Run("Assigning a read-only attribute", func(t *testing.T) {
		win := initWindow(t, e, nil)
		res := win.MustEval(`
			const doc = window.document
			window.document = null
			const res = doc === window.document
			res.toString()
		`)
		assert.Equal(t, "true", res, "Value was unaffected by assignment")
	})
}
