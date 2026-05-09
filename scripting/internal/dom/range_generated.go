// This file is generated. Do not edit.

package dom

import (
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeRange[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("setStart", Range_setStart)
	jsClass.CreateOperation("setEnd", Range_setEnd)
	jsClass.CreateOperation("detach", Range_detach)
	jsClass.CreateOperation("toString", Range_toString)
}

func Range_setStart[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	offset, errArg2 := js.ConsumeArgument(cbCtx, "offset", nil, codec.DecodeInt)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	errCall := instance.SetStart(node, offset)
	return nil, errCall
}

func Range_setEnd[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	offset, errArg2 := js.ConsumeArgument(cbCtx, "offset", nil, codec.DecodeInt)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	errCall := instance.SetEnd(node, offset)
	return nil, errCall
}

func Range_detach[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.Range](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Detach()
	return nil, nil
}

func Range_toString[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.Range](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.String()
	return codec.EncodeString(cbCtx, result)
}
