// This file is generated. Do not edit.

package v8host

import (
	"errors"
	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	log "github.com/gost-dom/browser/internal/log"
	abstraction "github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	url "github.com/gost-dom/browser/url"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("URL", "", createURLPrototype)
}

func createURLPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newURLV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

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

func (w urlV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.Constructor")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	url, err1 := consumeArgument(cbCtx, "url", nil, w.decodeString)
	base, err2 := consumeArgument(cbCtx, "base", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return nil, err
		}
		return w.CreateInstanceBase(cbCtx.Context(), info.This(), url, base)
	}
	if cbCtx.noOfReadArguments >= 1 {
		if err1 != nil {
			return nil, err1
		}
		return w.CreateInstance(cbCtx.Context(), info.This(), url)
	}
	return nil, errors.New("URL.constructor: Missing arguments")
}

func (w urlV8Wrapper) toJSON(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.toJSON")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result, callErr := instance.ToJSON()
	if callErr != nil {
		return nil, callErr
	} else {
		return w.toString_(cbCtx.Context(), result)
	}
}

func (w urlV8Wrapper) href(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.href")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return w.toString_(cbCtx.Context(), result)
}

func (w urlV8Wrapper) setHref(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setHref")
	return nil, errors.New("URL.setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) origin(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.origin")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return w.toString_(cbCtx.Context(), result)
}

func (w urlV8Wrapper) protocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.protocol")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return w.toString_(cbCtx.Context(), result)
}

func (w urlV8Wrapper) setProtocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setProtocol")
	return nil, errors.New("URL.setProtocol: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) username(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.username")
	return nil, errors.New("URL.username: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) setUsername(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setUsername")
	return nil, errors.New("URL.setUsername: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) password(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.password")
	return nil, errors.New("URL.password: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) setPassword(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setPassword")
	return nil, errors.New("URL.setPassword: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) host(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.host")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return w.toString_(cbCtx.Context(), result)
}

func (w urlV8Wrapper) setHost(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setHost")
	return nil, errors.New("URL.setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) hostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.hostname")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return w.toString_(cbCtx.Context(), result)
}

func (w urlV8Wrapper) setHostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setHostname")
	return nil, errors.New("URL.setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) port(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.port")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return w.toString_(cbCtx.Context(), result)
}

func (w urlV8Wrapper) setPort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setPort")
	return nil, errors.New("URL.setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) pathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.pathname")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return w.toString_(cbCtx.Context(), result)
}

func (w urlV8Wrapper) setPathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setPathname")
	return nil, errors.New("URL.setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) search(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.search")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return w.toString_(cbCtx.Context(), result)
}

func (w urlV8Wrapper) setSearch(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setSearch")
	return nil, errors.New("URL.setSearch: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) searchParams(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.searchParams")
	return nil, errors.New("URL.searchParams: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) hash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.hash")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return w.toString_(cbCtx.Context(), result)
}

func (w urlV8Wrapper) setHash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setHash")
	return nil, errors.New("URL.setHash: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
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
	iso := scriptHost.iso
	wrapper := newURLSearchParamsV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

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

func (w urlSearchParamsV8Wrapper) append(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.append")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, err2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		instance.Append(name, value)
		return nil, nil
	}
	return nil, errors.New("URLSearchParams.append: Missing arguments")
}

func (w urlSearchParamsV8Wrapper) delete(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.delete")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, err2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		instance.DeleteValue(name, value)
		return nil, nil
	}
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.Delete(name)
		return nil, nil
	}
	return nil, errors.New("URLSearchParams.delete: Missing arguments")
}

func (w urlSearchParamsV8Wrapper) getAll(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.getAll")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.GetAll(name)
		return w.toSequenceString_(cbCtx.Context(), result)
	}
	return nil, errors.New("URLSearchParams.getAll: Missing arguments")
}

func (w urlSearchParamsV8Wrapper) has(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.has")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, err2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		result := instance.HasValue(name, value)
		return w.toBoolean(cbCtx.Context(), result)
	}
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Has(name)
		return w.toBoolean(cbCtx.Context(), result)
	}
	return nil, errors.New("URLSearchParams.has: Missing arguments")
}

func (w urlSearchParamsV8Wrapper) set(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.set")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	name, err1 := consumeArgument(cbCtx, "name", nil, w.decodeString)
	value, err2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		instance.Set(name, value)
		return nil, nil
	}
	return nil, errors.New("URLSearchParams.set: Missing arguments")
}

func (w urlSearchParamsV8Wrapper) sort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.sort")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Sort()
	return nil, nil
}

func (w urlSearchParamsV8Wrapper) toString(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.toString")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.String()
	return w.toString_(cbCtx.Context(), result)
}

func (w urlSearchParamsV8Wrapper) size(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.size")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Size()
	return w.toUnsignedLong(cbCtx.Context(), result)
}
