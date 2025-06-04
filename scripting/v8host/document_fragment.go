package v8host

import (
	"github.com/gost-dom/browser/dom"
)

type documentFragmentV8Wrapper struct {
	handleReffedObject[dom.DocumentFragment, jsTypeParam]
	parentNode *parentNodeV8Wrapper
}

func (w documentFragmentV8Wrapper) constructor(ctx jsCallbackContext) (jsValue, error) {
	result := dom.NewDocumentFragment(ctx.Scope().Window().Document())
	return w.store(result, ctx)
}

func createDocumentFragmentPrototype(host *V8ScriptHost) v8Class {
	wrapper := documentFragmentV8Wrapper{
		newHandleReffedObject[dom.DocumentFragment](host),
		newParentNodeV8Wrapper(host),
	}
	constructor := wrapV8Callback(host, wrapper.constructor)
	constructor.InstanceTemplate().SetInternalFieldCount(1)
	jsClass := newV8Class(host, constructor)
	wrapper.parentNode.installPrototype(jsClass)
	return newV8Class(host, constructor)
}
