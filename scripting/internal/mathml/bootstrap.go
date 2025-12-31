package mathml

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	InitializeMathMlElement[T](js.CreateClass(e, "MathMLElement", "Element", js.IllegalConstructor))
}
