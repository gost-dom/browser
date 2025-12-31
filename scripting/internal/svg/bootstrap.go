package svg

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "SVGElement", "Element", InitializeSvgElement, js.IllegalConstructor)
}
