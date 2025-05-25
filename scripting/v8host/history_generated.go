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
	handleReffedObject[htmlinterfaces.History]
}

func newHistoryV8Wrapper(scriptHost *V8ScriptHost) *historyV8Wrapper {
	return &historyV8Wrapper{newHandleReffedObject[htmlinterfaces.History](scriptHost)}
}

func createHistoryPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newHistoryV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w historyV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("go", wrapV8Callback(w.scriptHost, w.go_))
	prototypeTmpl.Set("back", wrapV8Callback(w.scriptHost, w.back))
	prototypeTmpl.Set("forward", wrapV8Callback(w.scriptHost, w.forward))
	prototypeTmpl.Set("pushState", wrapV8Callback(w.scriptHost, w.pushState))
	prototypeTmpl.Set("replaceState", wrapV8Callback(w.scriptHost, w.replaceState))

	prototypeTmpl.SetAccessorProperty("length",
		wrapV8Callback(w.scriptHost, w.length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("state",
		wrapV8Callback(w.scriptHost, w.state),
		nil,
		v8.None)
}

func (w historyV8Wrapper) Constructor(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: History.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w historyV8Wrapper) go_(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: History.go_")
	instance, err0 := js.As[htmlinterfaces.History](cbCtx.Instance())
	delta, err1 := consumeArgument(cbCtx, "delta", w.defaultDelta, w.decodeLong)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		callErr := instance.Go(delta)
		if callErr != nil {
			return cbCtx.ReturnWithError(callErr)
		}
		return cbCtx.ReturnWithValue(nil)
	}
	return cbCtx.ReturnWithError(errors.New("History.go: Missing arguments"))
}

func (w historyV8Wrapper) back(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: History.back")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	callErr := instance.Back()
	if callErr != nil {
		return cbCtx.ReturnWithError(callErr)
	}
	return cbCtx.ReturnWithValue(nil)
}

func (w historyV8Wrapper) forward(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: History.forward")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	callErr := instance.Forward()
	if callErr != nil {
		return cbCtx.ReturnWithError(callErr)
	}
	return cbCtx.ReturnWithValue(nil)
}

func (w historyV8Wrapper) pushState(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: History.pushState")
	instance, err0 := js.As[htmlinterfaces.History](cbCtx.Instance())
	data, err1 := consumeArgument(cbCtx, "data", nil, w.decodeAny)
	ignoreArgument(cbCtx)
	url, err3 := consumeArgument(cbCtx, "url", w.defaultUrl, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err3)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		callErr := instance.PushState(data, url)
		if callErr != nil {
			return cbCtx.ReturnWithError(callErr)
		}
		return cbCtx.ReturnWithValue(nil)
	}
	return cbCtx.ReturnWithError(errors.New("History.pushState: Missing arguments"))
}

func (w historyV8Wrapper) replaceState(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: History.replaceState")
	instance, err0 := js.As[htmlinterfaces.History](cbCtx.Instance())
	data, err1 := consumeArgument(cbCtx, "data", nil, w.decodeAny)
	ignoreArgument(cbCtx)
	url, err3 := consumeArgument(cbCtx, "url", w.defaultUrl, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err3)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		callErr := instance.ReplaceState(data, url)
		if callErr != nil {
			return cbCtx.ReturnWithError(callErr)
		}
		return cbCtx.ReturnWithValue(nil)
	}
	return cbCtx.ReturnWithError(errors.New("History.replaceState: Missing arguments"))
}

func (w historyV8Wrapper) length(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: History.length")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx, result)
}

func (w historyV8Wrapper) state(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: History.state")
	instance, err := js.As[htmlinterfaces.History](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.State()
	return w.toHistoryState(cbCtx, result)
}
