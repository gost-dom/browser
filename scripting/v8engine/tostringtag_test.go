package v8engine_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/gost-dom/browser/scripting/v8engine"
	"github.com/stretchr/testify/assert"
)

// TestInterfaceToStringTag verifies that interface prototypes carry the Web-IDL
// @@toStringTag, so Object.prototype.toString.call(obj) yields "[object <Name>]"
// rather than the generic "[object Object]".
func TestInterfaceToStringTag(t *testing.T) {
	win := browsertest.InitWindow(t, v8engine.DefaultEngine())
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
