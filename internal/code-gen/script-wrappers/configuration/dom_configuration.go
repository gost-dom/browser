package configuration

func ConfigureDOMSpecs(specs *WebIdlConfigurations) {
	domSpecs := specs.Module("dom")
	domSpecs.SetMultipleFiles(true)
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
