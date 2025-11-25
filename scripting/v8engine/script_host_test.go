package v8engine_test

import (
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/scripting/v8engine"
	"github.com/gost-dom/browser/v8browser"
	"github.com/onsi/gomega"
)

func TestScriptHostDocumentScriptLoading(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
	reader := strings.NewReader(`<html><body>
    <script>window.sut = document.documentElement.outerHTML</script>
    <div>I should not be in the output</div>
  </body></html>
`)
	host := v8engine.New()
	t.Cleanup(host.Close)
	options := html.WindowOptions{ScriptHost: host}
	win, err := html.NewWindowReader(reader, options)
	defer win.Close()
	Expect(err).ToNot(HaveOccurred())
	ctx := win.ScriptContext()
	Expect(
		ctx.Eval("window.sut"),
	).To(Equal(`<html><head></head><body>
    <script>window.sut = document.documentElement.outerHTML</script></body></html>`))
}

func TestScriptHostUnhandledPromiseRejection(t *testing.T) {
	// Promises are pushed to the microtask queue which are executed _after_ the
	// script is evaluated; but before returning from `run`.
	var recorder gosttest.LogRecorder
	g := gomega.NewWithT(t)
	b := v8browser.New(browser.WithLogger(slog.New(&recorder)))
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
}
