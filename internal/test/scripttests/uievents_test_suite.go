package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type UIEventTestSuite struct {
	ScriptHostSuite
}

func NewUIEventTestSuite(h html.ScriptHost) *UIEventTestSuite {
	return &UIEventTestSuite{ScriptHostSuite{scriptHost: h}}
}

func (s *UIEventTestSuite) TestEventInheritance() {
	s.Assert().
		NoError(s.run("const getSuperclassName = (o) => Object.getPrototypeOf(o.prototype).constructor.name"))
	s.Expect(s.eval(`getSuperclassName(PointerEvent)`)).
		To(Equal("MouseEvent"), "Pointer event superclass")
	s.Expect(s.eval(`getSuperclassName(MouseEvent)`)).
		To(Equal("UIEvent"), "MouseEvent event superclass")
	s.Expect(s.eval(`getSuperclassName(UIEvent)`)).
		To(Equal("Event"), "UIEvent event superclass")
}

func (s *UIEventTestSuite) TestClickEventIsAPointerEvent() {
	s.window.LoadHTML(`<body><div id="foo"></div></body>`)
	s.Assert().NoError(s.run(`
		let event
		document.getElementById("foo").addEventListener("click", e => { event = e })
	`))
	s.window.HTMLDocument().GetHTMLElementById("foo").Click()
	s.Expect(s.eval(`event instanceof PointerEvent`)).To(BeTrue(), "Event is an event")
}
