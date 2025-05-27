package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/test/scripttests"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/scripting/v8host"
	"github.com/stretchr/testify/suite"
)

type DocumentTestSuite struct {
	// Note, document tests also exists in the internal/test/scripting folder.
	scripttests.ScriptHostSuite
}

func TestDocument(t *testing.T) {
	suite.Run(t, &DocumentTestSuite{*scripttests.NewScriptHostSuite(v8host.New())})
}

func (s *DocumentTestSuite) SetupTest() {
	s.ScriptHostSuite.SetupTest()
	s.MustLoadHTML(` <html> <body> </body> </html>`)
}

func (s *DocumentTestSuite) TestDocumentElement() {
	s.Assert().Equal(
		"HTMLHtmlElement",
		s.MustEval("Object.getPrototypeOf(document.documentElement).constructor.name"))
}

func (s *DocumentTestSuite) TestLocationAttribute() {
	s.Assert().
		Equal(true,
			s.MustEval("document.location instanceof Location"),
			"document.location instanceof Location")
	s.Assert().
		Equal(true,
			s.MustEval("document.location === location"),
			"document.location is same as global location object")
}

func (s *DocumentTestSuite) TestDocumentStructur() {
	s.Assert().Equal("BODY", s.MustEval("document.body.tagName"), "<body> tagName")
	s.Assert().Equal("HEAD", s.MustEval("document.head.tagName"), "<head> tagName")
}

func (s *DocumentTestSuite) TestQuerySelector() {
	s.MustLoadHTML(
		`<body><div>0</div><div data-key="1">1</div><div data-key="2">2</div><body>`,
	)
	s.Assert().Equal(
		`<div data-key="1">1</div>`,
		s.MustEval("document.querySelector('[data-key]').outerHTML"))
	s.Expect(
		s.MustEval(`document.querySelector('[data-key="2"]').outerHTML`),
	).To(Equal(`<div data-key="2">2</div>`))
	s.Expect(
		s.MustEval(`document.querySelector('script')`),
	).To(BeNil())
}

func (s *DocumentTestSuite) TestQuerySelectorAll() {
	s.MustLoadHTML(
		`<body><div>0</div><div data-key="1">1</div><div data-key="2">2</div><body>`,
	)
	s.Expect(
		s.MustEval(
			"Array.from(document.querySelectorAll('[data-key]')).map(x => x.outerHTML).join(',')",
		),
	).To(Equal(`<div data-key="1">1</div>,<div data-key="2">2</div>`))
}

func (s *DocumentTestSuite) TestCreateDocumentFragment() {
	s.Expect(s.MustEval(`
		const fragment = document.createDocumentFragment();
		Object.getPrototypeOf(fragment) === DocumentFragment.prototype
	`)).To(BeTrue())
}
