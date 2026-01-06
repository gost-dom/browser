// This file is generated. Do not edit.

package html

import (
	dom "github.com/gost-dom/browser/dom"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeElement[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("insertAdjacentHTML", Element_insertAdjacentHTML)
	jsClass.CreateAttribute("innerHTML", Element_innerHTML, Element_setInnerHTML)
	jsClass.CreateAttribute("outerHTML", Element_outerHTML, Element_setOuterHTML)
}

func Element_insertAdjacentHTML[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	position, errArg1 := js.ConsumeArgument(cbCtx, "position", nil, codec.DecodeString)
	string, errArg2 := js.ConsumeArgument(cbCtx, "string", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	errCall := instance.InsertAdjacentHTML(position, string)
	return nil, errCall
}

func Element_innerHTML[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.InnerHTML()
	return codec.EncodeString(cbCtx, result)
}

func Element_setInnerHTML[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Element](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	return nil, instance.SetInnerHTML(val)
}

func Element_outerHTML[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OuterHTML()
	return codec.EncodeString(cbCtx, result)
}

func Element_setOuterHTML[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Element](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	return nil, instance.SetOuterHTML(val)
}
