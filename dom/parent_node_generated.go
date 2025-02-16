// This file is generated. Do not edit.

package dom

type ParentNode interface {
	Children() HTMLCollection
	FirstElementChild() Element
	LastElementChild() Element
	ChildElementCount() int
	Prepend(nodes ...Node) error
	Append(nodes ...Node) error
	ReplaceChildren(nodes ...Node) error
	QuerySelector(string) Element
	QuerySelectorAll(string) NodeList
}
