// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
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
	wrapper := newHTMLTemplateElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

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

func (w htmlTemplateElementV8Wrapper) Constructor(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLTemplateElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlTemplateElementV8Wrapper) content(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLTemplateElement.content")
	instance, err := js.As[html.HTMLTemplateElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Content()
	return cbCtx.getInstanceForNode(result)
}

func (w htmlTemplateElementV8Wrapper) shadowRootMode(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootMode")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) setShadowRootMode(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootMode")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) shadowRootDelegatesFocus(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootDelegatesFocus")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) setShadowRootDelegatesFocus(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootDelegatesFocus")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) shadowRootClonable(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootClonable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) setShadowRootClonable(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootClonable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) shadowRootSerializable(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootSerializable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) setShadowRootSerializable(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootSerializable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
