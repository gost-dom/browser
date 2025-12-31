// This file is generated. Do not edit.

package html

import (
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type History[T any] struct{}

func NewHistory[T any](scriptHost js.ScriptEngine[T]) History[T] {
	return History[T]{}
}

func (wrapper History[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w History[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("go", History_go)
	jsClass.CreateOperation("back", History_back)
	jsClass.CreateOperation("forward", History_forward)
	jsClass.CreateOperation("pushState", History_pushState)
	jsClass.CreateOperation("replaceState", History_replaceState)
	jsClass.CreateAttribute("length", History_length, nil)
	jsClass.CreateAttribute("state", History_state, nil)
}

func HistoryConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func History_go[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[htmlinterfaces.History](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	delta, errArg1 := js.ConsumeArgument(cbCtx, "delta", codec.ZeroValue, codec.DecodeInt)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Go(delta)
	return nil, errCall
}

func History_back[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.Back()
	return nil, errCall
}

func History_forward[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.Forward()
	return nil, errCall
}

func History_pushState[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[htmlinterfaces.History](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	data, errArg1 := js.ConsumeArgument(cbCtx, "data", nil, decodeHistoryState)
	cbCtx.ConsumeArg()
	url, errArg3 := js.ConsumeArgument(cbCtx, "url", codec.ZeroValue, codec.DecodeString)
	err = gosterror.First(errArg1, errArg3)
	if err != nil {
		return nil, err
	}
	errCall := instance.PushState(data, url)
	return nil, errCall
}

func History_replaceState[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[htmlinterfaces.History](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	data, errArg1 := js.ConsumeArgument(cbCtx, "data", nil, decodeHistoryState)
	cbCtx.ConsumeArg()
	url, errArg3 := js.ConsumeArgument(cbCtx, "url", codec.ZeroValue, codec.DecodeString)
	err = gosterror.First(errArg1, errArg3)
	if err != nil {
		return nil, err
	}
	errCall := instance.ReplaceState(data, url)
	return nil, errCall
}

func History_length[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}

func History_state[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.State()
	return encodeHistoryState(cbCtx, result)
}
