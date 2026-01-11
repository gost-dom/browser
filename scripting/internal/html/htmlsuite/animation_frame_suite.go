package htmlsuite

import (
	"log/slog"
	"testing"
	"time"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
)

func testAnimationFrameProvider(t *testing.T, e html.ScriptEngine) {
	t.Run("requestAnimationFrame simulates a small delay", func(t *testing.T) {
		// This isn't super-realistic. If more realism is needed, consider
		// - Configuring a framerate on the browser
		// - Grouping callbacks to simulate synchronized with refresh

		win := browsertest.InitWindow(t, e, browsertest.WithMinLogLevel(slog.LevelInfo))
		win.MustRun(`gost.assertEqual(typeof requestAnimationFrame, "function")`)
		win.MustRun(`
			let called = false
			requestAnimationFrame(() => { 
				called = true 
			})
		`)
		win.Clock().Advance(time.Millisecond)
		win.Assert().False("called")

		win.Clock().Advance(20 * time.Millisecond)
		win.Assert().True("called")
	})
}
