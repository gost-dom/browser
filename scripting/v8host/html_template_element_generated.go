// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	abstraction "github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("HTMLTemplateElement", "HTMLElement", createHTMLTemplateElementPrototype)
}

type htmlTemplateElementV8Wrapper struct {
	handleReffedObject[html.HTMLTemplateElement]
}

func newHTMLTemplateElementV8Wrapper(scriptHost *V8ScriptHost) *htmlTemplateElementV8Wrapper {
	return &htmlTemplateElementV8Wrapper{newHandleReffedObject[html.HTMLTemplateElement](scriptHost)}
}

func createHTMLTemplateElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHTMLTemplateElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlTemplateElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {

	prototypeTmpl.SetAccessorProperty("content",
		wrapV8Callback(w.scriptHost, w.content),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootMode",
		wrapV8Callback(w.scriptHost, w.shadowRootMode),
		wrapV8Callback(w.scriptHost, w.setShadowRootMode),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootDelegatesFocus",
		wrapV8Callback(w.scriptHost, w.shadowRootDelegatesFocus),
		wrapV8Callback(w.scriptHost, w.setShadowRootDelegatesFocus),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootClonable",
		wrapV8Callback(w.scriptHost, w.shadowRootClonable),
		wrapV8Callback(w.scriptHost, w.setShadowRootClonable),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootSerializable",
		wrapV8Callback(w.scriptHost, w.shadowRootSerializable),
		wrapV8Callback(w.scriptHost, w.setShadowRootSerializable),
		v8.None)
}

func (w htmlTemplateElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLTemplateElement.Constructor")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlTemplateElementV8Wrapper) content(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLTemplateElement.content")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.HTMLTemplateElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Content()
	return cbCtx.Context().getInstanceForNode(result)
}

func (w htmlTemplateElementV8Wrapper) shadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLTemplateElement.shadowRootMode")
	return nil, errors.New("HTMLTemplateElement.shadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) setShadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLTemplateElement.setShadowRootMode")
	return nil, errors.New("HTMLTemplateElement.setShadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) shadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLTemplateElement.shadowRootDelegatesFocus")
	return nil, errors.New("HTMLTemplateElement.shadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) setShadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLTemplateElement.setShadowRootDelegatesFocus")
	return nil, errors.New("HTMLTemplateElement.setShadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) shadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLTemplateElement.shadowRootClonable")
	return nil, errors.New("HTMLTemplateElement.shadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) setShadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLTemplateElement.setShadowRootClonable")
	return nil, errors.New("HTMLTemplateElement.setShadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) shadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLTemplateElement.shadowRootSerializable")
	return nil, errors.New("HTMLTemplateElement.shadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlTemplateElementV8Wrapper) setShadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLTemplateElement.setShadowRootSerializable")
	return nil, errors.New("HTMLTemplateElement.setShadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
