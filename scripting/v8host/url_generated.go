// This file is generated. Do not edit.

package v8host

import (
	"errors"
	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
	url "github.com/gost-dom/browser/url"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("URL", "", createURLPrototype)
}

func createURLPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newURLV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor)

	return constructor
}
func (w urlV8Wrapper) installPrototype(ft *v8.FunctionTemplate) {
	jsClass := newV8Class(w.scriptHost, ft)
	jsClass.CreatePrototypeMethod("toJSON", w.toJSON)
	prototypeTmpl := ft.PrototypeTemplate()
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
	prototypeTmpl.SetAccessorProperty("searchParams",
		wrapV8Callback(w.scriptHost, w.searchParams),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("hash",
		wrapV8Callback(w.scriptHost, w.hash),
		wrapV8Callback(w.scriptHost, w.setHash),
		v8.None)
}

func (w urlV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.Constructor")
	url, errArg1 := consumeArgument(cbCtx, "url", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	base, found, errArg := consumeOptionalArg(cbCtx, "base", w.decodeString)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceBase(cbCtx, url, base)
	}
	return w.CreateInstance(cbCtx, url)
}

func (w urlV8Wrapper) toJSON(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.toJSON")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result, errCall := instance.ToJSON()
	if errCall != nil {
		return nil, errCall
	}
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) href(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.href")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Href()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setHref(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHref")
	return cbCtx.ReturnWithError(errors.New("URL.setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) origin(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.origin")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Origin()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) protocol(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.protocol")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Protocol()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setProtocol(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setProtocol")
	return cbCtx.ReturnWithError(errors.New("URL.setProtocol: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) username(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.username")
	return cbCtx.ReturnWithError(errors.New("URL.username: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) setUsername(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setUsername")
	return cbCtx.ReturnWithError(errors.New("URL.setUsername: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) password(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.password")
	return cbCtx.ReturnWithError(errors.New("URL.password: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) setPassword(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setPassword")
	return cbCtx.ReturnWithError(errors.New("URL.setPassword: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) host(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.host")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Host()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setHost(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHost")
	return cbCtx.ReturnWithError(errors.New("URL.setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) hostname(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.hostname")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Hostname()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setHostname(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHostname")
	return cbCtx.ReturnWithError(errors.New("URL.setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) port(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.port")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Port()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setPort(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setPort")
	return cbCtx.ReturnWithError(errors.New("URL.setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) pathname(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.pathname")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Pathname()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setPathname(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setPathname")
	return cbCtx.ReturnWithError(errors.New("URL.setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) search(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.search")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Search()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setSearch(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setSearch")
	return cbCtx.ReturnWithError(errors.New("URL.setSearch: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) searchParams(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.searchParams")
	return cbCtx.ReturnWithError(errors.New("URL.searchParams: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) hash(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.hash")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Hash()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setHash(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHash")
	return cbCtx.ReturnWithError(errors.New("URL.setHash: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func init() {
	registerJSClass("URLSearchParams", "", createURLSearchParamsPrototype)
}

type urlSearchParamsV8Wrapper struct {
	handleReffedObject[urlinterfaces.URLSearchParams, jsTypeParam]
}

func newURLSearchParamsV8Wrapper(scriptHost *V8ScriptHost) *urlSearchParamsV8Wrapper {
	return &urlSearchParamsV8Wrapper{newHandleReffedObject[urlinterfaces.URLSearchParams](scriptHost)}
}

func createURLSearchParamsPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newURLSearchParamsV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor)

	wrapper.CustomInitialiser(constructor)
	return constructor
}
func (w urlSearchParamsV8Wrapper) installPrototype(ft *v8.FunctionTemplate) {
	jsClass := newV8Class(w.scriptHost, ft)
	jsClass.CreatePrototypeMethod("append", w.append)
	jsClass.CreatePrototypeMethod("delete", w.delete)
	jsClass.CreatePrototypeMethod("get", w.get)
	jsClass.CreatePrototypeMethod("getAll", w.getAll)
	jsClass.CreatePrototypeMethod("has", w.has)
	jsClass.CreatePrototypeMethod("set", w.set)
	jsClass.CreatePrototypeMethod("sort", w.sort)
	jsClass.CreatePrototypeMethod("toString", w.toString)
	prototypeTmpl := ft.PrototypeTemplate()
	prototypeTmpl.SetAccessorProperty("size",
		wrapV8Callback(w.scriptHost, w.size),
		nil,
		v8.None)
}

func (w urlSearchParamsV8Wrapper) append(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.append")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, errArg2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Append(name, value)
	return nil, nil
}

func (w urlSearchParamsV8Wrapper) delete(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.delete")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	value, found, errArg := consumeOptionalArg(cbCtx, "value", w.decodeString)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		instance.DeleteValue(name, value)
		return nil, nil
	}
	instance.Delete(name)
	return nil, nil
}

func (w urlSearchParamsV8Wrapper) get(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.get")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.Get(name)
	return w.toNillableString_(cbCtx, result, hasValue)
}

func (w urlSearchParamsV8Wrapper) getAll(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.getAll")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetAll(name)
	return w.toSequenceString_(cbCtx, result)
}

func (w urlSearchParamsV8Wrapper) has(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.has")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	value, found, errArg := consumeOptionalArg(cbCtx, "value", w.decodeString)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		result := instance.HasValue(name, value)
		return w.toBoolean(cbCtx, result)
	}
	result := instance.Has(name)
	return w.toBoolean(cbCtx, result)
}

func (w urlSearchParamsV8Wrapper) set(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.set")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, errArg2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Set(name, value)
	return nil, nil
}

func (w urlSearchParamsV8Wrapper) sort(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.sort")
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Sort()
	return nil, nil
}

func (w urlSearchParamsV8Wrapper) toString(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.toString")
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.String()
	return w.toString_(cbCtx, result)
}

func (w urlSearchParamsV8Wrapper) size(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.size")
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Size()
	return w.toUnsignedLong(cbCtx, result)
}
