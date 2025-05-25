// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

type parentNodeV8Wrapper struct {
	handleReffedObject[dom.ParentNode]
}

func newParentNodeV8Wrapper(scriptHost *V8ScriptHost) *parentNodeV8Wrapper {
	return &parentNodeV8Wrapper{newHandleReffedObject[dom.ParentNode](scriptHost)}
}

func createParentNodePrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newParentNodeV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w parentNodeV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("prepend", wrapV8Callback(w.scriptHost, w.prepend))
	prototypeTmpl.Set("append", wrapV8Callback(w.scriptHost, w.append))
	prototypeTmpl.Set("replaceChildren", wrapV8Callback(w.scriptHost, w.replaceChildren))
	prototypeTmpl.Set("querySelector", wrapV8Callback(w.scriptHost, w.querySelector))
	prototypeTmpl.Set("querySelectorAll", wrapV8Callback(w.scriptHost, w.querySelectorAll))

	prototypeTmpl.SetAccessorProperty("firstElementChild",
		wrapV8Callback(w.scriptHost, w.firstElementChild),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("lastElementChild",
		wrapV8Callback(w.scriptHost, w.lastElementChild),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("childElementCount",
		wrapV8Callback(w.scriptHost, w.childElementCount),
		nil,
		v8.None)
}

func (w parentNodeV8Wrapper) Constructor(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: ParentNode.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w parentNodeV8Wrapper) querySelector(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: ParentNode.querySelector")
	instance, err0 := js.As[dom.ParentNode](cbCtx.Instance())
	selectors, err1 := consumeArgument(cbCtx, "selectors", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.QuerySelector(selectors)
		if callErr != nil {
			return nil, callErr
		} else {
			return cbCtx.ScriptCtx().getInstanceForNode(result)
		}
	}
	return nil, errors.New("ParentNode.querySelector: Missing arguments")
}

func (w parentNodeV8Wrapper) querySelectorAll(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: ParentNode.querySelectorAll")
	instance, err0 := js.As[dom.ParentNode](cbCtx.Instance())
	selectors, err1 := consumeArgument(cbCtx, "selectors", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.QuerySelectorAll(selectors)
		if callErr != nil {
			return nil, callErr
		} else {
			return w.toNodeList(cbCtx.ScriptCtx(), result)
		}
	}
	return nil, errors.New("ParentNode.querySelectorAll: Missing arguments")
}

func (w parentNodeV8Wrapper) firstElementChild(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: ParentNode.firstElementChild")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.FirstElementChild()
	return cbCtx.ScriptCtx().getInstanceForNode(result)
}

func (w parentNodeV8Wrapper) lastElementChild(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: ParentNode.lastElementChild")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LastElementChild()
	return cbCtx.ScriptCtx().getInstanceForNode(result)
}

func (w parentNodeV8Wrapper) childElementCount(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: ParentNode.childElementCount")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ChildElementCount()
	return w.toUnsignedLong(cbCtx.ScriptCtx(), result)
}
