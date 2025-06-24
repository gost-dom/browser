package scripttests

import (
	"errors"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/onsi/gomega/types"
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

func (s *ElementSuite) TestIDLInterfaceNamesForElements() {
	ctx := s.Window.ScriptContext()
	s.Expect("document.createElement('a')").To(BeJSInstanceOf("HTMLAnchorElement", ctx))
	s.Expect("document.createElement('p')").To(BeJSInstanceOf("HTMLParagraphElement", ctx))
	s.Expect("document.createElement('div')").To(BeJSInstanceOf("HTMLDivElement", ctx))

}

type BeJSInstanceOfMatcher struct {
	class string
	ctx   html.ScriptContext
}

func BeJSInstanceOf(
	expected string,
	ctx html.ScriptContext,
) types.GomegaMatcher {
	return BeJSInstanceOfMatcher{expected, ctx}
}

func (m BeJSInstanceOfMatcher) Match(actual interface{}) (success bool, err error) {
	str, ok := actual.(string)
	if !ok {
		return false, errors.New("Actual must be a string")
	}
	v, err := m.ctx.Eval(str + " instanceof " + m.class)
	success, ok = v.(bool)
	if !ok {
		panic("Should have received a bool")
	}
	return
}

func (m BeJSInstanceOfMatcher) FailureMessage(actual any) string {
	return "Expected an instance of " + m.class
}
func (m BeJSInstanceOfMatcher) NegatedFailureMessage(actual any) string {
	return "Expected to not be an instance of " + m.class
}

var sampleHTML = `<body>
	<a ref="#foo" id="a">Link</a>
	<p id="p">Paragraph</p>
	<div id="div">Div</div>
</body>`
