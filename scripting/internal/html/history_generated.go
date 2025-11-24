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

func (w History[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: History.Constructor - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: History.Constructor", js.LogAttr("res", res))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w History[T]) go_(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: History.go_ - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: History.go_", js.LogAttr("res", res))
	}()
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

func (w History[T]) back(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: History.back - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: History.back", js.LogAttr("res", res))
	}()
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.Back()
	return nil, errCall
}

func (w History[T]) forward(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: History.forward - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: History.forward", js.LogAttr("res", res))
	}()
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.Forward()
	return nil, errCall
}

func (w History[T]) pushState(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: History.pushState - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: History.pushState", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[htmlinterfaces.History](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	data, errArg1 := js.ConsumeArgument(cbCtx, "data", nil, w.decodeHistoryState)
	cbCtx.ConsumeArg()
	url, errArg3 := js.ConsumeArgument(cbCtx, "url", codec.ZeroValue, codec.DecodeString)
	err = errors.Join(errArg1, errArg3)
	if err != nil {
		return nil, err
	}
	errCall := instance.PushState(data, url)
	return nil, errCall
}

func (w History[T]) replaceState(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: History.replaceState - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: History.replaceState", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[htmlinterfaces.History](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	data, errArg1 := js.ConsumeArgument(cbCtx, "data", nil, w.decodeHistoryState)
	cbCtx.ConsumeArg()
	url, errArg3 := js.ConsumeArgument(cbCtx, "url", codec.ZeroValue, codec.DecodeString)
	err = errors.Join(errArg1, errArg3)
	if err != nil {
		return nil, err
	}
	errCall := instance.ReplaceState(data, url)
	return nil, errCall
}

func (w History[T]) length(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: History.length - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: History.length", js.LogAttr("res", res))
	}()
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}

func (w History[T]) state(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: History.state - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: History.state", js.LogAttr("res", res))
	}()
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.State()
	return w.toHistoryState(cbCtx, result)
}
