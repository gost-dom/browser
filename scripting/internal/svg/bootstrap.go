package svg

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	// The SVGElement is added for Datastar support, which checks an element
	// is an instanceof SVGElement. Would throw a JavaScript error if not
	// present in global scope.
	InitializeSvgElement(js.CreateClass(e, "SVGElement", "Element", js.IllegalConstructor))
}
