package html

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/entity"
	// dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
)

// DOMImplementation is highly experimental implementation of the dom.implementation
//
// This primarily exists to fill enough gaps to run Web Platform Tests
// exercising other features.
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

func (i *DOMImplementation) HasFeature() bool { return true }

func (i *DOMImplementation) CreateDocumentType(
	qualifiedName, publicID, systemID string,
) dom.DocumentType {
	return dom.NewDocumentType(qualifiedName, publicID, systemID, nil)
}
