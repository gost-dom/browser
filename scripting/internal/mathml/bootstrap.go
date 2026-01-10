package mathml

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	// The MathMLElement is added for Datastar support, which checks an element
	// is an instanceof MathMLElement. Would throw a JavaScript error if not
	// present in global scope.
	InitializeMathMlElement(js.CreateClass(e, "MathMLElement", "Element", js.IllegalConstructor))
}
