package scripttests

import (
	"context"
	"testing"

	"github.com/gost-dom/browser"
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

	t.Run("Refuse to start new script when context has been cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(t.Context())

		b := browser.New(browser.WithContext(ctx), browser.WithScriptEngine(e))
		win := b.NewWindow()
		val, err := win.Eval("42")
		assert.NoError(t, err)
		assert.EqualValues(t, 42, val)

		cancel()
		_, err = win.Eval("42")
		assert.Error(t, err)
		assert.ErrorIs(t, err, html.ErrCancelled)
	})
}
