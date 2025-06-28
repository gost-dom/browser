// This file is generated. Do not edit.

package dom

import (
	"errors"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type MutationObserver[T any] struct{}

func NewMutationObserver[T any](scriptHost js.ScriptEngine[T]) *MutationObserver[T] {
	return &MutationObserver[T]{}
}

func (wrapper MutationObserver[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w MutationObserver[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("observe", w.observe)
	jsClass.CreatePrototypeMethod("disconnect", w.disconnect)
	jsClass.CreatePrototypeMethod("takeRecords", w.takeRecords)
}

func (w MutationObserver[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: MutationObserver.Constructor")
	callback, errArg1 := js.ConsumeArgument(cbCtx, "callback", nil, w.decodeMutationCallback)
	if errArg1 != nil {
		return nil, errArg1
	}
	return w.CreateInstance(cbCtx, callback)
}

func (w MutationObserver[T]) observe(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: MutationObserver.observe")
	instance, errInst := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	target, errArg1 := js.ConsumeArgument(cbCtx, "target", nil, codec.DecodeNode)
	options, errArg2 := js.ConsumeArgument(cbCtx, "options", nil, w.decodeObserveOption)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	errCall := instance.Observe(target, options...)
	return nil, errCall
}

func (w MutationObserver[T]) disconnect(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: MutationObserver.disconnect")
	instance, err := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Disconnect()
	return nil, nil
}

func (w MutationObserver[T]) takeRecords(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: MutationObserver.takeRecords")
	instance, err := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.TakeRecords()
	return w.toSequenceMutationRecord(cbCtx, result)
}
