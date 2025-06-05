package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w *documentV8Wrapper[T]) CustomInitializer(class js.Class[T]) {
	// host := w.scriptHost
	// tmpl := constructor.InstanceTemplate()
	class.CreateInstanceAttribute("location",
		func(ctx js.CallbackContext[T]) (js.Value[T], error) {
			return ctx.Scope().GlobalThis().Get("location")
		}, nil)
	class.CreatePrototypeAttribute("head", w.head, nil)
	class.CreatePrototypeAttribute("body", w.body, nil)
	class.CreatePrototypeMethod("getElementById", w.getElementById)
}

func (w *documentV8Wrapper[T]) CreateInstance(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	res := dom.NewDocument(nil)
	cbCtx.This().SetNativeValue(res)
	cbCtx.Scope().SetValue(res, cbCtx.This())
	return nil, nil
}

func (w *documentV8Wrapper[T]) getElementById(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err0 := js.As[dom.Document](cbCtx.Instance())
	id, err1 := consumeArgument(cbCtx, "id", nil, decodeString)
	if err := errors.Join(err0, err1); err != nil {
		return nil, err
	}
	return encodeEntity(cbCtx, instance.GetElementById(id))
}

func (w *documentV8Wrapper[T]) head(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return encodeEntity(cbCtx, instance.Head())
}

func (w *documentV8Wrapper[T]) body(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return encodeEntity(cbCtx, instance.Body())
}

func (w *documentV8Wrapper[T]) createElement(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	var name string
	name, err1 := consumeArgument(cbCtx, "name", nil, decodeString)
	instance, err2 := js.As[dom.Document](cbCtx.Instance())
	err := errors.Join(err1, err2)
	if err != nil {
		return nil, err
	}
	return encodeEntity(cbCtx, instance.CreateElement(name))
}

func (w *documentV8Wrapper[T]) createTextNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	data, err1 := consumeArgument(cbCtx, "data", nil, decodeString)
	instance, err2 := js.As[dom.Document](cbCtx.Instance())
	err := errors.Join(err1, err2)
	if err != nil {
		return nil, err
	}
	return encodeEntity(cbCtx, instance.CreateText(data))
}
