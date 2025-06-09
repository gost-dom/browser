// This file is generated. Do not edit.

package url

import (
	"errors"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	url "github.com/gost-dom/browser/url"
)

func (wrapper URLV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w URLV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
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

func (w URLV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.Constructor")
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

func (w URLV8Wrapper[T]) toJSON(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.toJSON")
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

func (w URLV8Wrapper[T]) href(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.href")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return codec.EncodeString(cbCtx, result)
}

func (w URLV8Wrapper[T]) setHref(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHref")
	return nil, errors.New("URL.setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) origin(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.origin")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return codec.EncodeString(cbCtx, result)
}

func (w URLV8Wrapper[T]) protocol(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.protocol")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return codec.EncodeString(cbCtx, result)
}

func (w URLV8Wrapper[T]) setProtocol(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setProtocol")
	return nil, errors.New("URL.setProtocol: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) username(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.username")
	return nil, errors.New("URL.username: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) setUsername(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setUsername")
	return nil, errors.New("URL.setUsername: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) password(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.password")
	return nil, errors.New("URL.password: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) setPassword(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setPassword")
	return nil, errors.New("URL.setPassword: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) host(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.host")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return codec.EncodeString(cbCtx, result)
}

func (w URLV8Wrapper[T]) setHost(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHost")
	return nil, errors.New("URL.setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) hostname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.hostname")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return codec.EncodeString(cbCtx, result)
}

func (w URLV8Wrapper[T]) setHostname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHostname")
	return nil, errors.New("URL.setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) port(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.port")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return codec.EncodeString(cbCtx, result)
}

func (w URLV8Wrapper[T]) setPort(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setPort")
	return nil, errors.New("URL.setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) pathname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.pathname")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return codec.EncodeString(cbCtx, result)
}

func (w URLV8Wrapper[T]) setPathname(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setPathname")
	return nil, errors.New("URL.setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) search(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.search")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return codec.EncodeString(cbCtx, result)
}

func (w URLV8Wrapper[T]) setSearch(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setSearch")
	return nil, errors.New("URL.setSearch: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) searchParams(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.searchParams")
	return nil, errors.New("URL.searchParams: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w URLV8Wrapper[T]) hash(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.hash")
	instance, err := js.As[*url.URL](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return codec.EncodeString(cbCtx, result)
}

func (w URLV8Wrapper[T]) setHash(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: URL.setHash")
	return nil, errors.New("URL.setHash: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
