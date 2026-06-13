package v8engine_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/gost-dom/browser/scripting/v8engine"
	"github.com/stretchr/testify/assert"
)

// TestBase64 covers the WindowOrWorkerGlobalScope atob/btoa methods. V8 does not
// provide these (they are Web APIs, not part of ECMAScript), so they are
// implemented natively.
func TestBase64(t *testing.T) {
	win := browsertest.InitWindow(t, v8engine.DefaultEngine())
	eq := func(name, script string, want any) {
		t.Helper()
		assert.Equal(t, want, win.MustEval(script), name)
	}

	eq("btoa", `btoa("hello")`, "aGVsbG8=")
	eq("atob", `atob("aGVsbG8=")`, "hello")
	eq("btoa empty", `btoa("")`, "")
	// Round-trips a binary string with bytes outside the printable ASCII range.
	eq("round-trip binary", `atob(btoa("\x00\x01\xfe\xff")) === "\x00\x01\xfe\xff"`, true)
	// Implemented natively, so Function.prototype.toString reports native code.
	eq("atob is native", `/\[native code\]/.test(atob.toString())`, true)
	eq("btoa is native", `/\[native code\]/.test(btoa.toString())`, true)
}
