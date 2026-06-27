package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

// testBase64 covers the WindowOrWorkerGlobalScope atob/btoa methods.
func testBase64(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e)
	eq := func(name, script string, want any) {
		t.Helper()
		assert.Equal(t, want, win.MustEval(script), name)
	}

	eq("btoa", `btoa("hello")`, "aGVsbG8=")
	eq("atob", `atob("aGVsbG8=")`, "hello")
	eq("btoa empty", `btoa("")`, "")
	// Round-trips a binary string with bytes outside the printable ASCII range.
	eq("round-trip binary", `atob(btoa("\x00\x01\xfe\xff")) === "\x00\x01\xfe\xff"`, true)
}
