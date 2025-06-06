// This file is generated. Do not edit.

package html

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "HTMLAnchorElement", "HTMLElement", NewHTMLAnchorElementV8Wrapper)
	js.RegisterClass(reg, "HTMLElement", "Element", NewHTMLElementV8Wrapper)
	js.RegisterClass(reg, "HTMLFormElement", "HTMLElement", NewHTMLFormElementV8Wrapper)
	js.RegisterClass(reg, "HTMLHyperlinkElementUtils", "", NewHTMLHyperlinkElementUtilsV8Wrapper)
	js.RegisterClass(reg, "HTMLInputElement", "HTMLElement", NewHTMLInputElementV8Wrapper)
	js.RegisterClass(reg, "HTMLOrSVGElement", "", NewHTMLOrSVGElementV8Wrapper)
	js.RegisterClass(reg, "HTMLTemplateElement", "HTMLElement", NewHTMLTemplateElementV8Wrapper)
	js.RegisterClass(reg, "History", "", NewHistoryV8Wrapper)
	js.RegisterClass(reg, "Location", "", NewLocationV8Wrapper)
	js.RegisterClass(reg, "Window", "EventTarget", NewWindowV8Wrapper)
}
