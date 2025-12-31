package dom

import "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	installDOMParser(e)
	Register(e)
	js.CreateClass(e, "ShadowRoot", "DocumentFragment", shadowRootConstructor)
	InitializeCustomEvent(js.CreateClass(e, "CustomEvent", "Event", customEventConstructor))
}
