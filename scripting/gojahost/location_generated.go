// This file is generated. Do not edit.

package gojahost

import (
	g "github.com/dop251/goja"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
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

func (w locationWrapper) Constructor(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.Constructor")
	cbCtx := newArgumentHelper(w.ctx, c)
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w locationWrapper) assign(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.assign")
	panic("Location.assign: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) replace(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.replace")
	panic("Location.replace: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) reload(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.reload")
	panic("Location.reload: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) href(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.href")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.Href()
	return cbCtx.ReturnWithValue(w.toString_(result))
}

func (w locationWrapper) setHref(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.setHref")
	panic("Location.setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) origin(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.origin")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.Origin()
	return cbCtx.ReturnWithValue(w.toString_(result))
}

func (w locationWrapper) protocol(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.protocol")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.Protocol()
	return cbCtx.ReturnWithValue(w.toString_(result))
}

func (w locationWrapper) setProtocol(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.setProtocol")
	panic("Location.setProtocol: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) host(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.host")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.Host()
	return cbCtx.ReturnWithValue(w.toString_(result))
}

func (w locationWrapper) setHost(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.setHost")
	panic("Location.setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) hostname(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.hostname")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.Hostname()
	return cbCtx.ReturnWithValue(w.toString_(result))
}

func (w locationWrapper) setHostname(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.setHostname")
	panic("Location.setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) port(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.port")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.Port()
	return cbCtx.ReturnWithValue(w.toString_(result))
}

func (w locationWrapper) setPort(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.setPort")
	panic("Location.setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) pathname(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.pathname")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.Pathname()
	return cbCtx.ReturnWithValue(w.toString_(result))
}

func (w locationWrapper) setPathname(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.setPathname")
	panic("Location.setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) search(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.search")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.Search()
	return cbCtx.ReturnWithValue(w.toString_(result))
}

func (w locationWrapper) setSearch(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.setSearch")
	panic("Location.setSearch: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) hash(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.hash")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.Hash()
	return cbCtx.ReturnWithValue(w.toString_(result))
}

func (w locationWrapper) setHash(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.setHash")
	panic("Location.setHash: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) ancestorOrigins(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Location.ancestorOrigins")
	panic("Location.ancestorOrigins: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
