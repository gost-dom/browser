package scripttests

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/onsi/gomega"
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

	t.Run("UnhandledPromiseBehaviour", func(t *testing.T) {
		// Promises are pushed to the microtask queue which are executed _after_ the
		// script is evaluated; but before returning from `run`.
		var recorder gosttest.LogRecorder
		g := gomega.NewWithT(t)
		b := browser.New(
			browser.WithScriptEngine(e),
			browser.WithLogger(slog.New(&recorder)))
		defer b.Close()
		win := b.NewWindow()

		g.Expect(win.Eval(`
		let value
		Promise.resolve("dummy").then(val => { value = val }).then(()=> { throw new Error("foo") })
		value
	`)).To(BeNil(), "Promise should not have evaluated yet")
		g.Expect(recorder.FilterLevel(slog.LevelError)).To(gomega.HaveLen(1))

		recorder.Reset()
		g.Expect(win.Run(`
		new Promise((resolve, reject) => {
			reject()
		})
	`)).To(Succeed())
		g.Expect(recorder.FilterLevel(slog.LevelError)).To(gomega.HaveLen(1))

		recorder.Reset()
		g.Expect(win.Run(`
		new Promise((resolve, reject) => { setTimeout(reject, 1000) })
	`)).To(Succeed())

		win.Clock().Advance(time.Millisecond * 999)
		g.Expect(recorder.FilterLevel(slog.LevelError)).To(gomega.HaveLen(0))
		win.Clock().Advance(time.Millisecond * 1)
		g.Expect(recorder.FilterLevel(slog.LevelError)).To(gomega.HaveLen(1))

	})
}
