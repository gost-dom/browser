// This file is generated. Do not edit.

package html

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	InitializeHistory(js.CreateClass(e, "History", "", HistoryConstructor))
	InitializeLocation(js.CreateClass(e, "Location", "", LocationConstructor))
	InitializeHTMLElement(js.CreateClass(e, "HTMLElement", "Element", HTMLElementConstructor))
	InitializeHTMLAnchorElement(js.CreateClass(e, "HTMLAnchorElement", "HTMLElement", HTMLAnchorElementConstructor))
	InitializeHTMLFormElement(js.CreateClass(e, "HTMLFormElement", "HTMLElement", HTMLFormElementConstructor))
	InitializeHTMLInputElement(js.CreateClass(e, "HTMLInputElement", "HTMLElement", HTMLInputElementConstructor))
	InitializeHTMLTemplateElement(js.CreateClass(e, "HTMLTemplateElement", "HTMLElement", HTMLTemplateElementConstructor))
}
