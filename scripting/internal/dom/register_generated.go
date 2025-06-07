// This file is generated. Do not edit.

package dom

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "Attr", "Node", NewAttrV8Wrapper)
	js.RegisterClass(reg, "DOMTokenList", "", NewDOMTokenListV8Wrapper)
	js.RegisterClass(reg, "Document", "Node", NewDocumentV8Wrapper)
	js.RegisterClass(reg, "DocumentFragment", "Node", NewDocumentFragmentV8Wrapper)
	js.RegisterClass(reg, "Element", "Node", NewElementV8Wrapper)
	js.RegisterClass(reg, "Event", "", NewEventV8Wrapper)
	js.RegisterClass(reg, "EventTarget", "", NewEventTargetV8Wrapper)
	js.RegisterClass(reg, "MutationObserver", "", NewMutationObserverV8Wrapper)
	js.RegisterClass(reg, "MutationRecord", "", NewMutationRecordV8Wrapper)
	js.RegisterClass(reg, "NamedNodeMap", "", NewNamedNodeMapV8Wrapper)
	js.RegisterClass(reg, "Node", "EventTarget", NewNodeV8Wrapper)
	js.RegisterClass(reg, "NodeList", "", NewNodeListV8Wrapper)
	js.RegisterClass(reg, "NonDocumentTypeChildNode", "", NewNonDocumentTypeChildNodeV8Wrapper)
	js.RegisterClass(reg, "ParentNode", "", NewParentNodeV8Wrapper)
	js.RegisterClass(reg, "PointerEvent", "MouseEvent", NewPointerEventV8Wrapper)
	js.RegisterClass(reg, "MouseEvent", "UIEvent", NewMouseEventV8Wrapper)
	js.RegisterClass(reg, "UIEvent", "Event", NewUIEventV8Wrapper)
}
