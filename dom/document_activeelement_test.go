package dom_test

import (
	"fmt"
	"testing"

	. "github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/testing/eventtest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/suite"
)

type DocumentActiveElementTestSuite struct {
	suite.Suite
	doc htmltest.HTMLDocumentHelper
}

func (s *DocumentActiveElementTestSuite) SetupTest() {
	s.doc = htmltest.ParseHTMLDocumentHelper(s.T(), activeDocumentTestHtml)
}

func TestDocumentActiveElement(t *testing.T) {
	suite.Run(t, new(DocumentActiveElementTestSuite))
}

func (s *DocumentActiveElementTestSuite) TestDefaultActiveElement() {
	s.Assert().Equal(s.doc.Body(), s.doc.ActiveElement(), "Active element is body by default")
}

func (s *DocumentActiveElementTestSuite) TestFocusEvents() {
	var logs []string
	var logHandler = func(name string) EventHandler {
		return eventtest.NewTestHandler(func(e *Event) {
			logs = append(logs, fmt.Sprintf("%s - %s", e.Type, name))
		})
	}
	nodes, err := s.doc.QuerySelectorAll("div, input")
	s.Assert().NoError(err)
	for _, e := range nodes.All() {
		id, _ := (e.(Element).GetAttribute("id"))
		e.AddEventListener("focus", logHandler(id))
		e.AddEventListener("focusin", logHandler(id))
	}
	target := s.doc.GetHTMLElementById("input-1-1-1")
	target.Focus()
	s.Assert().Equal([]string{
		"focus - input-1-1-1",
		"focusin - input-1-1-1",
		"focusin - child-1-1",
		"focusin - root-1",
	}, logs)
	s.Assert().Equal(target, s.doc.ActiveElement())
}

const activeDocumentTestHtml = `<body>
	<div id="root-1">
		<div id="child-1-1">
			<input id="input-1-1-1" type="text" />
			<input id="input-1-1-2" type="text" />
		</div>
		<div id="child-1-2">
			<button id="button-1-2-1">Button 1</button>
		</div>
	</div>
<body>`
