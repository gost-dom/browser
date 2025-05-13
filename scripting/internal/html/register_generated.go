// This file is generated. Do not edit.

package html

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "HTMLAnchorElement", "HTMLElement", NewHTMLAnchorElement)
	js.RegisterClass(reg, "HTMLElement", "Element", NewHTMLElement)
	js.RegisterClass(reg, "HTMLFormElement", "HTMLElement", NewHTMLFormElement)
	js.RegisterClass(reg, "HTMLHyperlinkElementUtils", "", NewHTMLHyperlinkElementUtils)
	js.RegisterClass(reg, "HTMLInputElement", "HTMLElement", NewHTMLInputElement)
	js.RegisterClass(reg, "HTMLOrSVGElement", "", NewHTMLOrSVGElement)
	js.RegisterClass(reg, "HTMLTemplateElement", "HTMLElement", NewHTMLTemplateElement)
	js.RegisterClass(reg, "History", "", NewHistory)
	js.RegisterClass(reg, "Location", "", NewLocation)
	js.RegisterClass(reg, "Window", "EventTarget", NewWindow)
}

func InitBuilder[T any](reg js.ClassBuilder[T]) {
	Bootstrap[T](reg)
}
