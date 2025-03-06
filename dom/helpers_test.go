package dom_test

import (
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
)

func NewTestHandler(
	f func(*event.Event),
) event.EventHandler {
	return event.NewEventHandlerFunc(event.NoError(f))
}

// Designed as a "mixin" for a testify Suite, that can create a document on
// demand.
type DocumentSuite struct {
	doc dom.Document
}

func (s *DocumentSuite) CreateDocument() dom.Document {
	return CreateHTMLDocument()
}

func (s *DocumentSuite) Document() dom.Document {
	if s.doc == nil {
		s.doc = s.CreateDocument()
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
