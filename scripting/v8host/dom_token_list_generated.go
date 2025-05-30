// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("DOMTokenList", "", createDOMTokenListPrototype)
}

type domTokenListV8Wrapper struct {
	handleReffedObject[dom.DOMTokenList]
}

func newDOMTokenListV8Wrapper(scriptHost *V8ScriptHost) *domTokenListV8Wrapper {
	return &domTokenListV8Wrapper{newHandleReffedObject[dom.DOMTokenList](scriptHost)}
}

func createDOMTokenListPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newDOMTokenListV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	wrapper.CustomInitialiser(constructor)
	return constructor
}
func (w domTokenListV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("item", wrapV8Callback(w.scriptHost, w.item))
	prototypeTmpl.Set("contains", wrapV8Callback(w.scriptHost, w.contains))
	prototypeTmpl.Set("add", wrapV8Callback(w.scriptHost, w.add))
	prototypeTmpl.Set("remove", wrapV8Callback(w.scriptHost, w.remove))
	prototypeTmpl.Set("toggle", wrapV8Callback(w.scriptHost, w.toggle))
	prototypeTmpl.Set("replace", wrapV8Callback(w.scriptHost, w.replace))
	prototypeTmpl.Set("supports", wrapV8Callback(w.scriptHost, w.supports))

	prototypeTmpl.SetAccessorProperty("length",
		wrapV8Callback(w.scriptHost, w.length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("value",
		wrapV8Callback(w.scriptHost, w.value),
		wrapV8Callback(w.scriptHost, w.setValue),
		v8.None)
	prototypeTmpl.Set("toString", wrapV8Callback(w.scriptHost, w.value))
}

func (w domTokenListV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w domTokenListV8Wrapper) item(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.item")
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	index, errArg1 := consumeArgument(cbCtx, "index", nil, w.decodeUnsignedLong)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.Item(index)
	return w.toNillableString_(cbCtx, result, hasValue)
}

func (w domTokenListV8Wrapper) contains(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.contains")
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	token, errArg1 := consumeArgument(cbCtx, "token", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Contains(token)
	return w.toBoolean(cbCtx, result)
}

func (w domTokenListV8Wrapper) add(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.add")
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	tokens, errArg1 := consumeArgument(cbCtx, "tokens", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Add(tokens)
	return nil, errCall
}

func (w domTokenListV8Wrapper) remove(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.remove")
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	tokens, errArg1 := consumeArgument(cbCtx, "tokens", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.Remove(tokens)
	return nil, nil
}

func (w domTokenListV8Wrapper) replace(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.replace")
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	token, errArg1 := consumeArgument(cbCtx, "token", nil, w.decodeString)
	newToken, errArg2 := consumeArgument(cbCtx, "newToken", nil, w.decodeString)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.Replace(token, newToken)
	return w.toBoolean(cbCtx, result)
}

func (w domTokenListV8Wrapper) supports(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.supports")
	return cbCtx.ReturnWithError(errors.New("DOMTokenList.supports: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w domTokenListV8Wrapper) length(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.length")
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx, result)
}

func (w domTokenListV8Wrapper) value(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.value")
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Value()
	return w.toString_(cbCtx, result)
}

func (w domTokenListV8Wrapper) setValue(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.setValue")
	instance, err0 := js.As[dom.DOMTokenList](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetValue(val)
	return cbCtx.ReturnWithValue(nil)
}
