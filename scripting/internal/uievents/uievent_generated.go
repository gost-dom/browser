// This file is generated. Do not edit.

package uievents

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	uievents "github.com/gost-dom/browser/internal/uievents"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type UIEvent[T any] struct{}

func NewUIEvent[T any](scriptHost js.ScriptEngine[T]) UIEvent[T] {
	return UIEvent[T]{}
}

func (wrapper UIEvent[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w UIEvent[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateAttribute("view", UIEvent_view, nil)
	jsClass.CreateAttribute("detail", UIEvent_detail, nil)
}

func UIEventConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	type_, errType := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	options, errOpts := js.ConsumeArgument(cbCtx, "options", codec.ZeroValue, codec.DecodeJsObject)
	err = gosterror.First(errType, errOpts)
	if err != nil {
		return nil, err
	}
	var data uievents.UIEventInit
	e := event.Event{Type: type_}
	if options != nil {
		err = errors.Join(
			codec.DecodeEvent(cbCtx, options, &e),
			decodeUIEventInit(cbCtx, options, &data))
		if err != nil {
			return nil, err
		}
	}
	e.Data = data
	return codec.EncodeConstructedValue(cbCtx, &e)
}

func UIEvent_view[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "UIEvent.UIEvent_view: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func UIEvent_detail[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "UIEvent.UIEvent_detail: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
