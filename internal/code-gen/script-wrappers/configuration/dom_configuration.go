package configuration

func ConfigureDOMSpecs(specs *WebIdlConfigurations) {
	domSpecs := specs.Module("dom")
	domSpecs.SetMultipleFiles(true)
	configureDOMNode(domSpecs)
	configureDOMEvent(domSpecs)
}

func configureDOMNode(domSpecs *WebIdlConfiguration) {
	domNode := domSpecs.Type("Node")
	domNode.Method("nodeType").SetCustomImplementation()
	domNode.Method("getRootNode").Argument("options").SetHasDefault()
	domNode.Method("cloneNode").Argument("subtree").SetHasDefault()
	domNode.Method("textContent").SetCustomImplementation()

	domNode.Method("hasChildNodes").Ignore()
	domNode.Method("normalize").Ignore()
	domNode.Method("isEqualNode").Ignore()
	domNode.Method("compareDocumentPosition").Ignore()
	domNode.Method("lookupPrefix").Ignore()
	domNode.Method("lookupNamespaceURI").Ignore()
	domNode.Method("isDefaultNamespace").Ignore()
	domNode.Method("replaceChild").Ignore()
	domNode.Method("baseURI").Ignore()
	domNode.Method("parentNode").Ignore()
	domNode.Method("lastChild").Ignore()
	domNode.Method("nodeValue").Ignore()
}

func configureDOMEvent(domSpecs *WebIdlConfiguration) {
	event := domSpecs.Type("Event")
	event.SkipWrapper = true
	event.Method("type").SetCustomImplementation()
	event.Method("bubbles").SetCustomImplementation()
	event.Method("cancelable").SetCustomImplementation()
	event.Method("eventPhase").SetCustomImplementation()
	event.Method("constructor").Argument("eventInitDict").SetHasDefault()
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
