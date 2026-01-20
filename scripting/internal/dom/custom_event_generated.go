// This file is generated. Do not edit.

package dom

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeCustomEvent[T any](jsClass js.Class[T]) {
	jsClass.CreateAttribute("detail", CustomEvent_detail, nil)
}

func CustomEventConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	type_, errType := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	options, errOpts := js.ConsumeArgument(cbCtx, "options", codec.ZeroValue, codec.DecodeJsObject)
	err = gosterror.First(errType, errOpts)
	if err != nil {
		return nil, err
	}
	var data event.CustomEventInit
	e := event.Event{Type: type_}
	if options != nil {
		err = errors.Join(
			codec.DecodeEvent(cbCtx, options, &e),
			decodeCustomEventInit(cbCtx, options, &data))
		if err != nil {
			return nil, err
		}
	}
	e.Data = data
	return codec.EncodeConstructedValue(cbCtx, &e)
}

func CustomEvent_detail[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var eventInit event.CustomEventInit
	eventInit, err = codec.RetrieveEventInit[event.CustomEventInit](cbCtx)
	if err != nil {
		return nil, err
	}
	result := eventInit.Detail
	return codec.EncodeAny(cbCtx, result)
}
