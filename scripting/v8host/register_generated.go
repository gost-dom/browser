// This file is generated. Do not edit.

package v8host

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "Attr", "Node", newAttrV8Wrapper)
	js.RegisterClass(reg, "DOMTokenList", "", newDOMTokenListV8Wrapper)
	js.RegisterClass(reg, "Document", "Node", newDocumentV8Wrapper)
	js.RegisterClass(reg, "Element", "Node", newElementV8Wrapper)
	js.RegisterClass(reg, "Event", "", newEventV8Wrapper)
	js.RegisterClass(reg, "EventTarget", "", newEventTargetV8Wrapper)
	js.RegisterClass(reg, "MutationObserver", "", newMutationObserverV8Wrapper)
	js.RegisterClass(reg, "MutationRecord", "", newMutationRecordV8Wrapper)
	js.RegisterClass(reg, "NamedNodeMap", "", newNamedNodeMapV8Wrapper)
	js.RegisterClass(reg, "Node", "EventTarget", newNodeV8Wrapper)
	js.RegisterClass(reg, "NodeList", "", newNodeListV8Wrapper)
	js.RegisterClass(reg, "NonDocumentTypeChildNode", "", newNonDocumentTypeChildNodeV8Wrapper)
	js.RegisterClass(reg, "ParentNode", "", newParentNodeV8Wrapper)
	js.RegisterClass(reg, "HTMLAnchorElement", "HTMLElement", newHTMLAnchorElementV8Wrapper)
	js.RegisterClass(reg, "HTMLElement", "Element", newHTMLElementV8Wrapper)
	js.RegisterClass(reg, "HTMLFormElement", "HTMLElement", newHTMLFormElementV8Wrapper)
	js.RegisterClass(reg, "HTMLHyperlinkElementUtils", "", newHTMLHyperlinkElementUtilsV8Wrapper)
	js.RegisterClass(reg, "HTMLInputElement", "HTMLElement", newHTMLInputElementV8Wrapper)
	js.RegisterClass(reg, "HTMLOrSVGElement", "", newHTMLOrSVGElementV8Wrapper)
	js.RegisterClass(reg, "HTMLTemplateElement", "HTMLElement", newHTMLTemplateElementV8Wrapper)
	js.RegisterClass(reg, "History", "", newHistoryV8Wrapper)
	js.RegisterClass(reg, "Location", "", newLocationV8Wrapper)
	js.RegisterClass(reg, "Window", "EventTarget", newWindowV8Wrapper)
	js.RegisterClass(reg, "PointerEvent", "MouseEvent", newPointerEventV8Wrapper)
	js.RegisterClass(reg, "MouseEvent", "UIEvent", newMouseEventV8Wrapper)
	js.RegisterClass(reg, "UIEvent", "Event", newUIEventV8Wrapper)
	js.RegisterClass(reg, "URL", "", newURLV8Wrapper)
	js.RegisterClass(reg, "URLSearchParams", "", newURLSearchParamsV8Wrapper)
	js.RegisterClass(reg, "FormData", "", newFormDataV8Wrapper)
	js.RegisterClass(reg, "XMLHttpRequest", "XMLHttpRequestEventTarget", newXMLHttpRequestV8Wrapper)
	js.RegisterClass(reg, "XMLHttpRequestEventTarget", "EventTarget", newXMLHttpRequestEventTargetV8Wrapper)
}
