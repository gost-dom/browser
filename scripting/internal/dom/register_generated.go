// This file is generated. Do not edit.

package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	InitializeAbortController(js.CreateClass(e, "AbortController", "", AbortControllerConstructor))
	InitializeDOMTokenList(js.CreateClass(e, "DOMTokenList", "", nil))
	InitializeEvent(js.CreateClass(e, "Event", "", EventConstructor))
	InitializeEventTarget(js.CreateClass(e, "EventTarget", "", EventTargetConstructor))
	InitializeHTMLCollection(js.CreateClass(e, "HTMLCollection", "", nil))
	InitializeMutationObserver(js.CreateClass(e, "MutationObserver", "", MutationObserverConstructor))
	InitializeMutationRecord(js.CreateClass(e, "MutationRecord", "", nil))
	InitializeNamedNodeMap(js.CreateClass(e, "NamedNodeMap", "", nil))
	InitializeNodeList(js.CreateClass(e, "NodeList", "", nil))
	InitializeAbortSignal(js.CreateClass(e, "AbortSignal", "EventTarget", nil))
	InitializeNode(js.CreateClass(e, "Node", "EventTarget", nil))
	InitializeAttr(js.CreateClass(e, "Attr", "Node", nil))
	InitializeCharacterData(js.CreateClass(e, "CharacterData", "Node", nil))
	InitializeDocument(js.CreateClass(e, "Document", "Node", DocumentConstructor))
	InitializeDocumentFragment(js.CreateClass(e, "DocumentFragment", "Node", DocumentFragmentConstructor))
	InitializeElement(js.CreateClass(e, "Element", "Node", nil))
	InitializeText(js.CreateClass(e, "Text", "CharacterData", nil))
}
