// This file is generated. Do not edit.

package html

import (
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeMessageChannel[T any](jsClass js.Class[T]) {
	jsClass.CreateAttribute("port1", MessageChannel_port1, nil)
	jsClass.CreateAttribute("port2", MessageChannel_port2, nil)
}

func MessageChannelConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return CreateMessageChannel(cbCtx)
}

func MessageChannel_port1[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[htmlinterfaces.MessageChannel](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Port1()
	return encodeMessagePort(cbCtx, result)
}

func MessageChannel_port2[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[htmlinterfaces.MessageChannel](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Port2()
	return encodeMessagePort(cbCtx, result)
}
