// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("NodeList", "", createNodeListPrototype)
}

type nodeListV8Wrapper struct {
	handleReffedObject[dom.NodeList]
}

func newNodeListV8Wrapper(scriptHost *V8ScriptHost) *nodeListV8Wrapper {
	return &nodeListV8Wrapper{newHandleReffedObject[dom.NodeList](scriptHost)}
}

func createNodeListPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newNodeListV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	wrapper.CustomInitialiser(constructor)
	return constructor
}
func (w nodeListV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("item", wrapV8Callback(w.scriptHost, w.item))

	prototypeTmpl.SetAccessorProperty("length",
		wrapV8Callback(w.scriptHost, w.length),
		nil,
		v8.None)
}

func (w nodeListV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NodeList.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nodeListV8Wrapper) item(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NodeList.item")
	instance, err0 := js.As[dom.NodeList](cbCtx.Instance())
	index, err1 := consumeArgument(cbCtx, "index", nil, w.decodeUnsignedLong)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result := instance.Item(index)
		return w.toJSWrapper(cbCtx, result)
	}
	return cbCtx.ReturnWithError(errors.New("NodeList.item: Missing arguments"))
}

func (w nodeListV8Wrapper) length(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: NodeList.length")
	instance, err := js.As[dom.NodeList](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx, result)
}
