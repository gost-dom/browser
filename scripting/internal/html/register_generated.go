// This file is generated. Do not edit.

package html

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "History", "", History[T]{}.Initialize, HistoryConstructor)
	js.RegisterClass(e, "Location", "", Location[T]{}.Initialize, LocationConstructor)
	js.RegisterClass(e, "HTMLElement", "Element", HTMLElement[T]{}.Initialize, HTMLElementConstructor)
	js.RegisterClass(e, "HTMLAnchorElement", "HTMLElement", HTMLAnchorElement[T]{}.Initialize, HTMLAnchorElementConstructor)
	js.RegisterClass(e, "HTMLFormElement", "HTMLElement", HTMLFormElement[T]{}.Initialize, HTMLFormElementConstructor)
	js.RegisterClass(e, "HTMLInputElement", "HTMLElement", HTMLInputElement[T]{}.Initialize, HTMLInputElementConstructor)
	js.RegisterClass(e, "HTMLTemplateElement", "HTMLElement", HTMLTemplateElement[T]{}.Initialize, HTMLTemplateElementConstructor)
}
