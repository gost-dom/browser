// This file is generated. Do not edit.

package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "AbortController", "", NewAbortController)
	js.RegisterClass(e, "DOMTokenList", "", NewDOMTokenList)
	js.RegisterClass(e, "Event", "", NewEvent)
	js.RegisterClass(e, "EventTarget", "", NewEventTarget)
	js.RegisterClass(e, "HTMLCollection", "", NewHTMLCollection)
	js.RegisterClass(e, "MutationObserver", "", NewMutationObserver)
	js.RegisterClass(e, "MutationRecord", "", NewMutationRecord)
	js.RegisterClass(e, "NamedNodeMap", "", NewNamedNodeMap)
	js.RegisterClass(e, "NodeList", "", NewNodeList)
	js.RegisterClass(e, "AbortSignal", "EventTarget", NewAbortSignal)
	js.RegisterClass(e, "Node", "EventTarget", NewNode)
	js.RegisterClass(e, "Attr", "Node", NewAttr)
	js.RegisterClass(e, "CharacterData", "Node", NewCharacterData)
	js.RegisterClass(e, "Document", "Node", NewDocument)
	js.RegisterClass(e, "DocumentFragment", "Node", NewDocumentFragment)
	js.RegisterClass(e, "Element", "Node", NewElement)
	js.RegisterClass(e, "Text", "CharacterData", NewText)
}
