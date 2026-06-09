package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func testNode(t *testing.T, e html.ScriptEngine) {
	t.Run("Structure", func(t *testing.T) { testStructure(t, e) })
	t.Run("InsertBefore", func(t *testing.T) { testInsertBefore(t, e) })
	t.Run("InsertBefore with no ref", func(t *testing.T) { testInsertBeforeWithNoRef(t, e) })
	t.Run("InsertBefore with null ref", func(t *testing.T) { testInsertBeforeWithNullRef(t, e) })
	t.Run("Remove child", func(t *testing.T) { testRemoveChild(t, e) })
	t.Run("First child", func(t *testing.T) { testFirstChild(t, e) })
	t.Run("Contains", func(t *testing.T) { testContains(t, e) })
	t.Run("Static properties exist on both constructor and prototype", func(t *testing.T) {
		testStaticProperties(t, e)
	})
}

func testStructure(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e, browsertest.WithHtml(
		`<div id="parent-1"><div id="child-1"></div><div id="child-2"></div></div>`,
	))
	assert.Equal(t, true, win.MustEval(`
		const parent1 = document.getElementById("parent-1")
		const child1 = document.getElementById("child-1")
		child1.parentNode === parent1
	`))
}

func testInsertBefore(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e, browsertest.WithHtml(
		`<div id="parent-1"><div id="child-1"></div><div id="child-2"></div></div>`,
	))
	g := gomega.NewGomegaWithT(t)
	g.Expect(win.Eval(`
		const f = document.createDocumentFragment()
		const d1 = document.createElement("div")
		const d2 = document.createElement("div")
		d1.setAttribute("id", "d1")
		d2.setAttribute("id", "d2")
		f.appendChild(d1)
		f.appendChild(d2)
		const parent = document.getElementById("parent-1")
		ref = document.getElementById("child-2")
		parent.insertBefore(f, ref)
		Array.from(parent.childNodes).map(x => x.getAttribute("id")).join(", ")
	`)).To(Equal("child-1, d1, d2, child-2"))
}

func testInsertBeforeWithNullRef(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e, browsertest.WithHtml(
		`<div id="parent-1"><div id="child-1"></div><div id="child-2"></div></div>`,
	))
	g := gomega.NewGomegaWithT(t)
	g.Expect(win.Eval(`
		const f = document.createDocumentFragment()
		const d1 = document.createElement("div")
		const d2 = document.createElement("div")
		d1.setAttribute("id", "d1")
		d2.setAttribute("id", "d2")
		f.appendChild(d1)
		f.appendChild(d2)
		const parent = document.getElementById("parent-1")
		parent.insertBefore(f, null)
		Array.from(parent.childNodes).map(x => x.getAttribute("id")).join(", ")
	`)).To(Equal("child-1, child-2, d1, d2"))
}

func testInsertBeforeWithNoRef(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e, browsertest.WithHtml(
		`<div id="parent-1"><div id="child-1"></div><div id="child-2"></div></div>`,
	))
	g := gomega.NewWithT(t)
	g.Expect(win.Eval(`
		const f = document.createDocumentFragment()
		const d1 = document.createElement("div")
		const d2 = document.createElement("div")
		d1.setAttribute("id", "d1")
		d2.setAttribute("id", "d2")
		f.appendChild(d1)
		f.appendChild(d2)
		const parent = document.getElementById("parent-1")
		parent.insertBefore(f)
		Array.from(parent.childNodes).map(x => x.getAttribute("id")).join(", ")
	`)).To(Equal("child-1, child-2, d1, d2"))
}

func testRemoveChild(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e, browsertest.WithHtml(
		`<div id="parent-1"><div id="child">child</div></div>`,
	))
	s := gomega.NewGomegaWithT(t)
	s.Expect(win.Run(`
		const child = document.getElementById('child');
		const parent = document.getElementById('parent-1')
		parent.removeChild(child)
	`)).To(Succeed())
	s.Expect(
		win.Document().GetElementById("parent-1").ChildNodes().Length(),
	).To(Equal(0))
}

func testFirstChild(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e, browsertest.WithHtml(
		`<div id="parent-1"><div id="child">child</div></div>`,
	))
	s := gomega.NewGomegaWithT(t)

	s.Expect(
		win.Eval(`document.getElementById("parent-1").firstChild.getAttribute("id")`),
	).To(Equal("child"))
}

func testContains(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e, browsertest.WithHtml(`
		<div>
			<div id="parent-1">
				<div id="child">child</div>
			</div>
			<div id="parent-2"></div>
		</div>
		<script>
			const parent1 = document.getElementById("parent-1")
			const parent2 = document.getElementById("parent-2")
			const child = document.getElementById("child")
		</script>
	`))

	assert.Equal(t, true,
		win.MustEval(`parent1.contains(child)`),
		"node.contains when passed a child element")

	assert.Equal(t, false,
		win.MustEval(`parent1.contains(parent2)`),
		"node.contains when passed a child element")
}

func testStaticProperties(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e)
	win.MustRun("const el = document.documentElement")

	assert.EqualValues(t, 1, win.MustEval("el.ELEMENT_NODE"))
	assert.EqualValues(t, 2, win.MustEval("el.ATTRIBUTE_NODE"))

	assert.EqualValues(t, 1, win.MustEval("Node.ELEMENT_NODE"))
	assert.EqualValues(t, 2, win.MustEval("Node.ATTRIBUTE_NODE"))
	// Not going through them all
}
