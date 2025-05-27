// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

type htmlOrSVGElementV8Wrapper struct {
	handleReffedObject[html.HTMLOrSVGElement]
}

func newHTMLOrSVGElementV8Wrapper(scriptHost *V8ScriptHost) *htmlOrSVGElementV8Wrapper {
	return &htmlOrSVGElementV8Wrapper{newHandleReffedObject[html.HTMLOrSVGElement](scriptHost)}
}

func createHTMLOrSVGElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newHTMLOrSVGElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlOrSVGElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("focus", wrapV8Callback(w.scriptHost, w.focus))
	prototypeTmpl.Set("blur", wrapV8Callback(w.scriptHost, w.blur))

	prototypeTmpl.SetAccessorProperty("dataset",
		wrapV8Callback(w.scriptHost, w.dataset),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nonce",
		wrapV8Callback(w.scriptHost, w.nonce),
		wrapV8Callback(w.scriptHost, w.setNonce),
		v8.None)
	prototypeTmpl.SetAccessorProperty("autofocus",
		wrapV8Callback(w.scriptHost, w.autofocus),
		wrapV8Callback(w.scriptHost, w.setAutofocus),
		v8.None)
	prototypeTmpl.SetAccessorProperty("tabIndex",
		wrapV8Callback(w.scriptHost, w.tabIndex),
		wrapV8Callback(w.scriptHost, w.setTabIndex),
		v8.None)
}

func (w htmlOrSVGElementV8Wrapper) Constructor(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLOrSVGElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlOrSVGElementV8Wrapper) blur(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLOrSVGElement.blur")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Blur()
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlOrSVGElementV8Wrapper) dataset(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLOrSVGElement.dataset")
	return cbCtx.ReturnWithError(errors.New("HTMLOrSVGElement.dataset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlOrSVGElementV8Wrapper) nonce(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLOrSVGElement.nonce")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Nonce()
	return w.toString_(cbCtx, result)
}

func (w htmlOrSVGElementV8Wrapper) setNonce(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLOrSVGElement.setNonce")
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetNonce(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlOrSVGElementV8Wrapper) autofocus(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLOrSVGElement.autofocus")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Autofocus()
	return w.toBoolean(cbCtx, result)
}

func (w htmlOrSVGElementV8Wrapper) setAutofocus(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLOrSVGElement.setAutofocus")
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeBoolean)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetAutofocus(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlOrSVGElementV8Wrapper) tabIndex(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLOrSVGElement.tabIndex")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.TabIndex()
	return w.toLong(cbCtx, result)
}

func (w htmlOrSVGElementV8Wrapper) setTabIndex(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLOrSVGElement.setTabIndex")
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeLong)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetTabIndex(val)
	return cbCtx.ReturnWithValue(nil)
}
