// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
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
	wrapper := newHTMLInputElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlInputElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("checkValidity", wrapV8Callback(w.scriptHost, w.checkValidity))

	prototypeTmpl.SetAccessorProperty("type",
		wrapV8Callback(w.scriptHost, w.type_),
		wrapV8Callback(w.scriptHost, w.setType),
		v8.None)
}

func (w htmlInputElementV8Wrapper) Constructor(cbCtx *argumentHelper) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLInputElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlInputElementV8Wrapper) checkValidity(cbCtx *argumentHelper) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLInputElement.checkValidity")
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.CheckValidity()
	return w.toBoolean(cbCtx, result)
}

func (w htmlInputElementV8Wrapper) type_(cbCtx *argumentHelper) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLInputElement.type_")
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Type()
	return w.toString_(cbCtx, result)
}

func (w htmlInputElementV8Wrapper) setType(cbCtx *argumentHelper) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLInputElement.setType")
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetType(val)
	return cbCtx.ReturnWithValue(nil)
}
