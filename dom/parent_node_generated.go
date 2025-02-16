// This file is generated. Do not edit.

package dom

type ParentNode interface {
	Children() HTMLCollection
	FirstElementChild() Element
	LastElementChild() Element
	ChildElementCount() int
	Prepend(nodes ...Node)
	Append(nodes ...Node)
	ReplaceChildren(nodes ...Node)
	QuerySelector(string) Element
	QuerySelectorAll(string) NodeList
}
