package dom_test

import (
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/stretchr/testify/suite"
)

type GomegaSuite struct {
	suite.Suite
	gomega gomega.Gomega
}

func (s *GomegaSuite) Expect(actual interface{}, extra ...interface{}) types.Assertion {
	if s.gomega == nil {
		s.gomega = gomega.NewWithT(s.T())
	}
	return s.gomega.Expect(actual, extra...)
}

// Designed as a "mixin" for a testify Suite, that can create a document on
// demand.
type DocumentSuite struct {
	doc dom.Document
}

func (s *DocumentSuite) Document() dom.Document {
	if s.doc == nil {
		s.doc = CreateHTMLDocument()
	}
	return s.doc
}

func ParseHtmlString(s string) dom.Document {
	win, err := html.NewWindowReader(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	return win.Document()
}

func CreateHTMLDocument() dom.Document {
	return html.NewHTMLDocument(nil)
}
