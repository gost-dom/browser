package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

// testToStringTag verifies that interface prototypes carry the Web IDL
// @@toStringTag, so Object.prototype.toString.call(obj) yields "[object <Name>]"
// rather than the generic "[object Object]".
func testToStringTag(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e)
	eq := func(name, script string, want any) {
		t.Helper()
		assert.Equal(t, want, win.MustEval(script), name)
	}

	eq("window", `Object.prototype.toString.call(window)`, "[object Window]")
	eq("document", `Object.prototype.toString.call(document)`, "[object HTMLDocument]")
	eq("div element", `Object.prototype.toString.call(document.createElement("div"))`, "[object HTMLDivElement]")
	eq("Headers", `Object.prototype.toString.call(new Headers())`, "[object Headers]")
	// The most-derived class on the prototype chain provides the tag.
	eq("anchor element", `Object.prototype.toString.call(document.createElement("a"))`, "[object HTMLAnchorElement]")
}
