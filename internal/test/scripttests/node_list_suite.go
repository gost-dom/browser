package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type NodeListSuite struct {
	ScriptHostSuite
}

func NewNodeListSuite(h html.ScriptHost) *NodeListSuite {
	return &NodeListSuite{ScriptHostSuite: ScriptHostSuite{scriptHost: h}}
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
