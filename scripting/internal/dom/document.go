package dom

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w *DocumentV8Wrapper[T]) CustomInitializer(class js.Class[T]) {
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

func (w *DocumentV8Wrapper[T]) CreateInstance(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	res := dom.NewDocument(nil)
	cbCtx.This().SetNativeValue(res)
	cbCtx.Scope().SetValue(res, cbCtx.This())
	return nil, nil
}

func (w *DocumentV8Wrapper[T]) getElementById(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err0 := js.As[dom.Document](cbCtx.Instance())
	id, err1 := js.ConsumeArgument(cbCtx, "id", nil, codec.DecodeString)
	if err := errors.Join(err0, err1); err != nil {
		return nil, err
	}
	return codec.EncodeEntity(cbCtx, instance.GetElementById(id))
}

func (w *DocumentV8Wrapper[T]) head(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodeEntity(cbCtx, instance.Head())
}

func (w *DocumentV8Wrapper[T]) body(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodeEntity(cbCtx, instance.Body())
}

func (w *DocumentV8Wrapper[T]) createElement(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	var name string
	name, err1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	instance, err2 := js.As[dom.Document](cbCtx.Instance())
	err := errors.Join(err1, err2)
	if err != nil {
		return nil, err
	}
	return codec.EncodeEntity(cbCtx, instance.CreateElement(name))
}

func (w *DocumentV8Wrapper[T]) createTextNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	data, err1 := js.ConsumeArgument(cbCtx, "data", nil, codec.DecodeString)
	instance, err2 := js.As[dom.Document](cbCtx.Instance())
	err := errors.Join(err1, err2)
	if err != nil {
		return nil, err
	}
	return codec.EncodeEntity(cbCtx, instance.CreateText(data))
}
