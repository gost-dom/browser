// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("NamedNodeMap", "", createNamedNodeMapPrototype)
}

type namedNodeMapV8Wrapper struct {
	handleReffedObject[dom.NamedNodeMap]
}

func newNamedNodeMapV8Wrapper(scriptHost *V8ScriptHost) *namedNodeMapV8Wrapper {
	return &namedNodeMapV8Wrapper{newHandleReffedObject[dom.NamedNodeMap](scriptHost)}
}

func createNamedNodeMapPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newNamedNodeMapV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	wrapper.CustomInitialiser(constructor)
	return constructor
}
func (w namedNodeMapV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("item", wrapV8Callback(w.scriptHost, w.item))
	prototypeTmpl.Set("getNamedItem", wrapV8Callback(w.scriptHost, w.getNamedItem))
	prototypeTmpl.Set("getNamedItemNS", wrapV8Callback(w.scriptHost, w.getNamedItemNS))
	prototypeTmpl.Set("setNamedItem", wrapV8Callback(w.scriptHost, w.setNamedItem))
	prototypeTmpl.Set("setNamedItemNS", wrapV8Callback(w.scriptHost, w.setNamedItemNS))
	prototypeTmpl.Set("removeNamedItem", wrapV8Callback(w.scriptHost, w.removeNamedItem))
	prototypeTmpl.Set("removeNamedItemNS", wrapV8Callback(w.scriptHost, w.removeNamedItemNS))

	prototypeTmpl.SetAccessorProperty("length",
		wrapV8Callback(w.scriptHost, w.length),
		nil,
		v8.None)
}

func (w namedNodeMapV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NamedNodeMap.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w namedNodeMapV8Wrapper) item(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NamedNodeMap.item")
	instance, errInst := js.As[dom.NamedNodeMap](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	index, errArg1 := consumeArgument(cbCtx, "index", nil, w.decodeUnsignedLong)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Item(index)
	return w.toJSWrapper(cbCtx, result)
}

func (w namedNodeMapV8Wrapper) getNamedItem(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NamedNodeMap.getNamedItem")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.getNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) getNamedItemNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NamedNodeMap.getNamedItemNS")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.getNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) setNamedItem(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NamedNodeMap.setNamedItem")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.setNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) setNamedItemNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NamedNodeMap.setNamedItemNS")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.setNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) removeNamedItem(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NamedNodeMap.removeNamedItem")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.removeNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) removeNamedItemNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NamedNodeMap.removeNamedItemNS")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.removeNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper) length(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NamedNodeMap.length")
	instance, err := js.As[dom.NamedNodeMap](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx, result)
}
