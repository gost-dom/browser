package v8host

import js "github.com/gost-dom/browser/scripting/internal/js"

type htmlDocumentV8Wrapper struct {
	documentV8Wrapper[jsTypeParam]
}

func newHTMLDocumentV8Wrapper(host js.ScriptEngine[jsTypeParam]) htmlDocumentV8Wrapper {
	return htmlDocumentV8Wrapper{*newDocumentV8Wrapper(host)}
}

func (w htmlDocumentV8Wrapper) constructor(c jsCallbackContext) (jsValue, error) {
	return nil, c.ValueFactory().NewTypeError("illegal constructor")
}

func (w htmlDocumentV8Wrapper) initialize(c jsClass) {
	w.documentV8Wrapper.Initialize(c)
}
