// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("HTMLAnchorElement", "HTMLElement", createHTMLAnchorElementPrototype)
}

type htmlAnchorElementV8Wrapper struct {
	handleReffedObject[html.HTMLAnchorElement]
	htmlHyperlinkElementUtils *htmlHyperlinkElementUtilsV8Wrapper
}

func newHTMLAnchorElementV8Wrapper(scriptHost *V8ScriptHost) *htmlAnchorElementV8Wrapper {
	return &htmlAnchorElementV8Wrapper{
		newHandleReffedObject[html.HTMLAnchorElement](scriptHost),
		newHTMLHyperlinkElementUtilsV8Wrapper(scriptHost),
	}
}

func createHTMLAnchorElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newHTMLAnchorElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlAnchorElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {

	prototypeTmpl.SetAccessorProperty("target",
		wrapV8Callback(w.scriptHost, w.target),
		wrapV8Callback(w.scriptHost, w.setTarget),
		v8.None)
	w.htmlHyperlinkElementUtils.installPrototype(prototypeTmpl)
}

func (w htmlAnchorElementV8Wrapper) Constructor(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLAnchorElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlAnchorElementV8Wrapper) target(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLAnchorElement.target")
	instance, err := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target()
	return w.toString_(cbCtx.ScriptCtx(), result)
}

func (w htmlAnchorElementV8Wrapper) setTarget(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLAnchorElement.setTarget")
	instance, err0 := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx.ScriptCtx(), cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTarget(val)
	return nil, nil
}
