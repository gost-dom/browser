// This file is generated. Do not edit.

package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "AbortController", "", AbortController[T]{}.Initialize, AbortControllerConstructor)
	js.RegisterClass(e, "DOMTokenList", "", DOMTokenList[T]{}.Initialize, DOMTokenListConstructor)
	js.RegisterClass(e, "Event", "", Event[T]{}.Initialize, EventConstructor)
	js.RegisterClass(e, "EventTarget", "", EventTarget[T]{}.Initialize, EventTargetConstructor)
	js.RegisterClass(e, "HTMLCollection", "", HTMLCollection[T]{}.Initialize, HTMLCollectionConstructor)
	js.RegisterClass(e, "MutationObserver", "", MutationObserver[T]{}.Initialize, MutationObserverConstructor)
	js.RegisterClass(e, "MutationRecord", "", MutationRecord[T]{}.Initialize, MutationRecordConstructor)
	js.RegisterClass(e, "NamedNodeMap", "", NamedNodeMap[T]{}.Initialize, NamedNodeMapConstructor)
	js.RegisterClass(e, "NodeList", "", NodeList[T]{}.Initialize, NodeListConstructor)
	js.RegisterClass(e, "AbortSignal", "EventTarget", AbortSignal[T]{}.Initialize, AbortSignalConstructor)
	js.RegisterClass(e, "Node", "EventTarget", Node[T]{}.Initialize, NodeConstructor)
	js.RegisterClass(e, "Attr", "Node", Attr[T]{}.Initialize, AttrConstructor)
	js.RegisterClass(e, "CharacterData", "Node", CharacterData[T]{}.Initialize, CharacterDataConstructor)
	js.RegisterClass(e, "Document", "Node", Document[T]{}.Initialize, DocumentConstructor)
	js.RegisterClass(e, "DocumentFragment", "Node", DocumentFragment[T]{}.Initialize, DocumentFragmentConstructor)
	js.RegisterClass(e, "Element", "Node", Element[T]{}.Initialize, ElementConstructor)
	js.RegisterClass(e, "Text", "CharacterData", Text[T]{}.Initialize, TextConstructor)
}
