package v8host

type htmlDocumentV8Wrapper struct {
	documentV8Wrapper
}

func newHTMLDocumentV8Wrapper(host *V8ScriptHost) htmlDocumentV8Wrapper {
	return htmlDocumentV8Wrapper{*newDocumentV8Wrapper(host)}
}

func (w htmlDocumentV8Wrapper) constructor(c jsCallbackContext) (jsValue, error) {
	return nil, c.ValueFactory().NewTypeError("illegal constructor")
}

func (w htmlDocumentV8Wrapper) initialize(c jsClass) {
	w.documentV8Wrapper.initialize(c)
}
