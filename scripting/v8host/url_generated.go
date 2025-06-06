// This file is generated. Do not edit.

package v8host

import (
	"errors"
	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	url "github.com/gost-dom/browser/url"
)

func (wrapper urlV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w urlV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("toJSON", w.toJSON)
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
	jsClass.CreatePrototypeAttribute("searchParams", w.searchParams, nil)
	jsClass.CreatePrototypeAttribute("hash", w.hash, w.setHash)
}

func (w urlV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.Constructor")
	url, errArg1 := consumeArgument(cbCtx, "url", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	base, found, errArg := consumeOptionalArg(cbCtx, "base", codec.DecodeString)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceBase(cbCtx, url, base)
	}
	return w.CreateInstance(cbCtx, url)
}

func (w urlV8Wrapper[T]) toJSON(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.toJSON")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result, errCall := instance.ToJSON()
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeString(cbCtx, result)
}

func (w urlV8Wrapper[T]) href(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.href")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Href()
	return codec.EncodeString(cbCtx, result)
}

func (w urlV8Wrapper[T]) setHref(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHref")
	return cbCtx.ReturnWithError(errors.New("URL.setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) origin(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.origin")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Origin()
	return codec.EncodeString(cbCtx, result)
}

func (w urlV8Wrapper[T]) protocol(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.protocol")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Protocol()
	return codec.EncodeString(cbCtx, result)
}

func (w urlV8Wrapper[T]) setProtocol(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setProtocol")
	return cbCtx.ReturnWithError(errors.New("URL.setProtocol: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) username(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.username")
	return cbCtx.ReturnWithError(errors.New("URL.username: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) setUsername(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setUsername")
	return cbCtx.ReturnWithError(errors.New("URL.setUsername: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) password(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.password")
	return cbCtx.ReturnWithError(errors.New("URL.password: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) setPassword(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setPassword")
	return cbCtx.ReturnWithError(errors.New("URL.setPassword: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) host(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.host")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Host()
	return codec.EncodeString(cbCtx, result)
}

func (w urlV8Wrapper[T]) setHost(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHost")
	return cbCtx.ReturnWithError(errors.New("URL.setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) hostname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.hostname")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Hostname()
	return codec.EncodeString(cbCtx, result)
}

func (w urlV8Wrapper[T]) setHostname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHostname")
	return cbCtx.ReturnWithError(errors.New("URL.setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) port(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.port")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Port()
	return codec.EncodeString(cbCtx, result)
}

func (w urlV8Wrapper[T]) setPort(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setPort")
	return cbCtx.ReturnWithError(errors.New("URL.setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) pathname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.pathname")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Pathname()
	return codec.EncodeString(cbCtx, result)
}

func (w urlV8Wrapper[T]) setPathname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setPathname")
	return cbCtx.ReturnWithError(errors.New("URL.setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) search(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.search")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Search()
	return codec.EncodeString(cbCtx, result)
}

func (w urlV8Wrapper[T]) setSearch(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setSearch")
	return cbCtx.ReturnWithError(errors.New("URL.setSearch: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) searchParams(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.searchParams")
	return cbCtx.ReturnWithError(errors.New("URL.searchParams: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w urlV8Wrapper[T]) hash(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.hash")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Hash()
	return codec.EncodeString(cbCtx, result)
}

func (w urlV8Wrapper[T]) setHash(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHash")
	return cbCtx.ReturnWithError(errors.New("URL.setHash: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

type urlSearchParamsV8Wrapper[T any] struct {
	handleReffedObject[urlinterfaces.URLSearchParams, T]
}

func newURLSearchParamsV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *urlSearchParamsV8Wrapper[T] {
	return &urlSearchParamsV8Wrapper[T]{newHandleReffedObject[urlinterfaces.URLSearchParams, T](scriptHost)}
}

func (wrapper urlSearchParamsV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w urlSearchParamsV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("append", w.append)
	jsClass.CreatePrototypeMethod("delete", w.delete)
	jsClass.CreatePrototypeMethod("get", w.get)
	jsClass.CreatePrototypeMethod("getAll", w.getAll)
	jsClass.CreatePrototypeMethod("has", w.has)
	jsClass.CreatePrototypeMethod("set", w.set)
	jsClass.CreatePrototypeMethod("sort", w.sort)
	jsClass.CreatePrototypeMethod("toString", w.toString)
	jsClass.CreatePrototypeAttribute("size", w.size, nil)
}

func (w urlSearchParamsV8Wrapper[T]) append(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.append")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := consumeArgument(cbCtx, "value", nil, codec.DecodeString)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Append(name, value)
	return nil, nil
}

func (w urlSearchParamsV8Wrapper[T]) delete(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.delete")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	value, found, errArg := consumeOptionalArg(cbCtx, "value", codec.DecodeString)
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

func (w urlSearchParamsV8Wrapper[T]) get(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.get")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.Get(name)
	return codec.EncodeNillableString(cbCtx, result, hasValue)
}

func (w urlSearchParamsV8Wrapper[T]) getAll(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.getAll")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetAll(name)
	return w.toSequenceString_(cbCtx, result)
}

func (w urlSearchParamsV8Wrapper[T]) has(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.has")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	value, found, errArg := consumeOptionalArg(cbCtx, "value", codec.DecodeString)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		result := instance.HasValue(name, value)
		return codec.EncodeBoolean(cbCtx, result)
	}
	result := instance.Has(name)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w urlSearchParamsV8Wrapper[T]) set(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.set")
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	name, errArg1 := consumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := consumeArgument(cbCtx, "value", nil, codec.DecodeString)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Set(name, value)
	return nil, nil
}

func (w urlSearchParamsV8Wrapper[T]) sort(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.sort")
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Sort()
	return nil, nil
}

func (w urlSearchParamsV8Wrapper[T]) toString(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.toString")
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.String()
	return codec.EncodeString(cbCtx, result)
}

func (w urlSearchParamsV8Wrapper[T]) size(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URLSearchParams.size")
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Size()
	return codec.EncodeInt(cbCtx, result)
}
