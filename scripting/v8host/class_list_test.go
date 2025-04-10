package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/gost-dom/browser/scripting/v8host"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClassListTestTestSuite struct {
	gosttest.GomegaSuite
	scriptHost html.ScriptHost
	window     htmltest.WindowHelper
}

func (s *ClassListTestTestSuite) SetupTest() {
	s.window = htmltest.NewWindowHelper(s.T(), html.NewWindow(html.WindowOptions{
		ScriptHost: s.scriptHost,
	}))
}

func TestClassListTest(t *testing.T) {
	t.Parallel()
	suite.Run(t, &ClassListTestTestSuite{scriptHost: v8host.New()})
}

func (s *ClassListTestTestSuite) TestAdd() {
	s.window.MustLoadHTML("<div id='target' class='c1'></div>")
	assert.NoError(s.T(), s.window.Run(`
			const list = document.getElementById('target').classList;
			list.add('c2')
		`))
	element := s.window.HTMLDocument().GetHTMLElementById("target")
	s.Expect(element).To(HaveAttribute("class", "c1 c2"))
}

func (s *ClassListTestTestSuite) TestClassListIsIterable() {
	s.window.MustLoadHTML("<div id='target' class='a b c'></div>")
	s.Expect(s.window.Eval(`
		list = document.getElementById("target").classList
		typeof list[Symbol.iterator]`)).To(Equal("function"))
}

func (s *ClassListTestTestSuite) TestIterableIteratesClassNames() {
	s.window.MustLoadHTML("<div id='target' class='a b c'></div>")
	s.Expect(s.window.Eval(`
		const list = document.getElementById('target').classList;
		Array.from(list).join(",")
	`)).To(Equal("a,b,c"))
}

func (s *ClassListTestTestSuite) TestToggleExistingClassName() {
	s.window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.window.Eval(`
		document.getElementById("target").classList.toggle("b")
	`)).To(BeFalse())
	s.Expect(
		s.window.HTMLDocument().GetElementById("target"),
	).To(HaveAttribute("class", "a c"))
}

func (s *ClassListTestTestSuite) TestToggleNonExistingClassName() {
	s.window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.window.Eval(`
		document.getElementById("target").classList.toggle("x")
	`)).To(BeTrue())
	s.Expect(
		s.window.HTMLDocument().GetElementById("target"),
	).To(HaveAttribute("class", "a b c x"))
}

func (s *ClassListTestTestSuite) TestToggleForceExistingItem() {
	s.window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.window.Eval(`
		document.getElementById("target").classList.toggle("b", true)
	`)).To(BeTrue(), "Force toggle return value")
	div := s.window.Document().GetElementById("target")
	s.Expect(div).To(HaveAttribute("class", "a b c"))
}

func (s *ClassListTestTestSuite) TestToggleForceNonExistingItem() {
	s.window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.window.Eval(`
		document.getElementById("target").classList.toggle("x", true)
	`)).To(BeTrue())
	div := s.window.Document().GetElementById("target")
	s.Expect(div).To(HaveAttribute("class", "a b c x"))
}

func (s *ClassListTestTestSuite) TestToggleNoForceExistingItem() {
	s.window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.window.Eval(`
		document.getElementById("target").classList.toggle("b", false)
	`)).To(BeFalse())
	div := s.window.Document().GetElementById("target")
	s.Expect(div).To(HaveAttribute("class", "a c"))
}

func (s *ClassListTestTestSuite) TestToggleNoForceNonExistingItem() {
	s.window.MustLoadHTML(`<div id="target" class="a b c"></div>`)
	s.Expect(s.window.Eval(`
		document.getElementById("target").classList.toggle("x", false)
	`)).To(BeFalse())
	div := s.window.Document().GetElementById("target")
	s.Expect(div).To(HaveAttribute("class", "a b c"))
}
