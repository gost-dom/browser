// This file is generated. Do not edit.

package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "AbortController", "", NewAbortController)
	js.RegisterClass(reg, "AbortSignal", "EventTarget", NewAbortSignal)
	js.RegisterClass(reg, "Attr", "Node", NewAttr)
	js.RegisterClass(reg, "CharacterData", "Node", NewCharacterData)
	js.RegisterClass(reg, "ChildNode", "", NewChildNode)
	js.RegisterClass(reg, "DOMTokenList", "", NewDOMTokenList)
	js.RegisterClass(reg, "Document", "Node", NewDocument)
	js.RegisterClass(reg, "DocumentFragment", "Node", NewDocumentFragment)
	js.RegisterClass(reg, "Element", "Node", NewElement)
	js.RegisterClass(reg, "Event", "", NewEvent)
	js.RegisterClass(reg, "EventTarget", "", NewEventTarget)
	js.RegisterClass(reg, "HTMLCollection", "", NewHTMLCollection)
	js.RegisterClass(reg, "MutationObserver", "", NewMutationObserver)
	js.RegisterClass(reg, "MutationRecord", "", NewMutationRecord)
	js.RegisterClass(reg, "NamedNodeMap", "", NewNamedNodeMap)
	js.RegisterClass(reg, "Node", "EventTarget", NewNode)
	js.RegisterClass(reg, "NodeList", "", NewNodeList)
	js.RegisterClass(reg, "NonDocumentTypeChildNode", "", NewNonDocumentTypeChildNode)
	js.RegisterClass(reg, "ParentNode", "", NewParentNode)
	js.RegisterClass(reg, "Text", "CharacterData", NewText)
}
