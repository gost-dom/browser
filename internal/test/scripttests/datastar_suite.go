package scripttests

import (
	"context"
	"testing"
	"time"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/input/controller"
	"github.com/gost-dom/browser/input/key"
	app "github.com/gost-dom/browser/internal/test/integration/test-app"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/gost-dom/browser/testing/gosttest"
	"github.com/stretchr/testify/assert"
)

func RunDataStarTests(t *testing.T, e html.ScriptEngine) {
	t.Run("Simple datastar test", func(t *testing.T) {
		// Each test has it's own V8 instance, so everything can be parallelized.
		t.Parallel()

		// Make the test abort if it hasn't completed withing a second. This is an
		// excessive timeout - the test runs in about a millisecond, but
		// bootstrapping new v8 isolates, when one cannot be picked from a pool
		// incurs an overhead.
		ctx, cancel := context.WithTimeout(t.Context(), time.Second)
		defer cancel()

		// The BrowserHelper is an internal test helper in Gost-DOM, and may
		// eventually make it's way out of the "internal" space. The patterns are
		// easily adopted locally.
		//
		// It adds testing.TB-aware helpers on top of native functions, e.g, non-nil
		// error return values will be reported to t.Error(), automatically causing
		// the test to fail.
		b := htmltest.NewBrowserHelper(t,
			// browser.New is the primary entry point to creating a Gost-DOM browser
			browser.New(
				browser.WithScriptEngine(e),
				// By passing a context, the browser automatically disposes
				// resources when the context cancels. By deriving contexts from
				// t.Context(), the context will automatically cancel when the test
				// is done.
				//
				// Disposing the browser will allow V8 isolates to be reused,
				// significantly speeding up creation of new browsers.
				browser.WithContext(ctx),
				// app.CreateServer returns the root HTTP handler for the test
				// application. browser.WithHandler connects Gost-DOM directly to the
				// http Handler, bypassing the TCP transport layer, and elliminating
				// the need for starting/stopping HTTP servers, and managing ports -
				// as well as facilitating stubbing components.
				browser.WithHandler(app.CreateServer()),
				// WithLogger installs a log handler Gost-DOM logs to an *slog.Logger
				// from the standard library.
				browser.WithLogger(
					// The test logger pipes all log messages to t.Log() methods.
					// Error level logs are piped to t.Error() though (this is
					// configurable). This will cause tests to fail if a JavaScript
					// error is thrown, even if the assertions in Go code succeed.
					gosttest.NewTestingLogger(t), // You don't need this. Logs all JS->Go calls
					// gosttest.MinLogLevel(slog.LevelDebug),

				),
			))
		// The host name is ignored, but the server serves a Datastar test page on
		// the /ds/index.html route
		win := b.OpenWindow("https://example.com/ds/")
		doc := win.HTMLDocument() // Wrap Window.Document() and returns a "Document test helper"

		// GetHTMLElementById wraps GetElementById, but asserts that any non-nil
		// return values are valid HTMLElement instances, providing access to
		// methods such as Click() that don't exist on the Element.
		clickTarget := doc.GetHTMLElementById("click-target")

		// Verify textContent both before and after clicking.
		assert.Equal(t, "", clickTarget.TextContent())
		doc.GetHTMLElementById("fetch-events-button").Click()
		assert.NoError(t, win.Clock().ProcessEvents(ctx)) // Wait for pending promises to settle.
		assert.Equal(t, "Foobar", clickTarget.TextContent())
	})

	t.Run("Signals", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(t.Context(), time.Second)
		defer cancel()

		b := htmltest.NewBrowserHelper(t,
			browser.New(
				browser.WithScriptEngine(e),
				browser.WithContext(ctx),
				browser.WithHandler(app.CreateServer()),
				browser.WithLogger(
					gosttest.NewTestingLogger(t), // gosttest.MinLogLevel(slog.LevelDebug),

				),
			))
		win := b.OpenWindow("https://example.com/ds/")
		win.HTMLDocument().GetHTMLElementById("echo-input-field").Focus()
		ctrl := controller.KeyboardController{Window: win}
		ctrl.SendKey(key.RuneToKey('a'))
		ctrl.SendKey(key.RuneToKey('b'))
		ctrl.SendKey(key.RuneToKey('c'))
		win.Clock().RunAll()

		output := win.HTMLDocument().GetHTMLElementById("echo-output")
		assert.Equal(t, "abc", output.TextContent())

	})
}
