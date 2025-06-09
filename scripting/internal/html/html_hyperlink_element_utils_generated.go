// This file is generated. Do not edit.

package html

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLHyperlinkElementUtilsV8Wrapper[T any] struct{}

func NewHTMLHyperlinkElementUtilsV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *HTMLHyperlinkElementUtilsV8Wrapper[T] {
	return &HTMLHyperlinkElementUtilsV8Wrapper[T]{}
}

func (wrapper HTMLHyperlinkElementUtilsV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("href", w.href, w.setHref)
	jsClass.CreatePrototypeMethod("toString", w.href)
	jsClass.CreatePrototypeAttribute("origin", w.origin, nil)
	jsClass.CreatePrototypeAttribute("protocol", w.protocol, w.setProtocol)
	jsClass.CreatePrototypeAttribute("username", w.username, w.setUsername)
	jsClass.CreatePrototypeAttribute("password", w.password, w.setPassword)
	jsClass.CreatePrototypeAttribute("host", w.host, w.setHost)
	jsClass.CreatePrototypeAttribute("hostname", w.hostname, w.setHostname)
	jsClass.CreatePrototypeAttribute("port", w.port, w.setPort)
	jsClass.CreatePrototypeAttribute("pathname", w.pathname, w.setPathname)
	jsClass.CreatePrototypeAttribute("search", w.search, w.setSearch)
	jsClass.CreatePrototypeAttribute("hash", w.hash, w.setHash)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) href(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.href")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) setHref(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setHref")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHref(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) origin(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.origin")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) protocol(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.protocol")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) setProtocol(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setProtocol")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetProtocol(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) username(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.username")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Username()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) setUsername(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setUsername")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetUsername(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) password(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.password")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Password()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) setPassword(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setPassword")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetPassword(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) host(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.host")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) setHost(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setHost")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHost(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) hostname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.hostname")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) setHostname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setHostname")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHostname(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) port(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.port")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) setPort(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setPort")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetPort(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) pathname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.pathname")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) setPathname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setPathname")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetPathname(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) search(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.search")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) setSearch(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setSearch")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetSearch(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) hash(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.hash")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtilsV8Wrapper[T]) setHash(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLHyperlinkElementUtils.setHash")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHash(val)
	return nil, nil
}
