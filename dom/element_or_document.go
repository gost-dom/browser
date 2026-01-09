package dom

import (
	"strings"
)

// ElementOrDocument doesn't correspond to an interface in the web IDL specs. It
// contains common operations between dom Elements and documents.
type ElementOrDocument interface {
	GetElementsByTagName(name string) HTMLCollection
	GetElementsByTagNameNS(ns, name string) HTMLCollection
}

type tagNamePredicate struct {
	ns   string
	name string
}

func (p tagNamePredicate) Match(e Element) bool {
	return strings.EqualFold(p.name, e.LocalName()) && strings.EqualFold(p.ns, e.NamespaceURI())
}

type elementOrDocument struct {
	node *node
}

func (e elementOrDocument) GetElementsByTagName(name string) HTMLCollection {
	return newLiveHtmlCollection(e.node.self(), tagNamePredicate{"", name})
}

func (e elementOrDocument) GetElementsByTagNameNS(ns, name string) HTMLCollection {
	return newLiveHtmlCollection(e.node.self(), tagNamePredicate{ns, name})
}
