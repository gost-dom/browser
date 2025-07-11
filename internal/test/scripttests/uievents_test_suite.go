package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/uievents"
)

type UIEventTestSuite struct {
	ScriptHostSuite
}

func NewUIEventTestSuite(h html.ScriptHost) *UIEventTestSuite {
	return &UIEventTestSuite{ScriptHostSuite{scriptHost: h}}
}

func (s *UIEventTestSuite) TestEventInheritance() {
	s.Assert().
		NoError(s.RunScript("const getSuperclassName = (o) => Object.getPrototypeOf(o.prototype).constructor.name"))
	s.Expect(s.Eval(`getSuperclassName(PointerEvent)`)).
		To(Equal("MouseEvent"), "Pointer event superclass")
	s.Expect(s.Eval(`getSuperclassName(MouseEvent)`)).
		To(Equal("UIEvent"), "MouseEvent event superclass")
	s.Expect(s.Eval(`getSuperclassName(UIEvent)`)).
		To(Equal("Event"), "UIEvent event superclass")
	s.Expect(s.Eval(`getSuperclassName(KeyboardEvent)`)).
		To(Equal("UIEvent"), "KeyboardEvent event superclass")
}

func (s *UIEventTestSuite) TestKeyboardEvent() {
	s.Window.LoadHTML(`<body><div id="foo"></div></body>`)
	s.MustRunScript(`
		let event
		document.getElementById("foo").addEventListener("keydown", e => { event = e })
	`)

	uievents.KeydownInit(
		s.Window.HTMLDocument().GetHTMLElementById("foo"),
		uievents.KeyboardEventInit{Key: "a"},
	)

	s.Expect(s.Eval(`event instanceof KeyboardEvent`)).To(BeTrue(), "Event is a KeyboardEvent")
	s.Expect(s.Eval(`event.key`)).To(Equal("a"), "Event has key: 'a'")
}

func (s *UIEventTestSuite) TestClickEventIsAPointerEvent() {
	s.Window.LoadHTML(`<body><div id="foo"></div></body>`)
	s.MustRunScript(`
		let event
		document.getElementById("foo").addEventListener("click", e => { event = e })
	`)
	s.Window.HTMLDocument().GetHTMLElementById("foo").Click()
	s.Expect(s.Eval(`event instanceof PointerEvent`)).To(BeTrue(), "Event is an event")
}
