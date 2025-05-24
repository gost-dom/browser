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
	iso := w.scriptHost.iso
	prototypeTmpl.Set("prepend", v8.NewFunctionTemplateWithError(iso, w.prepend))
	prototypeTmpl.Set("append", v8.NewFunctionTemplateWithError(iso, w.append))
	prototypeTmpl.Set("replaceChildren", v8.NewFunctionTemplateWithError(iso, w.replaceChildren))
	prototypeTmpl.Set("querySelector", v8.NewFunctionTemplateWithError(iso, w.querySelector))
	prototypeTmpl.Set("querySelectorAll", v8.NewFunctionTemplateWithError(iso, w.querySelectorAll))

	prototypeTmpl.SetAccessorProperty("firstElementChild",
		v8.NewFunctionTemplateWithError(iso, w.firstElementChild),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("lastElementChild",
		v8.NewFunctionTemplateWithError(iso, w.lastElementChild),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("childElementCount",
		v8.NewFunctionTemplateWithError(iso, w.childElementCount),
		nil,
		v8.None)
}

func (w parentNodeV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.Constructor")
	args := newArgumentHelper(w.scriptHost, info)
	return args.ReturnWithTypeError("Illegal constructor")
}

func (w parentNodeV8Wrapper) querySelector(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.querySelector")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.ParentNode](args.Instance())
	selectors, err1 := tryParseArg(args, 0, w.decodeString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.QuerySelector(selectors)
		if callErr != nil {
			return nil, callErr
		} else {
			return args.Context().getInstanceForNode(result)
		}
	}
	return nil, errors.New("ParentNode.querySelector: Missing arguments")
}

func (w parentNodeV8Wrapper) querySelectorAll(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.querySelectorAll")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.ParentNode](args.Instance())
	selectors, err1 := tryParseArg(args, 0, w.decodeString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.QuerySelectorAll(selectors)
		if callErr != nil {
			return nil, callErr
		} else {
			return w.toNodeList(args.Context(), result)
		}
	}
	return nil, errors.New("ParentNode.querySelectorAll: Missing arguments")
}

func (w parentNodeV8Wrapper) firstElementChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.firstElementChild")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.ParentNode](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.FirstElementChild()
	return args.Context().getInstanceForNode(result)
}

func (w parentNodeV8Wrapper) lastElementChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.lastElementChild")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.ParentNode](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LastElementChild()
	return args.Context().getInstanceForNode(result)
}

func (w parentNodeV8Wrapper) childElementCount(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: ParentNode.childElementCount")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.ParentNode](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ChildElementCount()
	return w.toUnsignedLong(args.Context(), result)
}
