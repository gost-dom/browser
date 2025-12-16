package dom

import (
	"strings"
)

// ElementOrDocument doesn't correspond to an interface in the web IDL specs. It
// contains common operations between dom Elements and documents.
type ElementOrDocument interface {
	GetElementsByTagName(string) HTMLCollection
}

type tagNamePredicate struct {
	name string
}

func (p tagNamePredicate) Match(e Element) bool {
	return strings.EqualFold(p.name, e.LocalName())
}

type elementOrDocument struct {
	node *node
}

func (e elementOrDocument) GetElementsByTagName(name string) HTMLCollection {
	return newLiveHtmlCollection(e.node.self, tagNamePredicate{name})
}
