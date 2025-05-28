package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func (w *documentV8Wrapper) CustomInitialiser(constructor *v8.FunctionTemplate) {
	host := w.scriptHost
	// iso := host.iso
	tmpl := constructor.InstanceTemplate()
	tmpl.SetAccessorProperty(
		"location",
		v8.NewFunctionTemplateWithError(
			w.scriptHost.iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := w.scriptHost.mustGetContext(info.Context())
				return ctx.v8ctx.Global().Get("location")
			},
		),
		nil,
		v8.None,
	)
	proto := constructor.PrototypeTemplate()
	proto.SetAccessorProperty("head", wrapV8Callback(host, w.head), nil, v8.None)
	proto.SetAccessorProperty("body", wrapV8Callback(host, w.body), nil, v8.None)
	proto.Set("getElementById", wrapV8Callback(host, w.getElementById))
}

func (w *documentV8Wrapper) CreateInstance(cbCtx *v8CallbackContext) (jsValue, error) {
	return cbCtx.ReturnWithJSValueErr(
		w.store(dom.NewDocument(nil), cbCtx.ScriptCtx(), cbCtx.This()),
	)
}

func (w *documentV8Wrapper) getElementById(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err0 := js.As[dom.Document](cbCtx.Instance())
	id, err1 := consumeArgument(cbCtx, "id", nil, w.decodeString)
	if err := errors.Join(err0, err1); err != nil {
		return cbCtx.ReturnWithError(err)
	}
	return cbCtx.getInstanceForNode(instance.GetElementById(id))
}

func (w *documentV8Wrapper) head(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err == nil {
		return cbCtx.getInstanceForNode(instance.Head())
	} else {
		return cbCtx.ReturnWithError(err)
	}
}

func (w *documentV8Wrapper) body(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err == nil {
		return cbCtx.getInstanceForNode(instance.Body())
	} else {
		return cbCtx.ReturnWithError(err)
	}
}

func (w *documentV8Wrapper) toComment(
	cbCtx *v8CallbackContext,
	comment dom.Comment,
) (jsValue, error) {
	return cbCtx.getInstanceForNode(comment)
}

func (w *documentV8Wrapper) toAttr(cbCtx jsCallbackContext, comment dom.Attr) (jsValue, error) {
	return cbCtx.getInstanceForNode(comment)
}
func (w *documentV8Wrapper) createElement(cbCtx jsCallbackContext) (jsValue, error) {
	var name string
	name, err1 := cbCtx.consumeString()
	instance, err2 := js.As[dom.Document](cbCtx.Instance())
	err := errors.Join(err1, err2)
	if err == nil {
		e := instance.CreateElement(name)
		return cbCtx.getInstanceForNode(e)
	} else {
		return cbCtx.ReturnWithError(err)
	}
}
func (w *documentV8Wrapper) createTextNode(cbCtx jsCallbackContext) (jsValue, error) {
	var name string
	name, err1 := cbCtx.consumeString()
	instance, err2 := js.As[dom.Document](cbCtx.Instance())
	err := errors.Join(err1, err2)
	if err == nil {
		e := instance.CreateText(name)
		return cbCtx.getInstanceForNode(e)
	} else {
		return cbCtx.ReturnWithError(err)
	}
}
