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

func (w domTokenListV8Wrapper) Constructor(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w domTokenListV8Wrapper) item(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.item")
	instance, err0 := js.As[dom.DOMTokenList](cbCtx.Instance())
	index, err1 := consumeArgument(cbCtx, "index", nil, w.decodeUnsignedLong)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Item(index)
		return w.toNullableString_(cbCtx.ScriptCtx(), result)
	}
	return nil, errors.New("DOMTokenList.item: Missing arguments")
}

func (w domTokenListV8Wrapper) contains(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.contains")
	instance, err0 := js.As[dom.DOMTokenList](cbCtx.Instance())
	token, err1 := consumeArgument(cbCtx, "token", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Contains(token)
		return w.toBoolean(cbCtx.ScriptCtx(), result)
	}
	return nil, errors.New("DOMTokenList.contains: Missing arguments")
}

func (w domTokenListV8Wrapper) add(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.add")
	instance, err0 := js.As[dom.DOMTokenList](cbCtx.Instance())
	tokens, err1 := consumeArgument(cbCtx, "tokens", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.Add(tokens)
		return nil, callErr
	}
	return nil, errors.New("DOMTokenList.add: Missing arguments")
}

func (w domTokenListV8Wrapper) remove(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.remove")
	instance, err0 := js.As[dom.DOMTokenList](cbCtx.Instance())
	tokens, err1 := consumeArgument(cbCtx, "tokens", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.Remove(tokens)
		return nil, nil
	}
	return nil, errors.New("DOMTokenList.remove: Missing arguments")
}

func (w domTokenListV8Wrapper) replace(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.replace")
	instance, err0 := js.As[dom.DOMTokenList](cbCtx.Instance())
	token, err1 := consumeArgument(cbCtx, "token", nil, w.decodeString)
	newToken, err2 := consumeArgument(cbCtx, "newToken", nil, w.decodeString)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		result := instance.Replace(token, newToken)
		return w.toBoolean(cbCtx.ScriptCtx(), result)
	}
	return nil, errors.New("DOMTokenList.replace: Missing arguments")
}

func (w domTokenListV8Wrapper) supports(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.supports")
	return cbCtx.ReturnWithError(errors.New("DOMTokenList.supports: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w domTokenListV8Wrapper) length(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.length")
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx.ScriptCtx(), result)
}

func (w domTokenListV8Wrapper) value(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.value")
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return w.toString_(cbCtx.ScriptCtx(), result)
}

func (w domTokenListV8Wrapper) setValue(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: DOMTokenList.setValue")
	instance, err0 := js.As[dom.DOMTokenList](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx.ScriptCtx(), cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}
