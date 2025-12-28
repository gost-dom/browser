package configuration

func ConfigureDOMSpecs(domSpecs *WebAPIConfig) {
	configureDOMNode(domSpecs)
	configureDOMEvent(domSpecs)
	domSpecs.AddSearchModule("html")
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
	attr.MarkMembersAsIgnored("prefix", "specified")

	eventTarget := specs.Type("EventTarget")
	addEventListenerOptions := eventTarget.Method("addEventListener").Argument("options")
	addEventListenerOptions.SetDecoder("decodeEventListenerOptions")
	addEventListenerOptions.HasDefault = true
	addEventListenerOptions.DefaultValue = "defaultEventListenerOptions"
	removeEventListenerOptions := eventTarget.Method("removeEventListener").Argument("options")
	removeEventListenerOptions.SetDecoder("decodeEventListenerOptions")
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
		"getElementsByClassName",
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

		// Custom Implementation
		"createCDATASection",
	)

	document.MarkMembersAsIgnored(
		"fgColor", "linkColor", "vlinkColor", "alinkColor", "bgColor", "anchors", "applets", "all",
		"readyState",

		// HTML spec
		"clear",
		"captureEvents",
		"releaseEvents",
		"getElementsByName",
		"open",
		"close",
		"write",
		"writeln",
		"hasFocus",
		"execCommand",
		"queryCommandEnabled", "queryCommandIndeterm",
		"queryCommandState",
		"queryCommandSupported",
		"queryCommandValue",
		"domain",
		"referrer",
		"cookie",
		"dir",

		// TODO: Use these
		"title",
		"body",
		"head",
		"images",
		"embeds",
		"plugins",
		"forms",
		"links",
		"scripts",
		"currentScript",
		"defaultView",
		"designMode",
		"hidden",
		"visibilityState",
		"lastModified",
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
	domElement.Method("classList").SetCustomImplementation()
	domElement.RunCustomCode = true

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
		"getElementsByClassName",
		"namespaceURI",
		"prefix",
		"shadowRoot",
		"slot",
		"className",
		"decodeShadowRootInit",
		"attachShadow",
	)

	domElement.MarkMembersAsIgnored(
		// HTMX fails if these exist but throw
		"webkitMatchesSelector",

		"setHTMLUnsafe",
		"getHTML",
	)

	domTokenList := specs.Type("DOMTokenList")
	domTokenList.RunCustomCode = true

	// Toggle has custom implementation, because the force option has behaviour that
	// doesn't make supermuch sense to have in the internal DOM implementation
	domTokenList.Method("toggle").SetCustomImplementation()
	domTokenList.Method("supports").SetNotImplemented()

	domNode := specs.Type("Node")
	// domNode.Method("nodeType").SetCustomImplementation()
	domNode.Method("getRootNode").Argument("options").SetHasDefault()
	domNode.Method("textContent").SetCustomImplementation()

	domNode.Method("hasChildNodes").Ignore()
	domNode.Method("normalize").Ignore()
	domNode.Method("compareDocumentPosition").Ignore()
	domNode.Method("lookupPrefix").Ignore()
	domNode.Method("lookupNamespaceURI").Ignore()
	domNode.Method("isDefaultNamespace").Ignore()
	domNode.Method("baseURI").Ignore()
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
