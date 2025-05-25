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

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w urlV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("toJSON", wrapV8Callback(w.scriptHost, w.toJSON))

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

func (w urlV8Wrapper) Constructor(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.Constructor")
	url, err1 := consumeArgument(cbCtx, "url", nil, w.decodeString)
	base, err2 := consumeArgument(cbCtx, "base", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		return w.CreateInstanceBase(cbCtx, url, base)
	}
	if cbCtx.noOfReadArguments >= 1 {
		if err1 != nil {
			return cbCtx.ReturnWithError(err1)
		}
		return w.CreateInstance(cbCtx, url)
	}
	return cbCtx.ReturnWithError(errors.New("URL.constructor: Missing arguments"))
}

func (w urlV8Wrapper) toJSON(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.toJSON")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result, callErr := instance.ToJSON()
	if callErr != nil {
		return cbCtx.ReturnWithError(callErr)
	} else {
		return w.toString_(cbCtx, result)
	}
}

func (w urlV8Wrapper) href(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.href")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Href()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setHref(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.setHref")
	return cbCtx.ReturnWithError(errors.New("URL.setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) origin(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.origin")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Origin()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) protocol(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.protocol")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Protocol()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setProtocol(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.setProtocol")
	return cbCtx.ReturnWithError(errors.New("URL.setProtocol: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) username(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.username")
	return cbCtx.ReturnWithError(errors.New("URL.username: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) setUsername(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.setUsername")
	return cbCtx.ReturnWithError(errors.New("URL.setUsername: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) password(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.password")
	return cbCtx.ReturnWithError(errors.New("URL.password: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) setPassword(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.setPassword")
	return cbCtx.ReturnWithError(errors.New("URL.setPassword: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) host(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.host")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Host()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setHost(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.setHost")
	return cbCtx.ReturnWithError(errors.New("URL.setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) hostname(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.hostname")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Hostname()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setHostname(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.setHostname")
	return cbCtx.ReturnWithError(errors.New("URL.setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) port(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.port")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Port()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setPort(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.setPort")
	return cbCtx.ReturnWithError(errors.New("URL.setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) pathname(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.pathname")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Pathname()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setPathname(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.setPathname")
	return cbCtx.ReturnWithError(errors.New("URL.setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) search(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.search")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Search()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setSearch(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.setSearch")
	return cbCtx.ReturnWithError(errors.New("URL.setSearch: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) searchParams(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.searchParams")
	return cbCtx.ReturnWithError(errors.New("URL.searchParams: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper) hash(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.hash")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Hash()
	return w.toString_(cbCtx, result)
}

func (w urlV8Wrapper) setHash(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URL.setHash")
	return cbCtx.ReturnWithError(errors.New("URL.setHash: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func init() {
	registerJSClass("URLSearchParams", "", createURLSearchParamsPrototype)
}

type urlSearchParamsV8Wrapper struct {
	handleReffedObject[urlinterfaces.URLSearchParams]
}

func newURLSearchParamsV8Wrapper(scriptHost *V8ScriptHost) *urlSearchParamsV8Wrapper {
	return &urlSearchParamsV8Wrapper{newHandleReffedObject[urlinterfaces.URLSearchParams](scriptHost)}
}

func createURLSearchParamsPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newURLSearchParamsV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	wrapper.CustomInitialiser(constructor)
	return constructor
}
func (w urlSearchParamsV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("append", wrapV8Callback(w.scriptHost, w.append))
	prototypeTmpl.Set("delete", wrapV8Callback(w.scriptHost, w.delete))
	prototypeTmpl.Set("get", wrapV8Callback(w.scriptHost, w.get))
	prototypeTmpl.Set("getAll", wrapV8Callback(w.scriptHost, w.getAll))
	prototypeTmpl.Set("has", wrapV8Callback(w.scriptHost, w.has))
	prototypeTmpl.Set("set", wrapV8Callback(w.scriptHost, w.set))
	prototypeTmpl.Set("sort", wrapV8Callback(w.scriptHost, w.sort))
	prototypeTmpl.Set("toString", wrapV8Callback(w.scriptHost, w.toString))

	prototypeTmpl.SetAccessorProperty("size",
		wrapV8Callback(w.scriptHost, w.size),
		nil,
		v8.None)
}

func (w urlSearchParamsV8Wrapper) append(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URLSearchParams.append")
	instance, err0 := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, err2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		instance.Append(name, value)
		return cbCtx.ReturnWithValue(nil)
	}
	return cbCtx.ReturnWithError(errors.New("URLSearchParams.append: Missing arguments"))
}

func (w urlSearchParamsV8Wrapper) delete(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URLSearchParams.delete")
	instance, err0 := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, err2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		instance.DeleteValue(name, value)
		return cbCtx.ReturnWithValue(nil)
	}
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		instance.Delete(name)
		return cbCtx.ReturnWithValue(nil)
	}
	return cbCtx.ReturnWithError(errors.New("URLSearchParams.delete: Missing arguments"))
}

func (w urlSearchParamsV8Wrapper) get(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URLSearchParams.get")
	instance, err0 := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result, hasValue := instance.Get(name)
		return w.toNillableString_(cbCtx, result, hasValue)
	}
	return cbCtx.ReturnWithError(errors.New("URLSearchParams.get: Missing arguments"))
}

func (w urlSearchParamsV8Wrapper) getAll(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URLSearchParams.getAll")
	instance, err0 := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result := instance.GetAll(name)
		return w.toSequenceString_(cbCtx, result)
	}
	return cbCtx.ReturnWithError(errors.New("URLSearchParams.getAll: Missing arguments"))
}

func (w urlSearchParamsV8Wrapper) has(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URLSearchParams.has")
	instance, err0 := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, err2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result := instance.HasValue(name, value)
		return w.toBoolean(cbCtx, result)
	}
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result := instance.Has(name)
		return w.toBoolean(cbCtx, result)
	}
	return cbCtx.ReturnWithError(errors.New("URLSearchParams.has: Missing arguments"))
}

func (w urlSearchParamsV8Wrapper) set(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URLSearchParams.set")
	instance, err0 := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, err2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		instance.Set(name, value)
		return cbCtx.ReturnWithValue(nil)
	}
	return cbCtx.ReturnWithError(errors.New("URLSearchParams.set: Missing arguments"))
}

func (w urlSearchParamsV8Wrapper) sort(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URLSearchParams.sort")
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Sort()
	return cbCtx.ReturnWithValue(nil)
}

func (w urlSearchParamsV8Wrapper) toString(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URLSearchParams.toString")
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.String()
	return w.toString_(cbCtx, result)
}

func (w urlSearchParamsV8Wrapper) size(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: URLSearchParams.size")
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Size()
	return w.toUnsignedLong(cbCtx, result)
}
