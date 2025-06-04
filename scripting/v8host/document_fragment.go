package v8host

import (
	"github.com/gost-dom/browser/dom"
)

type documentFragmentV8Wrapper struct {
	handleReffedObject[dom.DocumentFragment, jsTypeParam]
	parentNode *parentNodeV8Wrapper
}

func newDocumentFragmentV8Wrapper(host *V8ScriptHost) jsInitializer {
	return documentFragmentV8Wrapper{
		newHandleReffedObject[dom.DocumentFragment](host),
		newParentNodeV8Wrapper(host),
	}
}

func createDocumentFragmentPrototype(host *V8ScriptHost) v8Class {
	wrapper := newDocumentFragmentV8Wrapper(host)
	constructor := wrapV8Callback(host, wrapper.constructor)
	constructor.InstanceTemplate().SetInternalFieldCount(1)
	jsClass := newV8Class(host, constructor)
	wrapper.initialize(jsClass)
	return jsClass
}

func (w documentFragmentV8Wrapper) constructor(ctx jsCallbackContext) (jsValue, error) {
	result := dom.NewDocumentFragment(ctx.Scope().Window().Document())
	return w.store(result, ctx)
}

func (w documentFragmentV8Wrapper) initialize(class v8Class) {
	w.parentNode.installPrototype(class)
}
