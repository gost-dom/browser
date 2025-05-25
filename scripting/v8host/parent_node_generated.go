// This file is generated. Do not edit.

package v8host

import (
	"errors"

	dom "github.com/gost-dom/browser/dom"
	log "github.com/gost-dom/browser/internal/log"
	abstraction "github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	v8 "github.com/gost-dom/v8go"
)

type parentNodeV8Wrapper struct {
	handleReffedObject[dom.ParentNode]
}

func newParentNodeV8Wrapper(scriptHost *V8ScriptHost) *parentNodeV8Wrapper {
	return &parentNodeV8Wrapper{newHandleReffedObject[dom.ParentNode](scriptHost)}
}

func createParentNodePrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newParentNodeV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w parentNodeV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("prepend", wrapV8Callback(w.scriptHost, w.prepend))
	prototypeTmpl.Set("append", wrapV8Callback(w.scriptHost, w.append))
	prototypeTmpl.Set("replaceChildren", wrapV8Callback(w.scriptHost, w.replaceChildren))
	prototypeTmpl.Set("querySelector", wrapV8Callback(w.scriptHost, w.querySelector))
	prototypeTmpl.Set("querySelectorAll", wrapV8Callback(w.scriptHost, w.querySelectorAll))

	prototypeTmpl.SetAccessorProperty("firstElementChild",
		wrapV8Callback(w.scriptHost, w.firstElementChild),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("lastElementChild",
		wrapV8Callback(w.scriptHost, w.lastElementChild),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("childElementCount",
		wrapV8Callback(w.scriptHost, w.childElementCount),
		nil,
		v8.None)
}

func (w parentNodeV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.Constructor")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w parentNodeV8Wrapper) querySelector(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.querySelector")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.ParentNode](cbCtx.Instance())
	selectors, err1 := consumeArgument(cbCtx, "selectors", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.QuerySelector(selectors)
		if callErr != nil {
			return nil, callErr
		} else {
			return cbCtx.Context().getInstanceForNode(result)
		}
	}
	return nil, errors.New("ParentNode.querySelector: Missing arguments")
}

func (w parentNodeV8Wrapper) querySelectorAll(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.querySelectorAll")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.ParentNode](cbCtx.Instance())
	selectors, err1 := consumeArgument(cbCtx, "selectors", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.QuerySelectorAll(selectors)
		if callErr != nil {
			return nil, callErr
		} else {
			return w.toNodeList(cbCtx.Context(), result)
		}
	}
	return nil, errors.New("ParentNode.querySelectorAll: Missing arguments")
}

func (w parentNodeV8Wrapper) firstElementChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.firstElementChild")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.FirstElementChild()
	return cbCtx.Context().getInstanceForNode(result)
}

func (w parentNodeV8Wrapper) lastElementChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.lastElementChild")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LastElementChild()
	return cbCtx.Context().getInstanceForNode(result)
}

func (w parentNodeV8Wrapper) childElementCount(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.childElementCount")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ChildElementCount()
	return w.toUnsignedLong(cbCtx.Context(), result)
}
