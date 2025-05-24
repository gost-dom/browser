// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	log "github.com/gost-dom/browser/internal/log"
	abstraction "github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("MutationObserver", "", createMutationObserverPrototype)
}

type mutationObserverV8Wrapper struct {
	handleReffedObject[dominterfaces.MutationObserver]
}

func newMutationObserverV8Wrapper(scriptHost *V8ScriptHost) *mutationObserverV8Wrapper {
	return &mutationObserverV8Wrapper{newHandleReffedObject[dominterfaces.MutationObserver](scriptHost)}
}

func createMutationObserverPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newMutationObserverV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w mutationObserverV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("observe", v8.NewFunctionTemplateWithError(iso, w.observe))
	prototypeTmpl.Set("disconnect", v8.NewFunctionTemplateWithError(iso, w.disconnect))
	prototypeTmpl.Set("takeRecords", v8.NewFunctionTemplateWithError(iso, w.takeRecords))
}

func (w mutationObserverV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationObserver.Constructor")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	callback, err1 := tryParseArg(args, 0, w.decodeMutationCallback)
	if args.noOfReadArguments >= 1 {
		if err1 != nil {
			return nil, err1
		}
		return w.CreateInstance(cbCtx.Context(), info.This(), callback)
	}
	return nil, errors.New("MutationObserver.constructor: Missing arguments")
}

func (w mutationObserverV8Wrapper) observe(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationObserver.observe")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dominterfaces.MutationObserver](cbCtx.Instance())
	target, err1 := tryParseArg(args, 0, w.decodeNode)
	options, err2 := tryParseArg(args, 1, w.decodeMutationObserverInit)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		callErr := instance.Observe(target, options...)
		return nil, callErr
	}
	return nil, errors.New("MutationObserver.observe: Missing arguments")
}

func (w mutationObserverV8Wrapper) disconnect(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationObserver.disconnect")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Disconnect()
	return nil, nil
}

func (w mutationObserverV8Wrapper) takeRecords(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationObserver.takeRecords")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.TakeRecords()
	return w.toSequenceMutationRecord(cbCtx.Context(), result)
}
