package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func (w *documentV8Wrapper) CustomInitialiser(constructor *v8.FunctionTemplate) {
	host := w.scriptHost
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

func (w *documentV8Wrapper) CreateInstance(cbCtx jsCallbackContext) (jsValue, error) {
	res := dom.NewDocument(nil)
	// w.store(res, cbCtx)
	cbCtx.This().SetNativeValue(res)
	cbCtx.Scope().SetValue(res, cbCtx.This())
	return nil, nil
}

func (w *documentV8Wrapper) getElementById(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err0 := js.As[dom.Document](cbCtx.Instance())
	id, err1 := consumeArgument(cbCtx, "id", nil, w.decodeString)
	if err := errors.Join(err0, err1); err != nil {
		return nil, err
	}
	return encodeEntity(cbCtx, instance.GetElementById(id))
}

func (w *documentV8Wrapper) head(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err == nil {
		return encodeEntity(cbCtx, instance.Head())
	} else {
		return nil, err
	}
}

func (w *documentV8Wrapper) body(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err == nil {
		return encodeEntity(cbCtx, instance.Body())
	} else {
		return nil, err
	}
}

func (w *documentV8Wrapper) createElement(cbCtx jsCallbackContext) (jsValue, error) {
	var name string
	name, err1 := cbCtx.consumeString()
	instance, err2 := js.As[dom.Document](cbCtx.Instance())
	err := errors.Join(err1, err2)
	if err == nil {
		return encodeEntity(cbCtx, instance.CreateElement(name))
	} else {
		return nil, err
	}
}
func (w *documentV8Wrapper) createTextNode(cbCtx jsCallbackContext) (jsValue, error) {
	var name string
	name, err1 := cbCtx.consumeString()
	instance, err2 := js.As[dom.Document](cbCtx.Instance())
	err := errors.Join(err1, err2)
	if err == nil {
		return encodeEntity(cbCtx, instance.CreateText(name))
	} else {
		return nil, err
	}
}
