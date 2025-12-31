// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeAttr[T any](jsClass js.Class[T]) {
	jsClass.CreateAttribute("localName", Attr_localName, nil)
	jsClass.CreateAttribute("name", Attr_name, nil)
	jsClass.CreateAttribute("value", Attr_value, Attr_setValue)
	jsClass.CreateAttribute("ownerElement", Attr_ownerElement, nil)
}

func AttrConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func Attr_localName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LocalName()
	return codec.EncodeString(cbCtx, result)
}

func Attr_name[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Name()
	return codec.EncodeString(cbCtx, result)
}

func Attr_value[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return codec.EncodeString(cbCtx, result)
}

func Attr_setValue[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Attr](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}

func Attr_ownerElement[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OwnerElement()
	return codec.EncodeEntity(cbCtx, result)
}
