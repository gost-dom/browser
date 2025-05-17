// This file is generated. Do not edit.

package v8host

import (
	"errors"
	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	log "github.com/gost-dom/browser/internal/log"
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
	iso := w.scriptHost.iso
	prototypeTmpl.Set("toJSON", v8.NewFunctionTemplateWithError(iso, w.toJSON))

	prototypeTmpl.SetAccessorProperty("href",
		v8.NewFunctionTemplateWithError(iso, w.href),
		v8.NewFunctionTemplateWithError(iso, w.setHref),
		v8.None)
	prototypeTmpl.SetAccessorProperty("origin",
		v8.NewFunctionTemplateWithError(iso, w.origin),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("protocol",
		v8.NewFunctionTemplateWithError(iso, w.protocol),
		v8.NewFunctionTemplateWithError(iso, w.setProtocol),
		v8.None)
	prototypeTmpl.SetAccessorProperty("username",
		v8.NewFunctionTemplateWithError(iso, w.username),
		v8.NewFunctionTemplateWithError(iso, w.setUsername),
		v8.None)
	prototypeTmpl.SetAccessorProperty("password",
		v8.NewFunctionTemplateWithError(iso, w.password),
		v8.NewFunctionTemplateWithError(iso, w.setPassword),
		v8.None)
	prototypeTmpl.SetAccessorProperty("host",
		v8.NewFunctionTemplateWithError(iso, w.host),
		v8.NewFunctionTemplateWithError(iso, w.setHost),
		v8.None)
	prototypeTmpl.SetAccessorProperty("hostname",
		v8.NewFunctionTemplateWithError(iso, w.hostname),
		v8.NewFunctionTemplateWithError(iso, w.setHostname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("port",
		v8.NewFunctionTemplateWithError(iso, w.port),
		v8.NewFunctionTemplateWithError(iso, w.setPort),
		v8.None)
	prototypeTmpl.SetAccessorProperty("pathname",
		v8.NewFunctionTemplateWithError(iso, w.pathname),
		v8.NewFunctionTemplateWithError(iso, w.setPathname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("search",
		v8.NewFunctionTemplateWithError(iso, w.search),
		v8.NewFunctionTemplateWithError(iso, w.setSearch),
		v8.None)
	prototypeTmpl.SetAccessorProperty("searchParams",
		v8.NewFunctionTemplateWithError(iso, w.searchParams),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("hash",
		v8.NewFunctionTemplateWithError(iso, w.hash),
		v8.NewFunctionTemplateWithError(iso, w.setHash),
		v8.None)
}

func (w urlV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(w.scriptHost, info)
	url, err1 := tryParseArg(args, 0, w.decodeUSVString)
	base, err2 := tryParseArg(args, 1, w.decodeUSVString)
	ctx := w.mustGetContext(info)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return nil, err
		}
		return w.CreateInstanceBase(ctx, info.This(), url, base)
	}
	if args.noOfReadArguments >= 1 {
		if err1 != nil {
			return nil, err1
		}
		return w.CreateInstance(ctx, info.This(), url)
	}
	return nil, errors.New("URL.constructor: Missing arguments")
}

func (w urlV8Wrapper) toJSON(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.toJSON")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result, callErr := instance.ToJSON()
	if callErr != nil {
		return nil, callErr
	} else {
		return w.toUSVString(ctx, result)
	}
}

func (w urlV8Wrapper) href(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.href")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return w.toUSVString(ctx, result)
}

func (w urlV8Wrapper) setHref(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setHref")
	return nil, errors.New("URL.setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) origin(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.origin")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return w.toUSVString(ctx, result)
}

func (w urlV8Wrapper) protocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.protocol")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return w.toUSVString(ctx, result)
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
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return w.toUSVString(ctx, result)
}

func (w urlV8Wrapper) setHost(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setHost")
	return nil, errors.New("URL.setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) hostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.hostname")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return w.toUSVString(ctx, result)
}

func (w urlV8Wrapper) setHostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setHostname")
	return nil, errors.New("URL.setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) port(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.port")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return w.toUSVString(ctx, result)
}

func (w urlV8Wrapper) setPort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setPort")
	return nil, errors.New("URL.setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) pathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.pathname")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return w.toUSVString(ctx, result)
}

func (w urlV8Wrapper) setPathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.setPathname")
	return nil, errors.New("URL.setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w urlV8Wrapper) search(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URL.search")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return w.toUSVString(ctx, result)
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
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return w.toUSVString(ctx, result)
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

	return constructor
}
func (w urlSearchParamsV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("append", v8.NewFunctionTemplateWithError(iso, w.append))
	prototypeTmpl.Set("delete", v8.NewFunctionTemplateWithError(iso, w.delete))
	prototypeTmpl.Set("get", v8.NewFunctionTemplateWithError(iso, w.get))
	prototypeTmpl.Set("getAll", v8.NewFunctionTemplateWithError(iso, w.getAll))
	prototypeTmpl.Set("has", v8.NewFunctionTemplateWithError(iso, w.has))
	prototypeTmpl.Set("set", v8.NewFunctionTemplateWithError(iso, w.set))
	prototypeTmpl.Set("sort", v8.NewFunctionTemplateWithError(iso, w.sort))

	prototypeTmpl.SetAccessorProperty("size",
		v8.NewFunctionTemplateWithError(iso, w.size),
		nil,
		v8.None)
}

func (w urlSearchParamsV8Wrapper) append(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.append")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	name, err1 := tryParseArg(args, 0, w.decodeUSVString)
	value, err2 := tryParseArg(args, 1, w.decodeUSVString)
	if args.noOfReadArguments >= 2 {
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
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	name, err1 := tryParseArg(args, 0, w.decodeUSVString)
	value, err2 := tryParseArg(args, 1, w.decodeUSVString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		instance.DeleteValue(name, value)
		return nil, nil
	}
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.Delete(name)
		return nil, nil
	}
	return nil, errors.New("URLSearchParams.delete: Missing arguments")
}

func (w urlSearchParamsV8Wrapper) get(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.get")
	ctx := w.mustGetContext(info)
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	name, err1 := tryParseArg(args, 0, w.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Get(name)
		return w.toNullableUSVString(ctx, result)
	}
	return nil, errors.New("URLSearchParams.get: Missing arguments")
}

func (w urlSearchParamsV8Wrapper) getAll(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.getAll")
	ctx := w.mustGetContext(info)
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	name, err1 := tryParseArg(args, 0, w.decodeUSVString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.GetAll(name)
		return w.toSequenceUSVString(ctx, result)
	}
	return nil, errors.New("URLSearchParams.getAll: Missing arguments")
}

func (w urlSearchParamsV8Wrapper) has(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.has")
	ctx := w.mustGetContext(info)
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	name, err1 := tryParseArg(args, 0, w.decodeUSVString)
	value, err2 := tryParseArg(args, 1, w.decodeUSVString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		result := instance.HasValue(name, value)
		return w.toBoolean(ctx, result)
	}
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Has(name)
		return w.toBoolean(ctx, result)
	}
	return nil, errors.New("URLSearchParams.has: Missing arguments")
}

func (w urlSearchParamsV8Wrapper) set(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.set")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	name, err1 := tryParseArg(args, 0, w.decodeUSVString)
	value, err2 := tryParseArg(args, 1, w.decodeUSVString)
	if args.noOfReadArguments >= 2 {
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
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	instance.Sort()
	return nil, nil
}

func (w urlSearchParamsV8Wrapper) size(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: URLSearchParams.size")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Size()
	return w.toUnsignedLong(ctx, result)
}
