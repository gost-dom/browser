// This file is generated. Do not edit.

package dom

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
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
	jsClass.CreatePrototypeAttribute("localName", w.localName, nil)
	jsClass.CreatePrototypeAttribute("name", w.name, nil)
	jsClass.CreatePrototypeAttribute("value", w.value, w.setValue)
	jsClass.CreatePrototypeAttribute("ownerElement", w.ownerElement, nil)
}

func (w Attr[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Attr.Constructor", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Attr[T]) localName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Attr.localName", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LocalName()
	return codec.EncodeString(cbCtx, result)
}

func (w Attr[T]) name(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Attr.name", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Name()
	return codec.EncodeString(cbCtx, result)
}

func (w Attr[T]) value(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Attr.value", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return codec.EncodeString(cbCtx, result)
}

func (w Attr[T]) setValue(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Attr.setValue", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err0 := js.As[dom.Attr](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}

func (w Attr[T]) ownerElement(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Attr.ownerElement", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OwnerElement()
	return codec.EncodeEntity(cbCtx, result)
}
