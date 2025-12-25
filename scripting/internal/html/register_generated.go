// This file is generated. Do not edit.

package html

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "History", "", NewHistory)
	js.RegisterClass(e, "Location", "", NewLocation)
	js.RegisterClass(e, "HTMLElement", "Element", NewHTMLElement)
	js.RegisterClass(e, "HTMLAnchorElement", "HTMLElement", NewHTMLAnchorElement)
	js.RegisterClass(e, "HTMLFormElement", "HTMLElement", NewHTMLFormElement)
	js.RegisterClass(e, "HTMLInputElement", "HTMLElement", NewHTMLInputElement)
	js.RegisterClass(e, "HTMLTemplateElement", "HTMLElement", NewHTMLTemplateElement)
}
