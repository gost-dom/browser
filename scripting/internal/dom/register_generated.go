// This file is generated. Do not edit.

package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	InitializeAbortController(js.CreateClass(e, "AbortController", "", AbortControllerConstructor))
	InitializeDOMTokenList(js.CreateClass(e, "DOMTokenList", "", DOMTokenListConstructor))
	InitializeEvent(js.CreateClass(e, "Event", "", EventConstructor))
	InitializeEventTarget(js.CreateClass(e, "EventTarget", "", EventTargetConstructor))
	InitializeHTMLCollection(js.CreateClass(e, "HTMLCollection", "", HTMLCollectionConstructor))
	InitializeMutationObserver(js.CreateClass(e, "MutationObserver", "", MutationObserverConstructor))
	InitializeMutationRecord(js.CreateClass(e, "MutationRecord", "", MutationRecordConstructor))
	InitializeNamedNodeMap(js.CreateClass(e, "NamedNodeMap", "", NamedNodeMapConstructor))
	InitializeNodeList(js.CreateClass(e, "NodeList", "", NodeListConstructor))
	InitializeAbortSignal(js.CreateClass(e, "AbortSignal", "EventTarget", AbortSignalConstructor))
	InitializeNode(js.CreateClass(e, "Node", "EventTarget", NodeConstructor))
	InitializeAttr(js.CreateClass(e, "Attr", "Node", AttrConstructor))
	InitializeCharacterData(js.CreateClass(e, "CharacterData", "Node", CharacterDataConstructor))
	InitializeDocument(js.CreateClass(e, "Document", "Node", DocumentConstructor))
	InitializeDocumentFragment(js.CreateClass(e, "DocumentFragment", "Node", DocumentFragmentConstructor))
	InitializeElement(js.CreateClass(e, "Element", "Node", ElementConstructor))
	InitializeText(js.CreateClass(e, "Text", "CharacterData", TextConstructor))
}
