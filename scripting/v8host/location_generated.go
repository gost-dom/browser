// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("Location", "", newLocationV8Wrapper)
}

type locationV8Wrapper[T any] struct {
	handleReffedObject[html.Location, T]
}

func newLocationV8Wrapper(scriptHost jsScriptEngine) *locationV8Wrapper[jsTypeParam] {
	return &locationV8Wrapper[jsTypeParam]{newHandleReffedObject[html.Location, jsTypeParam](scriptHost)}
}

func (wrapper locationV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w locationV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
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

func (w locationV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w locationV8Wrapper[T]) assign(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.assign")
	instance, errInst := js.As[html.Location](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	url, errArg1 := consumeArgument(cbCtx, "url", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Assign(url)
	return nil, errCall
}

func (w locationV8Wrapper[T]) replace(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.replace")
	instance, errInst := js.As[html.Location](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	url, errArg1 := consumeArgument(cbCtx, "url", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Replace(url)
	return nil, errCall
}

func (w locationV8Wrapper[T]) reload(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.reload")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	errCall := instance.Reload()
	return nil, errCall
}

func (w locationV8Wrapper[T]) href(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.href")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Href()
	return w.toString_(cbCtx, result)
}

func (w locationV8Wrapper[T]) setHref(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.setHref")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetHref(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w locationV8Wrapper[T]) origin(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.origin")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Origin()
	return w.toString_(cbCtx, result)
}

func (w locationV8Wrapper[T]) protocol(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.protocol")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Protocol()
	return w.toString_(cbCtx, result)
}

func (w locationV8Wrapper[T]) setProtocol(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.setProtocol")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetProtocol(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w locationV8Wrapper[T]) host(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.host")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Host()
	return w.toString_(cbCtx, result)
}

func (w locationV8Wrapper[T]) setHost(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.setHost")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetHost(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w locationV8Wrapper[T]) hostname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.hostname")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Hostname()
	return w.toString_(cbCtx, result)
}

func (w locationV8Wrapper[T]) setHostname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.setHostname")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetHostname(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w locationV8Wrapper[T]) port(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.port")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Port()
	return w.toString_(cbCtx, result)
}

func (w locationV8Wrapper[T]) setPort(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.setPort")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetPort(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w locationV8Wrapper[T]) pathname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.pathname")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Pathname()
	return w.toString_(cbCtx, result)
}

func (w locationV8Wrapper[T]) setPathname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.setPathname")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetPathname(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w locationV8Wrapper[T]) search(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.search")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Search()
	return w.toString_(cbCtx, result)
}

func (w locationV8Wrapper[T]) setSearch(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.setSearch")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetSearch(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w locationV8Wrapper[T]) hash(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.hash")
	instance, err := js.As[html.Location](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Hash()
	return w.toString_(cbCtx, result)
}

func (w locationV8Wrapper[T]) setHash(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.setHash")
	instance, err0 := js.As[html.Location](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetHash(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w locationV8Wrapper[T]) ancestorOrigins(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Location.ancestorOrigins")
	return cbCtx.ReturnWithError(errors.New("Location.ancestorOrigins: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
