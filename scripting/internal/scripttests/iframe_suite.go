package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

// testIFrame covers HTMLIFrameElement.contentWindow / contentDocument, which
// expose the iframe's nested browsing context.
func testIFrame(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e)
	eq := func(name, script string, want any) {
		t.Helper()
		assert.Equal(t, want, win.MustEval(script), name)
	}

	eq("contentWindow is object",
		`typeof document.createElement("iframe").contentWindow`, "object")
	eq("contentDocument is object",
		`typeof document.createElement("iframe").contentDocument`, "object")
	// Reads off contentWindow resolve to genuine native built-ins of the realm.
	eq("contentWindow.String is native", `(() => {
		const f = document.createElement("iframe");
		const cw = f.contentWindow;
		return cw != null && /\[native code\]/.test(cw.String.toString());
	})()`, true)
}
