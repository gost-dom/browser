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
	Prepend(...Node) error
	/*
	   Note that the IDL operation accepts either string or node values. This interface
	   requires an explicit a [Node]. Use [Document.CreateText] to convert a string to
	   a Node.

	   See also: https://developer.mozilla.org/en-US/docs/Web/API/Element
	*/
	Append(...Node) error
	/*
	   Note that the IDL operation accepts either string or node values. This interface
	   requires an explicit a [Node]. Use [Document.CreateText] to convert a string to
	   a Node.

	   See also: https://developer.mozilla.org/en-US/docs/Web/API/Element
	*/
	ReplaceChildren(...Node) error
	// QuerySelector returns the first element that matches the CSS selector.
	// Returns nil if no match was found. Returns an error if selector cannot be
	// parsed.
	//
	// See also: https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
	QuerySelector(selector string) (Element, error)
	// QuerySelectorAll returns a static NodeList of all elements matching the
	// selector. Returns an error if selector cannot be parsed.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
	QuerySelectorAll(selector string) (NodeList, error)
}
