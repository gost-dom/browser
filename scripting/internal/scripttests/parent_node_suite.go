package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/onsi/gomega"
)

func testParentNode(t *testing.T, e html.ScriptEngine) {
	t.Run("querySelector are functions", func(t *testing.T) {
		b := initBrowser(t, nil, e)
		w := b.NewWindow()
		g := gomega.NewWithT(t)
		g.Expect(w.Eval(`typeof document.querySelector`)).
			To(Equal("function"), "document.querySelector is a function")
		g.Expect(w.Eval(`typeof document.querySelectorAll`)).
			To(Equal("function"), "document.querySelectorAll is a function")
	})

	t.Run("querySelectorAll is iterable", func(t *testing.T) {
		w := initWindow(t, e, nil)
		g := gomega.NewWithT(t)
		g.Expect(w.LoadHTML(
			`<div id="1" class="foo"><p id="child1"></p><div id="child2"></div><p id="child3"></p></div>`,
		)).To(Succeed())
		g.Expect(w.Eval(`
			const e = document.getElementById("1")
			Array.from(e.querySelectorAll("p")).map(x => x.getAttribute("id")).join(",")
		`)).To(Equal("child1,child3"))

		g.Expect(w.Eval(`
			const res = []
			for(let {id: i} of e.querySelectorAll("p")) {
				res.push(i)
			}
			res.join(",")
		`)).To(Equal("child1,child3"))
	})
}
