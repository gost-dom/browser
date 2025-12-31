package svg

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	InitializeSvgElement(js.CreateClass(e, "SVGElement", "Element", js.IllegalConstructor))
}
