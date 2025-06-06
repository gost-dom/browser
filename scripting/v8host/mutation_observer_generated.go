// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type mutationObserverV8Wrapper[T any] struct{}

func newMutationObserverV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *mutationObserverV8Wrapper[T] {
	return &mutationObserverV8Wrapper[T]{}
}

func (wrapper mutationObserverV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w mutationObserverV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("observe", w.observe)
	jsClass.CreatePrototypeMethod("disconnect", w.disconnect)
	jsClass.CreatePrototypeMethod("takeRecords", w.takeRecords)
}

func (w mutationObserverV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationObserver.Constructor")
	callback, errArg1 := consumeArgument(cbCtx, "callback", nil, w.decodeMutationCallback)
	if errArg1 != nil {
		return nil, errArg1
	}
	return w.CreateInstance(cbCtx, callback)
}

func (w mutationObserverV8Wrapper[T]) observe(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationObserver.observe")
	instance, errInst := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	target, errArg1 := consumeArgument(cbCtx, "target", nil, codec.DecodeNode)
	options, errArg2 := consumeArgument(cbCtx, "options", nil, w.decodeObserveOption)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	errCall := instance.Observe(target, options...)
	return nil, errCall
}

func (w mutationObserverV8Wrapper[T]) disconnect(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationObserver.disconnect")
	instance, err := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Disconnect()
	return nil, nil
}

func (w mutationObserverV8Wrapper[T]) takeRecords(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationObserver.takeRecords")
	instance, err := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.TakeRecords()
	return w.toSequenceMutationRecord(cbCtx, result)
}
