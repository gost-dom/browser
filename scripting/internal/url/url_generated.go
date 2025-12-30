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
	jsClass.CreateOperation("toJSON", URL_toJSON)
	jsClass.CreateAttribute("href", URL_href, URL_setHref)
	jsClass.CreateOperation("toString", URL_href)
	jsClass.CreateAttribute("origin", URL_origin, nil)
	jsClass.CreateAttribute("protocol", URL_protocol, URL_setProtocol)
	jsClass.CreateAttribute("username", URL_username, URL_setUsername)
	jsClass.CreateAttribute("password", URL_password, URL_setPassword)
	jsClass.CreateAttribute("host", URL_host, URL_setHost)
	jsClass.CreateAttribute("hostname", URL_hostname, URL_setHostname)
	jsClass.CreateAttribute("port", URL_port, URL_setPort)
	jsClass.CreateAttribute("pathname", URL_pathname, URL_setPathname)
	jsClass.CreateAttribute("search", URL_search, URL_setSearch)
	jsClass.CreateAttribute("searchParams", URL_searchParams, nil)
	jsClass.CreateAttribute("hash", URL_hash, URL_setHash)
}

func URLConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	url, errArg1 := js.ConsumeArgument(cbCtx, "url", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	base, found, errArg := js.ConsumeOptionalArg(cbCtx, "base", codec.DecodeString)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return CreateURLBase(cbCtx, url, base)
	}
	return CreateURL(cbCtx, url)
}

func URL_toJSON[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func URL_href[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return codec.EncodeString(cbCtx, result)
}

func URL_setHref[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func URL_origin[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return codec.EncodeString(cbCtx, result)
}

func URL_protocol[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return codec.EncodeString(cbCtx, result)
}

func URL_setProtocol[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_setProtocol: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func URL_username[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_username: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func URL_setUsername[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_setUsername: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func URL_password[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_password: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func URL_setPassword[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_setPassword: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func URL_host[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return codec.EncodeString(cbCtx, result)
}

func URL_setHost[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func URL_hostname[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return codec.EncodeString(cbCtx, result)
}

func URL_setHostname[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func URL_port[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return codec.EncodeString(cbCtx, result)
}

func URL_setPort[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func URL_pathname[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return codec.EncodeString(cbCtx, result)
}

func URL_setPathname[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func URL_search[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return codec.EncodeString(cbCtx, result)
}

func URL_setSearch[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[*url.URL](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetSearch(val)
	return nil, nil
}

func URL_searchParams[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_searchParams: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func URL_hash[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return codec.EncodeString(cbCtx, result)
}

func URL_setHash[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "URL.URL_setHash: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
