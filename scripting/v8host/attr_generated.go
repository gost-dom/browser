// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("Attr", "Node", newAttrV8Wrapper)
}

type attrV8Wrapper[T any] struct {
	handleReffedObject[dom.Attr, T]
}

func newAttrV8Wrapper(scriptHost jsScriptEngine) *attrV8Wrapper[jsTypeParam] {
	return &attrV8Wrapper[jsTypeParam]{newHandleReffedObject[dom.Attr, jsTypeParam](scriptHost)}
}

func (wrapper attrV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w attrV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("localName", w.localName, nil)
	jsClass.CreatePrototypeAttribute("name", w.name, nil)
	jsClass.CreatePrototypeAttribute("value", w.value, w.setValue)
	jsClass.CreatePrototypeAttribute("ownerElement", w.ownerElement, nil)
}

func (w attrV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Attr.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w attrV8Wrapper[T]) localName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Attr.localName")
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.LocalName()
	return w.toString_(cbCtx, result)
}

func (w attrV8Wrapper[T]) name(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Attr.name")
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Name()
	return w.toString_(cbCtx, result)
}

func (w attrV8Wrapper[T]) value(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Attr.value")
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Value()
	return w.toString_(cbCtx, result)
}

func (w attrV8Wrapper[T]) setValue(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Attr.setValue")
	instance, err0 := js.As[dom.Attr](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetValue(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w attrV8Wrapper[T]) ownerElement(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Attr.ownerElement")
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.OwnerElement()
	return encodeEntity(cbCtx, result)
}
