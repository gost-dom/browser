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

func (s *UIEventTestSuite) TestEventIsInstanceOfEvent() {
	s.Expect(s.eval(`new Event("foo") instanceof Event`)).To(BeTrue())
}

func (s *UIEventTestSuite) TestClickEventInheritance() {
	s.window.LoadHTML(`<body><div id="foo"></div></body>`)
	s.Assert().NoError(s.run(`
		let event
		document.getElementById("foo").addEventListener("click", e => { event = e })
	`))
	s.window.Document().GetElementById("foo").Click()
	s.Expect(s.eval(`event instanceof Event`)).To(BeTrue())
}
