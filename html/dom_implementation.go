package html

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/entity"
	// dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
)

type DOMImplementation struct {
	entity.Entity
	OwnerDocument dom.Document
}

// var _ dominterfaces.DOMImplementation = &DOMImplementation{}

func (i *DOMImplementation) CreateDocument(
	namespace, qualifiedName string,
	docType dom.DocumentType,
) dom.Document {
	return dom.NewDocument(nil)
}

func (i *DOMImplementation) CreateHTMLDocument(title string) HTMLDocument {
	return NewValidHTMLDocument(nil)
}

func (i *DOMImplementation) HasFeature() bool { return true }
