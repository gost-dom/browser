package configuration

func ConfigureDOMSpecs(specs *WebIdlConfigurations) {
	domSpecs := specs.Module("dom")
	configureDOMNode(domSpecs)
	configureDOMEvent(domSpecs)
}

func configureDOMNode(specs *WebAPIConfig) {
	configureMutationObserver(specs)

	childNode := specs.Type("ChildNode")
	childNode.MarkMembersAsNotImplemented(
		"before", "after", "replaceWith",
	)
	cd := specs.Type("CharacterData")
	cd.MarkMembersAsNotImplemented(
		"substringData", "appendData", "insertData", "deleteData", "replaceData",
	)
	text := specs.Type("Text")
	text.MarkMembersAsNotImplemented(
		"splitText", "wholeText",
	)

	specs.Type("AbortController")
	abortSignal := specs.Type("AbortSignal")
	abortSignal.MarkMembersAsNotImplemented("reason")

	docFrag := specs.Type("DocumentFragment")
	docFrag.SkipConstructor = true

	attr := specs.Type("Attr")
	attr.MarkMembersAsIgnored("namespaceURI", "prefix", "specified")

	eventTarget := specs.Type("EventTarget")
	addEventListenerOptions := eventTarget.Method("addEventListener").Argument("options")
	addEventListenerOptions.SetDecoder("w.decodeEventListenerOptions")
	addEventListenerOptions.HasDefault = true
	addEventListenerOptions.DefaultValue = "defaultEventListenerOptions"
	removeEventListenerOptions := eventTarget.Method("removeEventListener").Argument("options")
	removeEventListenerOptions.SetDecoder("w.decodeEventListenerOptions")
	removeEventListenerOptions.HasDefault = true
	removeEventListenerOptions.DefaultValue = "defaultEventListenerOptions"

	namedNodeMap := specs.Type("NamedNodeMap")
	namedNodeMap.MarkMembersAsNotImplemented(
		"getNamedItem",
		"setNamedItem",
		"getNamedItemNS",
		"setNamedItemNS",
		"removeNamedItem",
		"removeNamedItemNS",
	)
	namedNodeMap.RunCustomCode = true

	specs.Type("NonDocumentTypeChildNode")
	document := specs.Type("Document")
	document.RunCustomCode = true // Set instance properties
	document.MarkMembersAsNotImplemented(
		"createNodeIterator",
		"createTreeWalker",
		"getElementsByTagName",
		"getElementsByTagNameNS",
		"getElementsByClassName",
		"createProcessingInstruction",
		"adoptNode",
		"createRange",
		"createEvent",
		"implementation",
		"documentURI",
		"doctype",
		"contentType",
		"inputEncoding",
		"charset", "characterSet",
		"compatMode", "URL",
		"createAttributeNS",
		"createElementNS",

		// Custom Implementation
		"createCDATASection",
	)

	// createElement has `is` option, relating to web components
	document.Method("createElement").SetCustomImplementation()

	nodeList := specs.Type("NodeList")
	nodeList.RunCustomCode = true

	htmlCollection := specs.Type("HTMLCollection")
	htmlCollection.RunCustomCode = true

	parentNode := specs.Type("ParentNode")
	parentNode.Method("append").Argument("nodes").Decoder = "w.decodeNodeOrText"
	parentNode.Method("prepend").Argument("nodes").Decoder = "w.decodeNodeOrText"
	parentNode.Method("replaceChildren").Argument("nodes").Decoder = "w.decodeNodeOrText"

	domElement := specs.Type("Element")
	// domElement.SkipWrapper = true
	domElement.RunCustomCode = true
	domElement.Method("classList").SetCustomImplementation()

	domElement.MarkMembersAsNotImplemented(
		"hasAttributes",
		"hasAttributeNS",
		"getAttributeNames",
		"getAttributeNS",
		"setAttributeNS",
		"removeAttributeNode",
		"removeAttributeNS",
		"toggleAttribute",
		"toggleAttributeForce",
		"setAttributeNode",
		"setAttributeNodeNS",
		"getAttributeNode",
		"getAttributeNodeNS",
		"getElementsByTagName",
		"getElementsByTagNameNS",
		"getElementsByClassName",
		"insertAdjacentText",
		"namespaceURI",
		"prefix",
		"localName",
		"shadowRoot",
		"slot",
		"className",
		"decodeShadowRootInit",
		"attachShadow",
	)

	domElement.MarkMembersAsIgnored(
		// HTMX fails if these exist but throw
		"webkitMatchesSelector",
	)

	domTokenList := specs.Type("DOMTokenList")
	domTokenList.RunCustomCode = true

	// Toggle has custom implementation, because the force option has behaviour that
	// doesn't make supermuch sense to have in the internal DOM implementation
	domTokenList.Method("toggle").SetCustomImplementation()
	domTokenList.Method("supports").SetNotImplemented()

	domNode := specs.Type("Node")
	domNode.Method("nodeType").SetCustomImplementation()
	domNode.Method("getRootNode").Argument("options").SetHasDefault()
	domNode.Method("textContent").SetCustomImplementation()

	domNode.Method("hasChildNodes").Ignore()
	domNode.Method("normalize").Ignore()
	domNode.Method("compareDocumentPosition").Ignore()
	domNode.Method("lookupPrefix").Ignore()
	domNode.Method("lookupNamespaceURI").Ignore()
	domNode.Method("isDefaultNamespace").Ignore()
	domNode.Method("replaceChild").Ignore()
	domNode.Method("baseURI").Ignore()
	domNode.Method("lastChild").Ignore()
	domNode.Method("nodeValue").Ignore()
}

func configureDOMEvent(domSpecs *WebAPIConfig) {
	event := domSpecs.Type("Event")
	// event.SkipWrapper = true
	event.Method("eventPhase").SetCustomImplementation()
	event.Method("initEvent").Ignore()
	event.Method("composed").Ignore()
	event.Method("composedPath").Ignore()
	event.Method("stopImmediatePropagation").Ignore()
	event.Method("isTrusted").Ignore()
	event.Method("cancelBubble").Ignore()
	event.Method("timeStamp").Ignore()
	event.Method("returnValue").Ignore()
	event.Method("srcElement").Ignore()
}

func configureMutationObserver(domSpecs *WebAPIConfig) {
	domSpecs.Type("MutationObserver")
	domSpecs.Type("MutationRecord")
}
