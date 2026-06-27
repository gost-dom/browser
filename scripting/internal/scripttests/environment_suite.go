package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

// testEnvironment covers the functional browser-environment surface:
// performance, crypto and the live document.readyState.
func testEnvironment(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e)
	eq := func(name, script string, want any) {
		t.Helper()
		assert.Equal(t, want, win.MustEval(script), name)
	}

	t.Run("performance", func(t *testing.T) {
		eq("now is number", `typeof performance.now() === "number"`, true)
		eq("now monotonic",
			`(() => { const a = performance.now(); const b = performance.now(); return b >= a; })()`, true)
		eq("timeOrigin positive",
			`typeof performance.timeOrigin === "number" && performance.timeOrigin > 0`, true)
	})

	t.Run("crypto", func(t *testing.T) {
		eq("getRandomValues returns same array",
			`(() => { const a = new Uint8Array(16); return crypto.getRandomValues(a) === a; })()`, true)
		eq("getRandomValues fills",
			`(() => { const a = new Uint8Array(32); crypto.getRandomValues(a); return a.some(x => x !== 0); })()`, true)
		eq("randomUUID format",
			`/^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/.test(crypto.randomUUID())`, true)
		eq("rejects Float32Array",
			`(() => { try { crypto.getRandomValues(new Float32Array(4)); return false; } catch (e) { return true; } })()`, true)
	})

	t.Run("document.readyState", func(t *testing.T) {
		eq("is a string", `typeof document.readyState === "string"`, true)
		// A freshly parsed document reports "complete".
		w := browsertest.InitWindow(t, e)
		assert.NoError(t, w.LoadHTML(`<html><body></body></html>`))
		assert.Equal(t, "complete", w.MustEval(`document.readyState`), "readyState after load")
	})
}
