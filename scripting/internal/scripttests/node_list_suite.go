package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type NodeListSuite struct {
	ScriptHostSuite
}

func NewNodeListSuite(h html.ScriptEngine) *NodeListSuite {
	return &NodeListSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *NodeListSuite) TestExtendsObject() {
	s.Expect(s.Eval(
		`Object.getPrototypeOf(NodeList.prototype) === Object.prototype`,
	)).To(BeTrue())
}

func (s *NodeListSuite) TestNodeListWithThreeElements() {
	s.MustLoadHTML(`<div id="1"></div><div id="2"></div><div id="3"></div>`)

	s.Expect(s.Eval(`
		const a = Array.from(document.body.childNodes);
		a.map(x => x.getAttribute("id")).join(",")
	`)).To(Equal("1,2,3"), "Node list is iterable")

	s.Expect(s.Eval("document.body.childNodes.length")).To(BeEquivalentTo(3), "NodeList.length")

	s.Expect(s.Eval(
		`document.body.childNodes.item(1).getAttribute("id")`,
	)).To(Equal("2"), "NodeList.item")

	s.Expect(s.Eval(
		`document.body.childNodes[1].getAttribute("id")`,
	)).To(Equal("2"), "NodeList indexed property handler")

	s.Expect(s.Eval(
		`document.body.childNodes.item(5) === null`,
	)).To(BeTrue(), "Get item out of range")

	s.Expect(s.Eval(
		`document.body.childNodes[5] === undefined`,
	)).To(BeTrue())
}

func (s *NodeListSuite) TestNodeListIteration() {
	s.MustLoadHTML(`<div id="1">Node 1</div><div id="2">Node 2</div><div id="3">Node 3</div>`)
	s.MustRunScript(`
		let nodes = []
		for (const node of document.body.childNodes) { 
			nodes.push(node) 
		}
	`)
	s.Expect(s.Eval("nodes.length")).To(BeEquivalentTo(3), "Array has three elements")
	s.Expect(s.Eval("nodes[0].textContent")).To(Equal("Node 1"))
	s.Expect(s.Eval("nodes[1].textContent")).To(Equal("Node 2"))
	s.Expect(s.Eval("nodes[2].textContent")).To(Equal("Node 3"))
	s.Expect(s.Eval("typeof nodes[3]")).To(Equal("undefined"))

	s.Expect(s.Eval("Array.from(document.body.childNodes.keys()).join(',')")).To(Equal("0,1,2"))
	s.Expect(s.Eval("Array.from(document.body.childNodes.values()).map(x => x.textContent).join(',')")).
		To(Equal("Node 1,Node 2,Node 3"))
}

func (s *NodeListSuite) TestNodeListForEach() {
	s.MustLoadHTML(`<div id="1"></div><div id="2"></div><div id="3"></div>`)
	s.MustRunScript(`
		let entries = []
		let thisPtr
		document.body.childNodes.forEach(function (value, key, collection) {
			thisPtr = this
			entries.push([value,key,collection])
		})
	`)
	s.Expect(s.Eval("thisPtr === globalThis")).To(BeTrue(), "'this' is the global scope")
	s.Expect(s.Eval("entries.length")).To(BeEquivalentTo(3), "Array has three elements")
	s.Expect(s.Eval("entries[0][0] instanceof HTMLElement")).
		To(BeTrue(), "First argument is the value")
	s.Expect(s.Eval("entries[0][1]")).To(BeEquivalentTo(0), "Second argument is the index")
	s.Expect(s.Eval("entries[0][2] === document.body.childNodes")).
		To(BeTrue(), "Third argument is the collection")
	s.Expect(s.Eval("entries[1][1]")).To(BeEquivalentTo(1), "Index of second entry")
	s.Expect(s.Eval("entries[2][1]")).To(BeEquivalentTo(2), "Index of last entry")
}
func (s *NodeListSuite) TestNodeListEntriesIteration() {
	s.MustLoadHTML(`<div id="1"></div><div id="2"></div><div id="3"></div>`)
	s.MustRunScript(`
		let entries = []
		for (const e of document.body.childNodes.entries()) { 
			entries.push(e) 
		}
	`)
	s.Expect(s.Eval("entries.length")).To(BeEquivalentTo(3), "Array has three elements")
	s.Expect(s.Eval("entries[0][0]")).To(BeEquivalentTo(0), "First element is the index")
	s.Expect(s.Eval("entries[0][1] instanceof HTMLElement")).To(BeTrue())
	s.Expect(s.Eval("entries[1][0]")).To(BeEquivalentTo(1), "First element is the index")
	s.Expect(s.Eval("entries[1][1] instanceof HTMLElement")).To(BeTrue())
}
