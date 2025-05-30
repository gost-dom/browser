// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

type htmlHyperlinkElementUtilsV8Wrapper struct {
	handleReffedObject[html.HTMLHyperlinkElementUtils, jsTypeParam]
}

func newHTMLHyperlinkElementUtilsV8Wrapper(scriptHost *V8ScriptHost) *htmlHyperlinkElementUtilsV8Wrapper {
	return &htmlHyperlinkElementUtilsV8Wrapper{newHandleReffedObject[html.HTMLHyperlinkElementUtils](scriptHost)}
}

func createHTMLHyperlinkElementUtilsPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newHTMLHyperlinkElementUtilsV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlHyperlinkElementUtilsV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {

	prototypeTmpl.SetAccessorProperty("href",
		wrapV8Callback(w.scriptHost, w.href),
		wrapV8Callback(w.scriptHost, w.setHref),
		v8.None)
	prototypeTmpl.Set("toString", wrapV8Callback(w.scriptHost, w.href))
	prototypeTmpl.SetAccessorProperty("origin",
		wrapV8Callback(w.scriptHost, w.origin),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("protocol",
		wrapV8Callback(w.scriptHost, w.protocol),
		wrapV8Callback(w.scriptHost, w.setProtocol),
		v8.None)
	prototypeTmpl.SetAccessorProperty("username",
		wrapV8Callback(w.scriptHost, w.username),
		wrapV8Callback(w.scriptHost, w.setUsername),
		v8.None)
	prototypeTmpl.SetAccessorProperty("password",
		wrapV8Callback(w.scriptHost, w.password),
		wrapV8Callback(w.scriptHost, w.setPassword),
		v8.None)
	prototypeTmpl.SetAccessorProperty("host",
		wrapV8Callback(w.scriptHost, w.host),
		wrapV8Callback(w.scriptHost, w.setHost),
		v8.None)
	prototypeTmpl.SetAccessorProperty("hostname",
		wrapV8Callback(w.scriptHost, w.hostname),
		wrapV8Callback(w.scriptHost, w.setHostname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("port",
		wrapV8Callback(w.scriptHost, w.port),
		wrapV8Callback(w.scriptHost, w.setPort),
		v8.None)
	prototypeTmpl.SetAccessorProperty("pathname",
		wrapV8Callback(w.scriptHost, w.pathname),
		wrapV8Callback(w.scriptHost, w.setPathname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("search",
		wrapV8Callback(w.scriptHost, w.search),
		wrapV8Callback(w.scriptHost, w.setSearch),
		v8.None)
	prototypeTmpl.SetAccessorProperty("hash",
		wrapV8Callback(w.scriptHost, w.hash),
		wrapV8Callback(w.scriptHost, w.setHash),
		v8.None)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlHyperlinkElementUtilsV8Wrapper) href(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.href")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Href()
	return w.toString_(cbCtx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setHref(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setHref")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetHref(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) origin(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.origin")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Origin()
	return w.toString_(cbCtx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) protocol(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.protocol")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Protocol()
	return w.toString_(cbCtx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setProtocol(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setProtocol")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetProtocol(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) username(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.username")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Username()
	return w.toString_(cbCtx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setUsername(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setUsername")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetUsername(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) password(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.password")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Password()
	return w.toString_(cbCtx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setPassword(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setPassword")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetPassword(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) host(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.host")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Host()
	return w.toString_(cbCtx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setHost(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setHost")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetHost(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) hostname(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.hostname")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Hostname()
	return w.toString_(cbCtx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setHostname(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setHostname")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetHostname(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) port(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.port")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Port()
	return w.toString_(cbCtx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setPort(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setPort")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetPort(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) pathname(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.pathname")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Pathname()
	return w.toString_(cbCtx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setPathname(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setPathname")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetPathname(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) search(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.search")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Search()
	return w.toString_(cbCtx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setSearch(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setSearch")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetSearch(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) hash(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.hash")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Hash()
	return w.toString_(cbCtx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setHash(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setHash")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetHash(val)
	return cbCtx.ReturnWithValue(nil)
}
