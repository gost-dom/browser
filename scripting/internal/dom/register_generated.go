// This file is generated. Do not edit.

package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "AbortController", "", NewAbortController, AbortControllerConstructor)
	js.RegisterClass(e, "DOMTokenList", "", NewDOMTokenList, DOMTokenListConstructor)
	js.RegisterClass(e, "Event", "", NewEvent, EventConstructor)
	js.RegisterClass(e, "EventTarget", "", NewEventTarget, EventTargetConstructor)
	js.RegisterClass(e, "HTMLCollection", "", NewHTMLCollection, HTMLCollectionConstructor)
	js.RegisterClass(e, "MutationObserver", "", NewMutationObserver, MutationObserverConstructor)
	js.RegisterClass(e, "MutationRecord", "", NewMutationRecord, MutationRecordConstructor)
	js.RegisterClass(e, "NamedNodeMap", "", NewNamedNodeMap, NamedNodeMapConstructor)
	js.RegisterClass(e, "NodeList", "", NewNodeList, NodeListConstructor)
	js.RegisterClass(e, "AbortSignal", "EventTarget", NewAbortSignal, AbortSignalConstructor)
	js.RegisterClass(e, "Node", "EventTarget", NewNode, NodeConstructor)
	js.RegisterClass(e, "Attr", "Node", NewAttr, AttrConstructor)
	js.RegisterClass(e, "CharacterData", "Node", NewCharacterData, CharacterDataConstructor)
	js.RegisterClass(e, "Document", "Node", NewDocument, DocumentConstructor)
	js.RegisterClass(e, "DocumentFragment", "Node", NewDocumentFragment, DocumentFragmentConstructor)
	js.RegisterClass(e, "Element", "Node", NewElement, ElementConstructor)
	js.RegisterClass(e, "Text", "CharacterData", NewText, TextConstructor)
}
