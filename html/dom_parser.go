package html

import (
	"io"

	"github.com/gost-dom/browser/dom"
)

// Deprecated: Will be removed
type domParser struct{}

// Deprecated: Will be removed
func (p domParser) ParseFragment(doc dom.Document, r io.Reader) (dom.DocumentFragment, error) {
	return ParseFragment(doc, r)
}

// Deprecated: Use [dom.ParseFragment]
func ParseFragment(owner dom.Document, r io.Reader) (dom.DocumentFragment, error) {
	return dom.ParseFragment(owner, r)
}

// Deprecated: Will be removed
func NewDOMParser() domParser { return domParser{} }

// Deprecated: Will be removed
type ElementSteps interface {
	AppendChild(parent dom.Node, child dom.Node) dom.Node
}
