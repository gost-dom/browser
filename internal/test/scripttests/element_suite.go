package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type ElementSuite struct {
	ScriptHostSuite
}

func NewElementSuite(h html.ScriptHost) *ElementSuite {
	return &ElementSuite{ScriptHostSuite: ScriptHostSuite{scriptHost: h}}
}

func (s *ElementSuite) TestAppendMultipleElements() {
	s.Expect(s.Eval(`
		const d = document.createElement("div")
		d.append(
			document.createElement("p"),
			document.createElement("p"),
		);
		d.outerHTML`)).To(Equal("<div><p></p><p></p></div>"))

	s.Expect(s.Eval(`
		const d2 = document.createElement("div")
		d2.append(
			document.createElement("p"),
			"Foo",
			"bar",
			document.createElement("p"),
		);
		d2.outerHTML`)).To(Equal("<div><p></p>Foobar<p></p></div>"))
}
