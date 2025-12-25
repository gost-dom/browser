// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Attr[T any] struct{}

func NewAttr[T any](scriptHost js.ScriptEngine[T]) *Attr[T] {
	return &Attr[T]{}
}

func (wrapper Attr[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Attr[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateAttribute("localName", w.localName, nil)
	jsClass.CreateAttribute("name", w.name, nil)
	jsClass.CreateAttribute("value", w.value, w.setValue)
	jsClass.CreateAttribute("ownerElement", w.ownerElement, nil)
}

func (w Attr[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Attr[T]) localName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LocalName()
	return codec.EncodeString(cbCtx, result)
}

func (w Attr[T]) name(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Name()
	return codec.EncodeString(cbCtx, result)
}

func (w Attr[T]) value(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return codec.EncodeString(cbCtx, result)
}

func (w Attr[T]) setValue(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Attr](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}

func (w Attr[T]) ownerElement(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OwnerElement()
	return codec.EncodeEntity(cbCtx, result)
}
