package dom

import "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	installDOMParser(e)
	InitializeDomException(js.CreateClass(e, "DOMException", "", domExceptionConstructor))
	Bootstrap(e)
	js.CreateClass(e, "ShadowRoot", "DocumentFragment", shadowRootConstructor)
}
