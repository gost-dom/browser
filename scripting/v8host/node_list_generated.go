// This file is generated. Do not edit.

package v8host

import (
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("NodeList", "", newNodeListV8Wrapper)
}

type nodeListV8Wrapper struct {
	handleReffedObject[dom.NodeList, jsTypeParam]
}

func newNodeListV8Wrapper(scriptHost jsScriptEngine) *nodeListV8Wrapper {
	return &nodeListV8Wrapper{newHandleReffedObject[dom.NodeList, jsTypeParam](scriptHost)}
}

func (wrapper nodeListV8Wrapper) Initialize(jsClass jsClass) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w nodeListV8Wrapper) installPrototype(jsClass jsClass) {
	jsClass.CreatePrototypeMethod("item", w.item)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
}

func (w nodeListV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NodeList.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nodeListV8Wrapper) item(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NodeList.item")
	instance, errInst := js.As[dom.NodeList](cbCtx.Instance())
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

func (w nodeListV8Wrapper) length(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NodeList.length")
	instance, err := js.As[dom.NodeList](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx, result)
}
