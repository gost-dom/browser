package v8engine_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/onsi/gomega"
)

func TestDomParser(t *testing.T) {
	win := browsertest.InitWindow(t, nil)
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
