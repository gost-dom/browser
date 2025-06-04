// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("MutationObserver", "", newMutationObserverV8Wrapper)
}

type mutationObserverV8Wrapper struct {
	handleReffedObject[dominterfaces.MutationObserver, jsTypeParam]
}

func newMutationObserverV8Wrapper(scriptHost *V8ScriptHost) *mutationObserverV8Wrapper {
	return &mutationObserverV8Wrapper{newHandleReffedObject[dominterfaces.MutationObserver](scriptHost)}
}

func createMutationObserverPrototype(scriptHost *V8ScriptHost) v8Class {
	wrapper := newMutationObserverV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return jsClass
}
func (wrapper mutationObserverV8Wrapper) initialize(jsClass v8Class) {
	wrapper.installPrototype(jsClass)
}

func (w mutationObserverV8Wrapper) installPrototype(jsClass v8Class) {
	jsClass.CreatePrototypeMethod("observe", w.observe)
	jsClass.CreatePrototypeMethod("disconnect", w.disconnect)
	jsClass.CreatePrototypeMethod("takeRecords", w.takeRecords)
}

func (w mutationObserverV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationObserver.constructor")
	callback, errArg1 := consumeArgument(cbCtx, "callback", nil, w.decodeMutationCallback)
	if errArg1 != nil {
		return nil, errArg1
	}
	return w.CreateInstance(cbCtx, callback)
}

func (w mutationObserverV8Wrapper) observe(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationObserver.observe")
	instance, errInst := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	target, errArg1 := consumeArgument(cbCtx, "target", nil, w.decodeNode)
	options, errArg2 := consumeArgument(cbCtx, "options", nil, w.decodeObserveOption)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	errCall := instance.Observe(target, options...)
	return nil, errCall
}

func (w mutationObserverV8Wrapper) disconnect(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationObserver.disconnect")
	instance, err := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Disconnect()
	return nil, nil
}

func (w mutationObserverV8Wrapper) takeRecords(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationObserver.takeRecords")
	instance, err := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.TakeRecords()
	return w.toSequenceMutationRecord(cbCtx, result)
}
