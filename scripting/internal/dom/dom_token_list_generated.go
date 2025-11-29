// This file is generated. Do not edit.

package dom

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	log "github.com/gost-dom/browser/internal/log"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"log/slog"
)

type DOMTokenList[T any] struct{}

func NewDOMTokenList[T any](scriptHost js.ScriptEngine[T]) *DOMTokenList[T] {
	return &DOMTokenList[T]{}
}

func (wrapper DOMTokenList[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w DOMTokenList[T]) installPrototype(jsClass js.Class[T]) {
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

func (w DOMTokenList[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "DOMTokenList"), slog.String("Method", "Constructor"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w DOMTokenList[T]) item(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "DOMTokenList"), slog.String("Method", "item"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	index, errArg1 := js.ConsumeArgument(cbCtx, "index", nil, codec.DecodeInt)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.Item(index)
	return codec.EncodeNillableString(cbCtx, result, hasValue)
}

func (w DOMTokenList[T]) contains(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "DOMTokenList"), slog.String("Method", "contains"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	token, errArg1 := js.ConsumeArgument(cbCtx, "token", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Contains(token)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w DOMTokenList[T]) add(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "DOMTokenList"), slog.String("Method", "add"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	tokens, errArg1 := js.ConsumeRestArguments(cbCtx, "tokens", codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Add(tokens...)
	return nil, errCall
}

func (w DOMTokenList[T]) remove(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "DOMTokenList"), slog.String("Method", "remove"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	tokens, errArg1 := js.ConsumeRestArguments(cbCtx, "tokens", codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.Remove(tokens...)
	return nil, nil
}

func (w DOMTokenList[T]) replace(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "DOMTokenList"), slog.String("Method", "replace"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	token, errArg1 := js.ConsumeArgument(cbCtx, "token", nil, codec.DecodeString)
	newToken, errArg2 := js.ConsumeArgument(cbCtx, "newToken", nil, codec.DecodeString)
	err = errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.Replace(token, newToken)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w DOMTokenList[T]) supports(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "DOMTokenList"), slog.String("Method", "supports"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "DOMTokenList.supports: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w DOMTokenList[T]) length(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "DOMTokenList"), slog.String("Method", "length"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}

func (w DOMTokenList[T]) value(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "DOMTokenList"), slog.String("Method", "value"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return codec.EncodeString(cbCtx, result)
}

func (w DOMTokenList[T]) setValue(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "DOMTokenList"), slog.String("Method", "setValue"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err0 := js.As[dom.DOMTokenList](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}
