// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerJSClass("NamedNodeMap", "", createNamedNodeMapPrototype)
}

type namedNodeMapV8Wrapper struct {
	handleReffedObject[dom.NamedNodeMap, jsTypeParam]
}

func newNamedNodeMapV8Wrapper(scriptHost *V8ScriptHost) *namedNodeMapV8Wrapper {
	return &namedNodeMapV8Wrapper{newHandleReffedObject[dom.NamedNodeMap](scriptHost)}
}

func createNamedNodeMapPrototype(scriptHost *V8ScriptHost) v8Class {
	wrapper := newNamedNodeMapV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	wrapper.CustomInitializer(jsClass)
	return jsClass
}

func (w namedNodeMapV8Wrapper) installPrototype(jsClass v8Class) {
	jsClass.CreatePrototypeMethod("item", w.item)
	jsClass.CreatePrototypeMethod("getNamedItem", w.getNamedItem)
	jsClass.CreatePrototypeMethod("getNamedItemNS", w.getNamedItemNS)
	jsClass.CreatePrototypeMethod("setNamedItem", w.setNamedItem)
	jsClass.CreatePrototypeMethod("setNamedItemNS", w.setNamedItemNS)
	jsClass.CreatePrototypeMethod("removeNamedItem", w.removeNamedItem)
	jsClass.CreatePrototypeMethod("removeNamedItemNS", w.removeNamedItemNS)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
}

func (w namedNodeMapV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w namedNodeMapV8Wrapper) item(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.item")
	instance, errInst := js.As[dom.NamedNodeMap](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	index, errArg1 := consumeArgument(cbCtx, "index", nil, w.decodeUnsignedLong)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Item(index)
	return encodeEntity(cbCtx, result)
}

func (w namedNodeMapV8Wrapper) getNamedItem(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.getNamedItem")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.getNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) getNamedItemNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.getNamedItemNS")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.getNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) setNamedItem(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.setNamedItem")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.setNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) setNamedItemNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.setNamedItemNS")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.setNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) removeNamedItem(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.removeNamedItem")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.removeNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) removeNamedItemNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.removeNamedItemNS")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.removeNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) length(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.length")
	instance, err := js.As[dom.NamedNodeMap](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx, result)
}
