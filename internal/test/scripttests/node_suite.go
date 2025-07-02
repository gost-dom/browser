package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type NodeTestSuite struct {
	ScriptHostSuite
}

func NewNodeTestSuite(h html.ScriptHost) *NodeTestSuite {
	return &NodeTestSuite{ScriptHostSuite: ScriptHostSuite{scriptHost: h}}
}

func (s *NodeTestSuite) TestInsertBefore() {
	s.MustLoadHTML(`<div id="parent-1"><div id="child-1"></div><div id="child-2"></div></div>`)
	s.Expect(s.Eval(`
		const f = document.createDocumentFragment()
		const d1 = document.createElement("div")
		const d2 = document.createElement("div")
		d1.setAttribute("id", "d1")
		d2.setAttribute("id", "d2")
		f.appendChild(d1)
		f.appendChild(d2)
		parent = document.getElementById("parent-1")
		ref = document.getElementById("child-2")
		parent.insertBefore(f, ref)
		Array.from(parent.childNodes).map(x => x.getAttribute("id")).join(", ")
	`)).To(Equal("child-1, d1, d2, child-2"))
}

func (s *NodeTestSuite) TestInsertBeforeWithNullRef() {
	s.MustLoadHTML(`<div id="parent-1"><div id="child-1"></div><div id="child-2"></div></div>`)
	s.Expect(s.Eval(`
		const f = document.createDocumentFragment()
		const d1 = document.createElement("div")
		const d2 = document.createElement("div")
		d1.setAttribute("id", "d1")
		d2.setAttribute("id", "d2")
		f.appendChild(d1)
		f.appendChild(d2)
		parent = document.getElementById("parent-1")
		parent.insertBefore(f, null)
		Array.from(parent.childNodes).map(x => x.getAttribute("id")).join(", ")
	`)).To(Equal("child-1, child-2, d1, d2"))
}

func (s *NodeTestSuite) TestInsertBeforeWithNoRef() {
	s.MustLoadHTML(`<div id="parent-1"><div id="child-1"></div><div id="child-2"></div></div>`)
	s.Expect(s.Eval(`
		const f = document.createDocumentFragment()
		const d1 = document.createElement("div")
		const d2 = document.createElement("div")
		d1.setAttribute("id", "d1")
		d2.setAttribute("id", "d2")
		f.appendChild(d1)
		f.appendChild(d2)
		parent = document.getElementById("parent-1")
		parent.insertBefore(f)
		Array.from(parent.childNodes).map(x => x.getAttribute("id")).join(", ")
	`)).To(Equal("child-1, child-2, d1, d2"))
}

func (s *NodeTestSuite) TestRemoveChild() {
	s.MustLoadHTML(`<div id="parent-1"><div id="child">child</div></div>`)
	s.Expect(s.RunScript(`
		const child = document.getElementById('child');
		const parent = document.getElementById('parent-1')
		parent.removeChild(child)
	`)).To(Succeed())
	s.Expect(
		s.Window.Document().GetElementById("parent-1").ChildNodes().Length(),
	).To(Equal(0))
}

func (s *NodeTestSuite) TestFirstChild() {
	s.MustLoadHTML(`<div id="parent-1"><div id="child">child</div></div>`)
	s.Expect(
		s.Eval(`document.getElementById("parent-1").firstChild.getAttribute("id")`),
	).To(Equal("child"))
}

func (s *NodeTestSuite) TestContains() {
	s.MustLoadHTML(`
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
	`)

	s.Expect(s.Eval(`parent1.contains(child)`)).To(BeTrue(),
		"node.contains when passed a child element")

	s.Expect(s.Eval(`parent1.contains(parent2)`)).To(BeFalse(),
		"node.contains when passed a child element")
}
