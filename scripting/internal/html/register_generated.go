// This file is generated. Do not edit.

package html

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	document, ok := e.Class("Document")
	if !ok {
		panic("gost-dom/html: Document: class not registered")
	}
	InitializeDocument(document)
	element, ok := e.Class("Element")
	if !ok {
		panic("gost-dom/html: Element: class not registered")
	}
	InitializeElement(element)
	InitializeHistory(js.CreateClass(e, "History", "", nil))
	InitializeLocation(js.CreateClass(e, "Location", "", nil))
	InitializeHTMLElement(js.CreateClass(e, "HTMLElement", "Element", nil))
	InitializeHTMLAnchorElement(js.CreateClass(e, "HTMLAnchorElement", "HTMLElement", nil))
	InitializeHTMLFormElement(js.CreateClass(e, "HTMLFormElement", "HTMLElement", nil))
	InitializeHTMLInputElement(js.CreateClass(e, "HTMLInputElement", "HTMLElement", nil))
	InitializeHTMLTemplateElement(js.CreateClass(e, "HTMLTemplateElement", "HTMLElement", nil))
}
