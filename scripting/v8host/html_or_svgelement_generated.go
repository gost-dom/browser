// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/gost-dom/v8go"
)

type htmlOrSVGElementV8Wrapper struct {
	handleReffedObject[html.HTMLOrSVGElement]
}

func newHTMLOrSVGElementV8Wrapper(scriptHost *V8ScriptHost) *htmlOrSVGElementV8Wrapper {
	return &htmlOrSVGElementV8Wrapper{newHandleReffedObject[html.HTMLOrSVGElement](scriptHost)}
}

func createHTMLOrSVGElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHTMLOrSVGElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlOrSVGElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("focus", v8.NewFunctionTemplateWithError(iso, w.focus))
	prototypeTmpl.Set("blur", v8.NewFunctionTemplateWithError(iso, w.blur))

	prototypeTmpl.SetAccessorProperty("dataset",
		v8.NewFunctionTemplateWithError(iso, w.dataset),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nonce",
		v8.NewFunctionTemplateWithError(iso, w.nonce),
		v8.NewFunctionTemplateWithError(iso, w.setNonce),
		v8.None)
	prototypeTmpl.SetAccessorProperty("autofocus",
		v8.NewFunctionTemplateWithError(iso, w.autofocus),
		v8.NewFunctionTemplateWithError(iso, w.setAutofocus),
		v8.None)
	prototypeTmpl.SetAccessorProperty("tabIndex",
		v8.NewFunctionTemplateWithError(iso, w.tabIndex),
		v8.NewFunctionTemplateWithError(iso, w.setTabIndex),
		v8.None)
}

func (w htmlOrSVGElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w htmlOrSVGElementV8Wrapper) blur(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLOrSVGElement.blur")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	instance.Blur()
	return nil, nil
}

func (w htmlOrSVGElementV8Wrapper) dataset(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLOrSVGElement.dataset")
	return nil, errors.New("HTMLOrSVGElement.dataset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlOrSVGElementV8Wrapper) nonce(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLOrSVGElement.nonce")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Nonce()
	return w.toString_(ctx, result)
}

func (w htmlOrSVGElementV8Wrapper) setNonce(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLOrSVGElement.setNonce")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetNonce(val)
	return nil, nil
}

func (w htmlOrSVGElementV8Wrapper) autofocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLOrSVGElement.autofocus")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Autofocus()
	return w.toBoolean(ctx, result)
}

func (w htmlOrSVGElementV8Wrapper) setAutofocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLOrSVGElement.setAutofocus")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeBoolean)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetAutofocus(val)
	return nil, nil
}

func (w htmlOrSVGElementV8Wrapper) tabIndex(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLOrSVGElement.tabIndex")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.TabIndex()
	return w.toLong(ctx, result)
}

func (w htmlOrSVGElementV8Wrapper) setTabIndex(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLOrSVGElement.setTabIndex")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeLong)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTabIndex(val)
	return nil, nil
}
