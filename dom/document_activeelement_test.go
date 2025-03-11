package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/stretchr/testify/suite"
)

type DocumentActiveElementTestSuite struct {
	suite.Suite
	doc dom.Document
}

func (s *DocumentActiveElementTestSuite) SetupTest() {
	s.doc = ParseHtmlString(activeDocumentTestHtml)
}

func TestDocumentActiveElement(t *testing.T) {
	suite.Run(t, new(DocumentActiveElementTestSuite))
}

func (s *DocumentActiveElementTestSuite) TestDefaultActiveElement() {
	s.Assert().Equal(s.doc.Body(), s.doc.ActiveElement(), "Active element is body by default")
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
