// This file is generated. Do not edit.

package html

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLHyperlinkElementUtils[T any] struct{}

func NewHTMLHyperlinkElementUtils[T any](scriptHost js.ScriptEngine[T]) *HTMLHyperlinkElementUtils[T] {
	return &HTMLHyperlinkElementUtils[T]{}
}

func (wrapper HTMLHyperlinkElementUtils[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLHyperlinkElementUtils[T]) installPrototype(jsClass js.Class[T]) {
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

func (w HTMLHyperlinkElementUtils[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.Constructor - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.Constructor", js.LogAttr("res", res))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLHyperlinkElementUtils[T]) href(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.href - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.href", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) setHref(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setHref - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setHref", js.LogAttr("res", res))
	}()
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHref(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtils[T]) origin(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.origin - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.origin", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) protocol(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.protocol - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.protocol", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) setProtocol(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setProtocol - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setProtocol", js.LogAttr("res", res))
	}()
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetProtocol(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtils[T]) username(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.username - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.username", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Username()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) setUsername(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setUsername - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setUsername", js.LogAttr("res", res))
	}()
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetUsername(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtils[T]) password(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.password - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.password", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Password()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) setPassword(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setPassword - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setPassword", js.LogAttr("res", res))
	}()
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetPassword(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtils[T]) host(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.host - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.host", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) setHost(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setHost - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setHost", js.LogAttr("res", res))
	}()
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHost(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtils[T]) hostname(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.hostname - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.hostname", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) setHostname(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setHostname - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setHostname", js.LogAttr("res", res))
	}()
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHostname(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtils[T]) port(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.port - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.port", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) setPort(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setPort - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setPort", js.LogAttr("res", res))
	}()
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetPort(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtils[T]) pathname(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.pathname - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.pathname", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) setPathname(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setPathname - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setPathname", js.LogAttr("res", res))
	}()
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetPathname(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtils[T]) search(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.search - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.search", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) setSearch(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setSearch - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setSearch", js.LogAttr("res", res))
	}()
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetSearch(val)
	return nil, nil
}

func (w HTMLHyperlinkElementUtils[T]) hash(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.hash - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.hash", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) setHash(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setHash - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setHash", js.LogAttr("res", res))
	}()
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHash(val)
	return nil, nil
}
