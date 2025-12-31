// This file is generated. Do not edit.

package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "AbortController", "", InitializeAbortController, AbortControllerConstructor)
	js.RegisterClass(e, "DOMTokenList", "", InitializeDOMTokenList, DOMTokenListConstructor)
	js.RegisterClass(e, "Event", "", InitializeEvent, EventConstructor)
	js.RegisterClass(e, "EventTarget", "", InitializeEventTarget, EventTargetConstructor)
	js.RegisterClass(e, "HTMLCollection", "", InitializeHTMLCollection, HTMLCollectionConstructor)
	js.RegisterClass(e, "MutationObserver", "", InitializeMutationObserver, MutationObserverConstructor)
	js.RegisterClass(e, "MutationRecord", "", InitializeMutationRecord, MutationRecordConstructor)
	js.RegisterClass(e, "NamedNodeMap", "", InitializeNamedNodeMap, NamedNodeMapConstructor)
	js.RegisterClass(e, "NodeList", "", InitializeNodeList, NodeListConstructor)
	js.RegisterClass(e, "AbortSignal", "EventTarget", InitializeAbortSignal, AbortSignalConstructor)
	js.RegisterClass(e, "Node", "EventTarget", InitializeNode, NodeConstructor)
	js.RegisterClass(e, "Attr", "Node", InitializeAttr, AttrConstructor)
	js.RegisterClass(e, "CharacterData", "Node", InitializeCharacterData, CharacterDataConstructor)
	js.RegisterClass(e, "Document", "Node", InitializeDocument, DocumentConstructor)
	js.RegisterClass(e, "DocumentFragment", "Node", InitializeDocumentFragment, DocumentFragmentConstructor)
	js.RegisterClass(e, "Element", "Node", InitializeElement, ElementConstructor)
	js.RegisterClass(e, "Text", "CharacterData", InitializeText, TextConstructor)
}
