// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
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
	wrapper := newMutationObserverV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w mutationObserverV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("observe", wrapV8Callback(w.scriptHost, w.observe))
	prototypeTmpl.Set("disconnect", wrapV8Callback(w.scriptHost, w.disconnect))
	prototypeTmpl.Set("takeRecords", wrapV8Callback(w.scriptHost, w.takeRecords))
}

func (w mutationObserverV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationObserver.Constructor")
	callback, err1 := consumeArgument(cbCtx, "callback", nil, w.decodeMutationCallback)
	if cbCtx.noOfReadArguments >= 1 {
		if err1 != nil {
			return cbCtx.ReturnWithError(err1)
		}
		return w.CreateInstance(cbCtx, callback)
	}
	return cbCtx.ReturnWithError(errors.New("MutationObserver.constructor: Missing arguments"))
}

func (w mutationObserverV8Wrapper) observe(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationObserver.observe")
	instance, err0 := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	target, err1 := consumeArgument(cbCtx, "target", nil, w.decodeNode)
	options, err2 := consumeArgument(cbCtx, "options", nil, w.decodeObserveOption)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		callErr := instance.Observe(target, options...)
		if callErr != nil {
			return cbCtx.ReturnWithError(callErr)
		}
		return cbCtx.ReturnWithValue(nil)
	}
	return cbCtx.ReturnWithError(errors.New("MutationObserver.observe: Missing arguments"))
}

func (w mutationObserverV8Wrapper) disconnect(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationObserver.disconnect")
	instance, err := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Disconnect()
	return cbCtx.ReturnWithValue(nil)
}

func (w mutationObserverV8Wrapper) takeRecords(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationObserver.takeRecords")
	instance, err := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.TakeRecords()
	return w.toSequenceMutationRecord(cbCtx, result)
}
