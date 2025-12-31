// This file is generated. Do not edit.

package html

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "History", "", InitializeHistory, HistoryConstructor)
	js.RegisterClass(e, "Location", "", InitializeLocation, LocationConstructor)
	js.RegisterClass(e, "HTMLElement", "Element", InitializeHTMLElement, HTMLElementConstructor)
	js.RegisterClass(e, "HTMLAnchorElement", "HTMLElement", InitializeHTMLAnchorElement, HTMLAnchorElementConstructor)
	js.RegisterClass(e, "HTMLFormElement", "HTMLElement", InitializeHTMLFormElement, HTMLFormElementConstructor)
	js.RegisterClass(e, "HTMLInputElement", "HTMLElement", InitializeHTMLInputElement, HTMLInputElementConstructor)
	js.RegisterClass(e, "HTMLTemplateElement", "HTMLElement", InitializeHTMLTemplateElement, HTMLTemplateElementConstructor)
}
