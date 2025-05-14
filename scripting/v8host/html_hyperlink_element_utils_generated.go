// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/gost-dom/v8go"
)

type htmlHyperlinkElementUtilsV8Wrapper struct {
	handleReffedObject[html.HTMLHyperlinkElementUtils]
}

func newHTMLHyperlinkElementUtilsV8Wrapper(scriptHost *V8ScriptHost) *htmlHyperlinkElementUtilsV8Wrapper {
	return &htmlHyperlinkElementUtilsV8Wrapper{newHandleReffedObject[html.HTMLHyperlinkElementUtils](scriptHost)}
}

func createHTMLHyperlinkElementUtilsPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHTMLHyperlinkElementUtilsV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlHyperlinkElementUtilsV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso

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
	prototypeTmpl.SetAccessorProperty("hash",
		v8.NewFunctionTemplateWithError(iso, w.hash),
		v8.NewFunctionTemplateWithError(iso, w.setHash),
		v8.None)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w htmlHyperlinkElementUtilsV8Wrapper) href(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.href")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return w.toUSVString(ctx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setHref(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.setHref")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeUSVString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHref(val)
	return nil, nil
}

func (w htmlHyperlinkElementUtilsV8Wrapper) origin(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.origin")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return w.toUSVString(ctx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) protocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.protocol")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return w.toUSVString(ctx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setProtocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.setProtocol")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeUSVString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetProtocol(val)
	return nil, nil
}

func (w htmlHyperlinkElementUtilsV8Wrapper) username(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.username")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Username()
	return w.toUSVString(ctx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setUsername(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.setUsername")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeUSVString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetUsername(val)
	return nil, nil
}

func (w htmlHyperlinkElementUtilsV8Wrapper) password(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.password")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Password()
	return w.toUSVString(ctx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setPassword(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.setPassword")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeUSVString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetPassword(val)
	return nil, nil
}

func (w htmlHyperlinkElementUtilsV8Wrapper) host(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.host")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return w.toUSVString(ctx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setHost(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.setHost")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeUSVString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHost(val)
	return nil, nil
}

func (w htmlHyperlinkElementUtilsV8Wrapper) hostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.hostname")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return w.toUSVString(ctx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setHostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.setHostname")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeUSVString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHostname(val)
	return nil, nil
}

func (w htmlHyperlinkElementUtilsV8Wrapper) port(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.port")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return w.toUSVString(ctx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setPort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.setPort")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeUSVString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetPort(val)
	return nil, nil
}

func (w htmlHyperlinkElementUtilsV8Wrapper) pathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.pathname")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return w.toUSVString(ctx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setPathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.setPathname")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeUSVString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetPathname(val)
	return nil, nil
}

func (w htmlHyperlinkElementUtilsV8Wrapper) search(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.search")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return w.toUSVString(ctx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setSearch(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.setSearch")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeUSVString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetSearch(val)
	return nil, nil
}

func (w htmlHyperlinkElementUtilsV8Wrapper) hash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.hash")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return w.toUSVString(ctx, result)
}

func (w htmlHyperlinkElementUtilsV8Wrapper) setHash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLHyperlinkElementUtils.setHash")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeUSVString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHash(val)
	return nil, nil
}
