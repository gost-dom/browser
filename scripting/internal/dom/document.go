package dom

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w *Document[T]) CustomInitializer(class js.Class[T]) {
	// host := w.scriptHost
	// tmpl := constructor.InstanceTemplate()
	class.CreateInstanceAttribute("location",
		func(ctx js.CallbackContext[T]) (js.Value[T], error) {
			return ctx.GlobalThis().Get("location")
		}, nil)
	class.CreatePrototypeAttribute("head", w.head, nil)
	class.CreatePrototypeAttribute("body", w.body, nil)
	class.CreatePrototypeMethod("getElementById", w.getElementById)
}

func (w *Document[T]) CreateInstance(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	res := dom.NewDocument(nil)
	cbCtx.This().SetNativeValue(res)
	cbCtx.SetValue(res, cbCtx.This())
	return nil, nil
}

func (w *Document[T]) getElementById(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err0 := js.As[dom.Document](cbCtx.Instance())
	id, err1 := js.ConsumeArgument(cbCtx, "id", nil, codec.DecodeString)
	if err := errors.Join(err0, err1); err != nil {
		return nil, err
	}
	return codec.EncodeEntity(cbCtx, instance.GetElementById(id))
}

func (w *Document[T]) head(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodeEntity(cbCtx, instance.Head())
}

func (w *Document[T]) body(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodeEntity(cbCtx, instance.Body())
}

func (w *Document[T]) createElement(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	var name string
	name, err1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	instance, err2 := js.As[dom.Document](cbCtx.Instance())
	err := errors.Join(err1, err2)
	if err != nil {
		return nil, err
	}
	return codec.EncodeEntity(cbCtx, instance.CreateElement(name))
}
