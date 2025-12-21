package htmltest

import (
	"testing"

	"github.com/gost-dom/browser/html"
)

type HTMLHelper struct {
	T testing.TB
}

func (h HTMLHelper) NewDocument() HTMLDocumentHelper {
	return NewHTMLDocumentHelper(h.T, html.NewHTMLDocument(nil))
}

// Create a single HTMLElement for testing element behaviour.
func (h HTMLHelper) CreateHTMLElement(tagname string) HTMLElementHelper {
	return NewHTMLElementHelper(h.T, h.NewDocument().CreateHTMLElement(tagname))
}
