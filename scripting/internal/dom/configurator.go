package dom

import "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureErrors[T any](e js.ScriptEngine[T]) {
	InitializeDomException(js.CreateClass(e, "DOMException", "", domExceptionConstructor))
}

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	installDOMParser(e)
	js.CreateClass(e, "ShadowRoot", "DocumentFragment", shadowRootConstructor)
}
