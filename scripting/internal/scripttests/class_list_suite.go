package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/stretchr/testify/assert"
)

type ClassListTestSuite struct {
	ScriptHostSuite
}

func NewClassListTestSuite(h html.ScriptEngine) *ClassListTestSuite {
	return &ClassListTestSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *ClassListTestSuite) TestAdd() {
	s.Window.MustLoadHTML("<div id='target' class='c1'></div>")
	assert.NoError(s.T(), s.Window.Run(`
			const list = document.getElementById('target').classList;
			list.add('c2')
		`))
	element := s.Window.HTMLDocument().GetHTMLElementById("target")
	s.Expect(element).To(HaveAttribute("class", "c1 c2"))
}

func (s *ClassListTestSuite) TestClassListIsIterable() {
	s.Window.MustLoadHTML("<div id='target' class='a b c'></div>")
	s.Expect(s.Window.Eval(`
		list = document.getElementById("target").classList
		typeof list[Symbol.iterator]`)).To(Equal("function"))
}

func (s *ClassListTestSuite) TestIterableIteratesClassNames() {
	s.Window.MustLoadHTML("<div id='target' class='a b c'></div>")
	s.Expect(s.Window.Eval(`
		const list = document.getElementById('target').classList;
		Array.from(list).join(",")
	`)).To(Equal("a,b,c"))
}

func (s *ClassListTestSuite) TestToggleExistingClassName() {
	s.Window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.Window.Eval(`
		document.getElementById("target").classList.toggle("b")
	`)).To(BeFalse())
	s.Expect(
		s.Window.HTMLDocument().GetElementById("target"),
	).To(HaveAttribute("class", "a c"))
}

func (s *ClassListTestSuite) TestToggleNonExistingClassName() {
	s.Window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.Window.Eval(`
		document.getElementById("target").classList.toggle("x")
	`)).To(BeTrue())
	s.Expect(
		s.Window.HTMLDocument().GetElementById("target"),
	).To(HaveAttribute("class", "a b c x"))
}

func (s *ClassListTestSuite) TestToggleForceExistingItem() {
	s.Window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.Window.Eval(`
		document.getElementById("target").classList.toggle("b", true)
	`)).To(BeTrue(), "Force toggle return value")
	div := s.Window.Document().GetElementById("target")
	s.Expect(div).To(HaveAttribute("class", "a b c"))
}

func (s *ClassListTestSuite) TestToggleForceNonExistingItem() {
	s.Window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.Window.Eval(`
		document.getElementById("target").classList.toggle("x", true)
	`)).To(BeTrue())
	div := s.Window.Document().GetElementById("target")
	s.Expect(div).To(HaveAttribute("class", "a b c x"))
}

func (s *ClassListTestSuite) TestToggleNoForceExistingItem() {
	s.Window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.Window.Eval(`
		document.getElementById("target").classList.toggle("b", false)
	`)).To(BeFalse())
	div := s.Window.Document().GetElementById("target")
	s.Expect(div).To(HaveAttribute("class", "a c"))
}

func (s *ClassListTestSuite) TestToggleNoForceNonExistingItem() {
	s.Window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.Window.Eval(`
		document.getElementById("target").classList.toggle("x", false)
	`)).To(BeFalse())
	div := s.Window.Document().GetElementById("target")
	s.Expect(div).To(HaveAttribute("class", "a b c"))
}
