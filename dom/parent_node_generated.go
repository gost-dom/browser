// This file is generated. Do not edit.

package dom

type ParentNode interface {
	Children() HTMLCollection
	FirstElementChild() Element
	LastElementChild() Element
	ChildElementCount() int
	/*
	   Note that the IDL operation accepts either string or node values. This interface
	   requires an explicit a [Node]. Use [Document.CreateText] to convert a string to
	   a Node.

	   See also: https://developer.mozilla.org/en-US/docs/Web/API/Element
	*/
	Prepend(nodes ...Node) error
	/*
	   Note that the IDL operation accepts either string or node values. This interface
	   requires an explicit a [Node]. Use [Document.CreateText] to convert a string to
	   a Node.

	   See also: https://developer.mozilla.org/en-US/docs/Web/API/Element
	*/
	Append(nodes ...Node) error
	/*
	   Note that the IDL operation accepts either string or node values. This interface
	   requires an explicit a [Node]. Use [Document.CreateText] to convert a string to
	   a Node.

	   See also: https://developer.mozilla.org/en-US/docs/Web/API/Element
	*/
	ReplaceChildren(nodes ...Node) error
	QuerySelector(string) (Element, error)
	QuerySelectorAll(string) (NodeList, error)
}
