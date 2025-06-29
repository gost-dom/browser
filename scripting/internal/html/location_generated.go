// This file is generated. Do not edit.

package html

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Location[T any] struct{}

func NewLocation[T any](scriptHost js.ScriptEngine[T]) *Location[T] {
	return &Location[T]{}
}

func (wrapper Location[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Location[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("assign", w.assign)
	jsClass.CreatePrototypeMethod("replace", w.replace)
	jsClass.CreatePrototypeMethod("reload", w.reload)
	jsClass.CreatePrototypeAttribute("href", w.href, w.setHref)
	jsClass.CreatePrototypeMethod("toString", w.href)
	jsClass.CreatePrototypeAttribute("origin", w.origin, nil)
	jsClass.CreatePrototypeAttribute("protocol", w.protocol, w.setProtocol)
	jsClass.CreatePrototypeAttribute("host", w.host, w.setHost)
	jsClass.CreatePrototypeAttribute("hostname", w.hostname, w.setHostname)
	jsClass.CreatePrototypeAttribute("port", w.port, w.setPort)
	jsClass.CreatePrototypeAttribute("pathname", w.pathname, w.setPathname)
	jsClass.CreatePrototypeAttribute("search", w.search, w.setSearch)
	jsClass.CreatePrototypeAttribute("hash", w.hash, w.setHash)
	jsClass.CreatePrototypeAttribute("ancestorOrigins", w.ancestorOrigins, nil)
}

func (w Location[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Location[T]) assign(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.assign")
	instance, errInst := js.As[html.Location](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	url, errArg1 := js.ConsumeArgument(cbCtx, "url", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Assign(url)
	return nil, errCall
}

func (w Location[T]) replace(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.replace")
	instance, errInst := js.As[html.Location](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	url, errArg1 := js.ConsumeArgument(cbCtx, "url", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Replace(url)
	return nil, errCall
}

func (w Location[T]) reload(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.reload")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.Reload()
	return nil, errCall
}

func (w Location[T]) href(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.href")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return codec.EncodeString(cbCtx, result)
}

func (w Location[T]) setHref(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.setHref")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHref(val)
	return nil, nil
}

func (w Location[T]) origin(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.origin")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return codec.EncodeString(cbCtx, result)
}

func (w Location[T]) protocol(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.protocol")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return codec.EncodeString(cbCtx, result)
}

func (w Location[T]) setProtocol(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.setProtocol")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetProtocol(val)
	return nil, nil
}

func (w Location[T]) host(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.host")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return codec.EncodeString(cbCtx, result)
}

func (w Location[T]) setHost(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.setHost")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHost(val)
	return nil, nil
}

func (w Location[T]) hostname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.hostname")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return codec.EncodeString(cbCtx, result)
}

func (w Location[T]) setHostname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.setHostname")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHostname(val)
	return nil, nil
}

func (w Location[T]) port(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.port")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return codec.EncodeString(cbCtx, result)
}

func (w Location[T]) setPort(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.setPort")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetPort(val)
	return nil, nil
}

func (w Location[T]) pathname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.pathname")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return codec.EncodeString(cbCtx, result)
}

func (w Location[T]) setPathname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.setPathname")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetPathname(val)
	return nil, nil
}

func (w Location[T]) search(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.search")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return codec.EncodeString(cbCtx, result)
}

func (w Location[T]) setSearch(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.setSearch")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetSearch(val)
	return nil, nil
}

func (w Location[T]) hash(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.hash")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return codec.EncodeString(cbCtx, result)
}

func (w Location[T]) setHash(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.setHash")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHash(val)
	return nil, nil
}

func (w Location[T]) ancestorOrigins(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Location.ancestorOrigins")
	return codec.EncodeCallbackErrorf(cbCtx, "Location.ancestorOrigins: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
