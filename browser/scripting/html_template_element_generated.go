// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func init() {
	RegisterJSClass("HTMLTemplateElement", "HTMLElement", createHtmlTemplateElementPrototype)
}

func createHtmlTemplateElementPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := newHtmlTemplateElementV8Wrapper(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()

	prototypeTmpl.SetAccessorProperty("content",
		v8.NewFunctionTemplateWithError(iso, wrapper.Content),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootMode",
		v8.NewFunctionTemplateWithError(iso, wrapper.ShadowRootMode),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootMode),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootDelegatesFocus",
		v8.NewFunctionTemplateWithError(iso, wrapper.ShadowRootDelegatesFocus),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootDelegatesFocus),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootClonable",
		v8.NewFunctionTemplateWithError(iso, wrapper.ShadowRootClonable),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootClonable),
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRootSerializable",
		v8.NewFunctionTemplateWithError(iso, wrapper.ShadowRootSerializable),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetShadowRootSerializable),
		v8.None)

	return constructor
}

func (e htmlTemplateElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(e.host.iso, "Illegal Constructor")
}

func (e htmlTemplateElementV8Wrapper) Content(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := e.host.MustGetContext(info.Context())
	instance, err := e.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Content()
	return ctx.GetInstanceForNode(result)
}

func (e htmlTemplateElementV8Wrapper) ShadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.ShadowRootMode")
}

func (e htmlTemplateElementV8Wrapper) SetShadowRootMode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootMode")
}

func (e htmlTemplateElementV8Wrapper) ShadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.ShadowRootDelegatesFocus")
}

func (e htmlTemplateElementV8Wrapper) SetShadowRootDelegatesFocus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootDelegatesFocus")
}

func (e htmlTemplateElementV8Wrapper) ShadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.ShadowRootClonable")
}

func (e htmlTemplateElementV8Wrapper) SetShadowRootClonable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootClonable")
}

func (e htmlTemplateElementV8Wrapper) ShadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.ShadowRootSerializable")
}

func (e htmlTemplateElementV8Wrapper) SetShadowRootSerializable(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: HTMLTemplateElement.SetShadowRootSerializable")
}
