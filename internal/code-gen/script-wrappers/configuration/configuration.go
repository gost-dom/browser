// Package configuration is part of an internal code generation tool for
// Gost-DOM. It is not to be used in other context, and is not used at runtime.
package configuration

func CreateSpecs() WebIdlConfigurations {
	specs := NewWrapperGeneratorsSpec()
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

	return specs
}
