// This file is generated. Do not edit.

package v8host

import (
	"errors"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("History", "", createHistoryPrototype)
}

type historyV8Wrapper struct {
	handleReffedObject[htmlinterfaces.History, jsTypeParam]
}

func newHistoryV8Wrapper(scriptHost *V8ScriptHost) *historyV8Wrapper {
	return &historyV8Wrapper{newHandleReffedObject[htmlinterfaces.History](scriptHost)}
}

func createHistoryPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newHistoryV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor)

	return constructor
}
func (w historyV8Wrapper) installPrototype(ft *v8.FunctionTemplate) {
	jsClass := newV8Class(w.scriptHost, ft)
	jsClass.CreatePrototypeMethod("go", w.go_)
	jsClass.CreatePrototypeMethod("back", w.back)
	jsClass.CreatePrototypeMethod("forward", w.forward)
	jsClass.CreatePrototypeMethod("pushState", w.pushState)
	jsClass.CreatePrototypeMethod("replaceState", w.replaceState)
	prototypeTmpl := ft.PrototypeTemplate()
	prototypeTmpl.SetAccessorProperty("length",
		wrapV8Callback(w.scriptHost, w.length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("state",
		wrapV8Callback(w.scriptHost, w.state),
		nil,
		v8.None)
}

func (w historyV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: History.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w historyV8Wrapper) go_(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: History.go_")
	instance, errInst := js.As[htmlinterfaces.History](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	delta, errArg1 := consumeArgument(cbCtx, "delta", w.defaultDelta, w.decodeLong)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Go(delta)
	return nil, errCall
}

func (w historyV8Wrapper) back(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: History.back")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	errCall := instance.Back()
	return nil, errCall
}

func (w historyV8Wrapper) forward(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: History.forward")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	errCall := instance.Forward()
	return nil, errCall
}

func (w historyV8Wrapper) pushState(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: History.pushState")
	instance, errInst := js.As[htmlinterfaces.History](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	data, errArg1 := consumeArgument(cbCtx, "data", nil, w.decodeHistoryState)
	cbCtx.ConsumeArg()
	url, errArg3 := consumeArgument(cbCtx, "url", w.defaultUrl, w.decodeString)
	err := errors.Join(errArg1, errArg3)
	if err != nil {
		return nil, err
	}
	errCall := instance.PushState(data, url)
	return nil, errCall
}

func (w historyV8Wrapper) replaceState(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: History.replaceState")
	instance, errInst := js.As[htmlinterfaces.History](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	data, errArg1 := consumeArgument(cbCtx, "data", nil, w.decodeHistoryState)
	cbCtx.ConsumeArg()
	url, errArg3 := consumeArgument(cbCtx, "url", w.defaultUrl, w.decodeString)
	err := errors.Join(errArg1, errArg3)
	if err != nil {
		return nil, err
	}
	errCall := instance.ReplaceState(data, url)
	return nil, errCall
}

func (w historyV8Wrapper) length(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: History.length")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx, result)
}

func (w historyV8Wrapper) state(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: History.state")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.State()
	return w.toHistoryState(cbCtx, result)
}
