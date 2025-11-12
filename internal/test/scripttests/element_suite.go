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

func NewElementSuite(h html.ScriptEngine) *ElementSuite {
	return &ElementSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
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

func (s *ElementSuite) TestAttributes() {
	s.MustLoadHTML(`<body><div foo="foo-value" bar="bar-value"></div><body>`)
	s.MustEval(`
		const div = document.querySelector("div")
		const attrs = Array.from(div.attributes)
		const names = attrs.map(x => x.name).join(",")
		const values = attrs.map(x => x.value).join(",")
	`)
	s.Assert().Equal("foo,bar", s.MustEval("names"))
	s.Assert().Equal("foo-value,bar-value", s.MustEval("values"))
}

func (s *ElementSuite) TestIDLInterfaceNamesForElements() {
	ctx := s.Window.ScriptContext()
	s.Expect("document.createElement('a')").To(BeJSInstanceOf("HTMLAnchorElement", ctx))
	s.Expect("document.createElement('p')").To(BeJSInstanceOf("HTMLParagraphElement", ctx))
	s.Expect("document.createElement('div')").To(BeJSInstanceOf("HTMLDivElement", ctx))
}

func (s *ElementSuite) TestChildren() {
	s.MustLoadHTML(`
		<body>
			<div id="target">
				Initial text
				<div id="child-1">Child 1</div>
				Some text
				<div id="child-2">Child 2</div>
				Final text
			</div>
		</body>`,
	)
	s.MustRunScript(`
		const target = document.getElementById("target")
		const child1ByItem = target.children.item(0)
		const child1ByIndex = target.children[0]
		const child2ByItem = target.children.item(1)
		const child2ByIndex = target.children[1]
		const length = target.children.length
		const arr = Array.from(target.children)
		const contents = arr.map(x => x.getAttribute("id")).join(",")
	`)
	assert := s.Assert()
	assert.Equal("Child 1", s.MustEval("child1ByItem.textContent"))
	assert.Equal("Child 1", s.MustEval("child1ByIndex.textContent"))
	assert.Equal("Child 2", s.MustEval("child2ByItem.textContent"))
	assert.Equal("Child 2", s.MustEval("child2ByIndex.textContent"))

	s.Expect(s.MustEval("length")).To(BeEquivalentTo(2))

	assert.Equal("child-1,child-2", s.MustEval("contents"))
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
