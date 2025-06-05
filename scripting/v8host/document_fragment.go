package v8host

import (
	"github.com/gost-dom/browser/dom"
)

type documentFragmentV8Wrapper struct {
	handleReffedObject[dom.DocumentFragment, jsTypeParam]
	parentNode *parentNodeV8Wrapper
}

func newDocumentFragmentV8Wrapper(host *V8ScriptHost) documentFragmentV8Wrapper {
	return documentFragmentV8Wrapper{
		newHandleReffedObject[dom.DocumentFragment](host),
		newParentNodeV8Wrapper(host),
	}
}

func (w documentFragmentV8Wrapper) constructor(ctx jsCallbackContext) (jsValue, error) {
	result := dom.NewDocumentFragment(ctx.Scope().Window().Document())
	return w.store(result, ctx)
}

func (w documentFragmentV8Wrapper) initialize(class jsClass) {
	w.parentNode.installPrototype(class)
}
