// This file is generated. Do not edit.

package gojahost

import (
	"errors"
	g "github.com/dop251/goja"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	installClass("Location", "", newLocationWrapper)
}

type locationWrapper struct {
	baseInstanceWrapper[html.Location]
}

func newLocationWrapper(instance *GojaContext) wrapper {
	return &locationWrapper{newBaseInstanceWrapper[html.Location](instance)}
}

func (w locationWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.Set("assign", wrapCallback(w.ctx, w.assign))
	prototype.Set("replace", wrapCallback(w.ctx, w.replace))
	prototype.Set("reload", wrapCallback(w.ctx, w.reload))
	prototype.DefineAccessorProperty("href", wrapCallback(w.ctx, w.href), wrapCallback(w.ctx, w.setHref), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("origin", wrapCallback(w.ctx, w.origin), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("protocol", wrapCallback(w.ctx, w.protocol), wrapCallback(w.ctx, w.setProtocol), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("host", wrapCallback(w.ctx, w.host), wrapCallback(w.ctx, w.setHost), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("hostname", wrapCallback(w.ctx, w.hostname), wrapCallback(w.ctx, w.setHostname), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("port", wrapCallback(w.ctx, w.port), wrapCallback(w.ctx, w.setPort), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("pathname", wrapCallback(w.ctx, w.pathname), wrapCallback(w.ctx, w.setPathname), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("search", wrapCallback(w.ctx, w.search), wrapCallback(w.ctx, w.setSearch), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("hash", wrapCallback(w.ctx, w.hash), wrapCallback(w.ctx, w.setHash), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("ancestorOrigins", wrapCallback(w.ctx, w.ancestorOrigins), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w locationWrapper) assign(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.assign")
	return cbCtx.ReturnWithError(errors.New("Location.assign: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w locationWrapper) replace(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.replace")
	return cbCtx.ReturnWithError(errors.New("Location.replace: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w locationWrapper) reload(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.reload")
	return cbCtx.ReturnWithError(errors.New("Location.reload: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w locationWrapper) href(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.href")
	instance, instErr := js.As[html.Location](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Href()
	return cbCtx.ReturnWithValueErr(w.toString_(cbCtx, result))
}

func (w locationWrapper) setHref(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.setHref")
	return cbCtx.ReturnWithError(errors.New("Location.setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w locationWrapper) origin(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.origin")
	instance, instErr := js.As[html.Location](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Origin()
	return cbCtx.ReturnWithValueErr(w.toString_(cbCtx, result))
}

func (w locationWrapper) protocol(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.protocol")
	instance, instErr := js.As[html.Location](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Protocol()
	return cbCtx.ReturnWithValueErr(w.toString_(cbCtx, result))
}

func (w locationWrapper) setProtocol(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.setProtocol")
	return cbCtx.ReturnWithError(errors.New("Location.setProtocol: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w locationWrapper) host(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.host")
	instance, instErr := js.As[html.Location](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Host()
	return cbCtx.ReturnWithValueErr(w.toString_(cbCtx, result))
}

func (w locationWrapper) setHost(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.setHost")
	return cbCtx.ReturnWithError(errors.New("Location.setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w locationWrapper) hostname(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.hostname")
	instance, instErr := js.As[html.Location](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Hostname()
	return cbCtx.ReturnWithValueErr(w.toString_(cbCtx, result))
}

func (w locationWrapper) setHostname(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.setHostname")
	return cbCtx.ReturnWithError(errors.New("Location.setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w locationWrapper) port(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.port")
	instance, instErr := js.As[html.Location](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Port()
	return cbCtx.ReturnWithValueErr(w.toString_(cbCtx, result))
}

func (w locationWrapper) setPort(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.setPort")
	return cbCtx.ReturnWithError(errors.New("Location.setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w locationWrapper) pathname(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.pathname")
	instance, instErr := js.As[html.Location](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Pathname()
	return cbCtx.ReturnWithValueErr(w.toString_(cbCtx, result))
}

func (w locationWrapper) setPathname(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.setPathname")
	return cbCtx.ReturnWithError(errors.New("Location.setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w locationWrapper) search(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.search")
	instance, instErr := js.As[html.Location](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Search()
	return cbCtx.ReturnWithValueErr(w.toString_(cbCtx, result))
}

func (w locationWrapper) setSearch(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.setSearch")
	return cbCtx.ReturnWithError(errors.New("Location.setSearch: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w locationWrapper) hash(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.hash")
	instance, instErr := js.As[html.Location](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Hash()
	return cbCtx.ReturnWithValueErr(w.toString_(cbCtx, result))
}

func (w locationWrapper) setHash(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.setHash")
	return cbCtx.ReturnWithError(errors.New("Location.setHash: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w locationWrapper) ancestorOrigins(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Location.ancestorOrigins")
	return cbCtx.ReturnWithError(errors.New("Location.ancestorOrigins: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
