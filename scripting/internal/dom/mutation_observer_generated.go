// This file is generated. Do not edit.

package dom

import (
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeMutationObserver[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("observe", MutationObserver_observe)
	jsClass.CreateOperation("disconnect", MutationObserver_disconnect)
	jsClass.CreateOperation("takeRecords", MutationObserver_takeRecords)
}

func MutationObserverConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	callback, errArg1 := js.ConsumeArgument(cbCtx, "callback", nil, decodeMutationCallback)
	if errArg1 != nil {
		return nil, errArg1
	}
	return CreateMutationObserver(cbCtx, callback)
}

func MutationObserver_observe[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	target, errArg1 := js.ConsumeArgument(cbCtx, "target", nil, codec.DecodeNode)
	options, errArg2 := js.ConsumeArgument(cbCtx, "options", nil, decodeObserveOption)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	errCall := instance.Observe(target, options...)
	return nil, errCall
}

func MutationObserver_disconnect[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Disconnect()
	return nil, nil
}

func MutationObserver_takeRecords[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.MutationObserver](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.TakeRecords()
	return encodeSequenceMutationRecord(cbCtx, result)
}
