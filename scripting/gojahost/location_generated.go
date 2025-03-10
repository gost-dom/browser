// This file is generated. Do not edit.

package gojahost

import (
	g "github.com/dop251/goja"
	html "github.com/gost-dom/browser/html"
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
	prototype.Set("assign", w.assign)
	prototype.Set("replace", w.replace)
	prototype.Set("reload", w.reload)
	prototype.DefineAccessorProperty("href", w.ctx.vm.ToValue(w.href), w.ctx.vm.ToValue(w.setHref), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("origin", w.ctx.vm.ToValue(w.origin), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("protocol", w.ctx.vm.ToValue(w.protocol), w.ctx.vm.ToValue(w.setProtocol), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("host", w.ctx.vm.ToValue(w.host), w.ctx.vm.ToValue(w.setHost), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("hostname", w.ctx.vm.ToValue(w.hostname), w.ctx.vm.ToValue(w.setHostname), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("port", w.ctx.vm.ToValue(w.port), w.ctx.vm.ToValue(w.setPort), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("pathname", w.ctx.vm.ToValue(w.pathname), w.ctx.vm.ToValue(w.setPathname), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("search", w.ctx.vm.ToValue(w.search), w.ctx.vm.ToValue(w.setSearch), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("hash", w.ctx.vm.ToValue(w.hash), w.ctx.vm.ToValue(w.setHash), g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("ancestorOrigins", w.ctx.vm.ToValue(w.ancestorOrigins), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w locationWrapper) assign(c g.FunctionCall) g.Value {
	panic("Location.assign: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) replace(c g.FunctionCall) g.Value {
	panic("Location.replace: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) reload(c g.FunctionCall) g.Value {
	panic("Location.reload: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) href(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Href()
	return w.toUSVString(result)
}

func (w locationWrapper) setHref(c g.FunctionCall) g.Value {
	panic("Location.setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) origin(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Origin()
	return w.toUSVString(result)
}

func (w locationWrapper) protocol(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Protocol()
	return w.toUSVString(result)
}

func (w locationWrapper) setProtocol(c g.FunctionCall) g.Value {
	panic("Location.setProtocol: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) host(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Host()
	return w.toUSVString(result)
}

func (w locationWrapper) setHost(c g.FunctionCall) g.Value {
	panic("Location.setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) hostname(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Hostname()
	return w.toUSVString(result)
}

func (w locationWrapper) setHostname(c g.FunctionCall) g.Value {
	panic("Location.setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) port(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Port()
	return w.toUSVString(result)
}

func (w locationWrapper) setPort(c g.FunctionCall) g.Value {
	panic("Location.setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) pathname(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Pathname()
	return w.toUSVString(result)
}

func (w locationWrapper) setPathname(c g.FunctionCall) g.Value {
	panic("Location.setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) search(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Search()
	return w.toUSVString(result)
}

func (w locationWrapper) setSearch(c g.FunctionCall) g.Value {
	panic("Location.setSearch: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) hash(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Hash()
	return w.toUSVString(result)
}

func (w locationWrapper) setHash(c g.FunctionCall) g.Value {
	panic("Location.setHash: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w locationWrapper) ancestorOrigins(c g.FunctionCall) g.Value {
	panic("Location.ancestorOrigins: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
