// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type domTokenListV8Wrapper[T any] struct {
	handleReffedObject[dom.DOMTokenList, T]
}

func newDOMTokenListV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *domTokenListV8Wrapper[T] {
	return &domTokenListV8Wrapper[T]{newHandleReffedObject[dom.DOMTokenList, T](scriptHost)}
}

func (wrapper domTokenListV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w domTokenListV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("item", w.item)
	jsClass.CreatePrototypeMethod("contains", w.contains)
	jsClass.CreatePrototypeMethod("add", w.add)
	jsClass.CreatePrototypeMethod("remove", w.remove)
	jsClass.CreatePrototypeMethod("toggle", w.toggle)
	jsClass.CreatePrototypeMethod("replace", w.replace)
	jsClass.CreatePrototypeMethod("supports", w.supports)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
	jsClass.CreatePrototypeAttribute("value", w.value, w.setValue)
	jsClass.CreatePrototypeMethod("toString", w.value)
}

func (w domTokenListV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: DOMTokenList.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w domTokenListV8Wrapper[T]) item(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: DOMTokenList.item")
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	index, errArg1 := consumeArgument(cbCtx, "index", nil, codec.DecodeInt)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.Item(index)
	return w.toNillableString_(cbCtx, result, hasValue)
}

func (w domTokenListV8Wrapper[T]) contains(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: DOMTokenList.contains")
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	token, errArg1 := consumeArgument(cbCtx, "token", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Contains(token)
	return w.toBoolean(cbCtx, result)
}

func (w domTokenListV8Wrapper[T]) add(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: DOMTokenList.add")
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	tokens, errArg1 := consumeRestArguments(cbCtx, "tokens", codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Add(tokens...)
	return nil, errCall
}

func (w domTokenListV8Wrapper[T]) replace(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: DOMTokenList.replace")
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	token, errArg1 := consumeArgument(cbCtx, "token", nil, codec.DecodeString)
	newToken, errArg2 := consumeArgument(cbCtx, "newToken", nil, codec.DecodeString)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.Replace(token, newToken)
	return w.toBoolean(cbCtx, result)
}

func (w domTokenListV8Wrapper[T]) supports(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: DOMTokenList.supports")
	return cbCtx.ReturnWithError(errors.New("DOMTokenList.supports: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w domTokenListV8Wrapper[T]) length(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: DOMTokenList.length")
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx, result)
}

func (w domTokenListV8Wrapper[T]) value(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: DOMTokenList.value")
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Value()
	return w.toString_(cbCtx, result)
}

func (w domTokenListV8Wrapper[T]) setValue(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: DOMTokenList.setValue")
	instance, err0 := js.As[dom.DOMTokenList](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetValue(val)
	return cbCtx.ReturnWithValue(nil)
}
