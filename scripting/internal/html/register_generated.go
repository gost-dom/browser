// This file is generated. Do not edit.

package html

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "History", "", NewHistory, HistoryConstructor)
	js.RegisterClass(e, "Location", "", NewLocation, LocationConstructor)
	js.RegisterClass(e, "HTMLElement", "Element", NewHTMLElement, HTMLElementConstructor)
	js.RegisterClass(e, "HTMLAnchorElement", "HTMLElement", NewHTMLAnchorElement, HTMLAnchorElementConstructor)
	js.RegisterClass(e, "HTMLFormElement", "HTMLElement", NewHTMLFormElement, HTMLFormElementConstructor)
	js.RegisterClass(e, "HTMLInputElement", "HTMLElement", NewHTMLInputElement, HTMLInputElementConstructor)
	js.RegisterClass(e, "HTMLTemplateElement", "HTMLElement", NewHTMLTemplateElement, HTMLTemplateElementConstructor)
}
