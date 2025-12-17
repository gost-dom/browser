package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"

	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/onsi/gomega"
)

func testDomParser(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e)
	g := gomega.NewWithT(t)
	g.Expect(win.Eval(`
		const parser = new DOMParser()
		const doc = parser.parseFromString("<div id='target'></div>", "text/html")
	`)).Error().ToNot(HaveOccurred())
	g.Expect(
		win.Eval("Object.getPrototypeOf(doc) === HTMLDocument.prototype"),
	).To(BeTrue(), "result is a Document")
	g.Expect(
		win.Eval("doc === window.document"),
	).To(BeFalse(), "Window.document isn't replaced")
	g.Expect(
		win.Eval("doc.getElementById('target') instanceof HTMLDivElement"),
	).To(BeTrue(), "Element is a div")

}
