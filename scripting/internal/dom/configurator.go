package dom

import "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	installDOMParser(e)
	Register(e)
	js.RegisterClass(e, "ShadowRoot", "DocumentFragment", newShadowRoot, shadowRootConstructor)
	js.RegisterClass(e, "CustomEvent", "Event", NewCustomEvent, customEventConstructor)
}
