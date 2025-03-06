package v8host

import (
	"github.com/gost-dom/browser/dom"

	v8 "github.com/gost-dom/v8go"
)

type documentFragmentV8Wrapper struct {
	handleReffedObject[dom.DocumentFragment]
	parentNode *parentNodeV8Wrapper
}

func (w documentFragmentV8Wrapper) constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.scriptHost.mustGetContext(info.Context())
	result := dom.NewDocumentFragment(ctx.window.Document())
	_, err := w.store(result, ctx, info.This())
	return nil, err
}

func createDocumentFragmentPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := documentFragmentV8Wrapper{
		newHandleReffedObject[dom.DocumentFragment](host),
		newParentNodeV8Wrapper(host),
	}
	constructor := v8.NewFunctionTemplateWithError(host.iso, wrapper.constructor)
	constructor.InstanceTemplate().SetInternalFieldCount(1)
	wrapper.parentNode.installPrototype(constructor.PrototypeTemplate())
	return constructor
}
