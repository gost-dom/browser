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

type attrV8Wrapper struct {
	handleReffedObject[dom.Attr, jsTypeParam]
}

func newAttrV8Wrapper(scriptHost jsScriptEngine) *attrV8Wrapper {
	return &attrV8Wrapper{newHandleReffedObject[dom.Attr](scriptHost)}
}

func (wrapper attrV8Wrapper) Initialize(jsClass jsClass) {
	wrapper.installPrototype(jsClass)
}

func (w attrV8Wrapper) installPrototype(jsClass jsClass) {
	jsClass.CreatePrototypeAttribute("localName", w.localName, nil)
	jsClass.CreatePrototypeAttribute("name", w.name, nil)
	jsClass.CreatePrototypeAttribute("value", w.value, w.setValue)
	jsClass.CreatePrototypeAttribute("ownerElement", w.ownerElement, nil)
}

func (w attrV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Attr.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w attrV8Wrapper) localName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Attr.localName")
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.LocalName()
	return w.toString_(cbCtx, result)
}

func (w attrV8Wrapper) name(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Attr.name")
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Name()
	return w.toString_(cbCtx, result)
}

func (w attrV8Wrapper) value(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Attr.value")
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Value()
	return w.toString_(cbCtx, result)
}

func (w attrV8Wrapper) setValue(cbCtx jsCallbackContext) (jsValue, error) {
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

func (w attrV8Wrapper) ownerElement(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Attr.ownerElement")
	instance, err := js.As[dom.Attr](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.OwnerElement()
	return encodeEntity(cbCtx, result)
}
