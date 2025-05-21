// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("HTMLInputElement", "HTMLElement", createHTMLInputElementPrototype)
}

type htmlInputElementV8Wrapper struct {
	handleReffedObject[html.HTMLInputElement]
}

func newHTMLInputElementV8Wrapper(scriptHost *V8ScriptHost) *htmlInputElementV8Wrapper {
	return &htmlInputElementV8Wrapper{newHandleReffedObject[html.HTMLInputElement](scriptHost)}
}

func createHTMLInputElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHTMLInputElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlInputElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("checkValidity", v8.NewFunctionTemplateWithError(iso, w.checkValidity))

	prototypeTmpl.SetAccessorProperty("type",
		v8.NewFunctionTemplateWithError(iso, w.type_),
		v8.NewFunctionTemplateWithError(iso, w.setType),
		v8.None)
}

func (w htmlInputElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w htmlInputElementV8Wrapper) checkValidity(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLInputElement.checkValidity")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.CheckValidity()
	return w.toBoolean(ctx, result)
}

func (w htmlInputElementV8Wrapper) type_(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLInputElement.type")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Type()
	return w.toString_(ctx, result)
}

func (w htmlInputElementV8Wrapper) setType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLInputElement.setType")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetType(val)
	return nil, nil
}
