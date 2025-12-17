package htmlsuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/onsi/gomega"
)

func testHtmlAnchorElement(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e)

	g := gomega.NewWithT(t)
	g.Expect(win.Eval(`
		const a = document.createElement("a");
		a.href = "http://example.com/foo";
		a.toString()
	`)).To(gomega.Equal("http://example.com/foo"))
}
