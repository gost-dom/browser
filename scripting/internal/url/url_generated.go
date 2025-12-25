// This file is generated. Do not edit.

package url

import (
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	url "github.com/gost-dom/browser/url"
)

func (wrapper URL[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w URL[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("toJSON", w.toJSON)
	jsClass.CreatePrototypeAttribute("href", w.href, w.setHref)
	jsClass.CreateOperation("toString", w.href)
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

func (w URL[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	url, errArg1 := js.ConsumeArgument(cbCtx, "url", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	base, found, errArg := js.ConsumeOptionalArg(cbCtx, "base", codec.DecodeString)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceBase(cbCtx, url, base)
	}
	return w.CreateInstance(cbCtx, url)
}

func (w URL[T]) toJSON(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result, errCall := instance.ToJSON()
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeString(cbCtx, result)
}

func (w URL[T]) href(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return codec.EncodeString(cbCtx, result)
}

func (w URL[T]) setHref(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URL[T]) origin(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return codec.EncodeString(cbCtx, result)
}

func (w URL[T]) protocol(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return codec.EncodeString(cbCtx, result)
}

func (w URL[T]) setProtocol(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.setProtocol: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URL[T]) username(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.username: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URL[T]) setUsername(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.setUsername: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URL[T]) password(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.password: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URL[T]) setPassword(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.setPassword: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URL[T]) host(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return codec.EncodeString(cbCtx, result)
}

func (w URL[T]) setHost(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URL[T]) hostname(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return codec.EncodeString(cbCtx, result)
}

func (w URL[T]) setHostname(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URL[T]) port(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return codec.EncodeString(cbCtx, result)
}

func (w URL[T]) setPort(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URL[T]) pathname(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return codec.EncodeString(cbCtx, result)
}

func (w URL[T]) setPathname(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URL[T]) search(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return codec.EncodeString(cbCtx, result)
}

func (w URL[T]) setSearch(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[*url.URL](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetSearch(val)
	return nil, nil
}

func (w URL[T]) searchParams(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.searchParams: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URL[T]) hash(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return codec.EncodeString(cbCtx, result)
}

func (w URL[T]) setHash(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.setHash: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
