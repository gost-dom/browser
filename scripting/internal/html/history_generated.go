// This file is generated. Do not edit.

package html

import (
	"errors"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type History[T any] struct{}

func NewHistory[T any](scriptHost js.ScriptEngine[T]) *History[T] {
	return &History[T]{}
}

func (wrapper History[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w History[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("go", w.go_)
	jsClass.CreatePrototypeMethod("back", w.back)
	jsClass.CreatePrototypeMethod("forward", w.forward)
	jsClass.CreatePrototypeMethod("pushState", w.pushState)
	jsClass.CreatePrototypeMethod("replaceState", w.replaceState)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
	jsClass.CreatePrototypeAttribute("state", w.state, nil)
}

func (w History[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: History.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w History[T]) go_(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: History.go_")
	instance, errInst := js.As[htmlinterfaces.History](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	delta, errArg1 := js.ConsumeArgument(cbCtx, "delta", w.defaultDelta, codec.DecodeInt)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Go(delta)
	return nil, errCall
}

func (w History[T]) back(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: History.back")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.Back()
	return nil, errCall
}

func (w History[T]) forward(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: History.forward")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.Forward()
	return nil, errCall
}

func (w History[T]) pushState(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: History.pushState")
	instance, errInst := js.As[htmlinterfaces.History](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	data, errArg1 := js.ConsumeArgument(cbCtx, "data", nil, w.decodeHistoryState)
	cbCtx.ConsumeArg()
	url, errArg3 := js.ConsumeArgument(cbCtx, "url", w.defaultUrl, codec.DecodeString)
	err := errors.Join(errArg1, errArg3)
	if err != nil {
		return nil, err
	}
	errCall := instance.PushState(data, url)
	return nil, errCall
}

func (w History[T]) replaceState(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: History.replaceState")
	instance, errInst := js.As[htmlinterfaces.History](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	data, errArg1 := js.ConsumeArgument(cbCtx, "data", nil, w.decodeHistoryState)
	cbCtx.ConsumeArg()
	url, errArg3 := js.ConsumeArgument(cbCtx, "url", w.defaultUrl, codec.DecodeString)
	err := errors.Join(errArg1, errArg3)
	if err != nil {
		return nil, err
	}
	errCall := instance.ReplaceState(data, url)
	return nil, errCall
}

func (w History[T]) length(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: History.length")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}

func (w History[T]) state(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: History.state")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.State()
	return w.toHistoryState(cbCtx, result)
}
