// This file is generated. Do not edit.

package html

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"log/slog"
)

type HTMLInputElement[T any] struct{}

func NewHTMLInputElement[T any](scriptHost js.ScriptEngine[T]) *HTMLInputElement[T] {
	return &HTMLInputElement[T]{}
}

func (wrapper HTMLInputElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLInputElement[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("checkValidity", w.checkValidity)
	jsClass.CreatePrototypeAttribute("name", w.name, w.setName)
	jsClass.CreatePrototypeAttribute("type", w.type_, w.setType)
	jsClass.CreatePrototypeAttribute("value", w.value, w.setValue)
}

func (w HTMLInputElement[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLInputElement"), slog.String("Method", "Constructor"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLInputElement[T]) checkValidity(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLInputElement"), slog.String("Method", "checkValidity"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CheckValidity()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w HTMLInputElement[T]) name(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLInputElement"), slog.String("Method", "name"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Name()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLInputElement[T]) setName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLInputElement"), slog.String("Method", "setName"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetName(val)
	return nil, nil
}

func (w HTMLInputElement[T]) type_(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLInputElement"), slog.String("Method", "type_"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLInputElement[T]) setType(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLInputElement"), slog.String("Method", "setType"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetType(val)
	return nil, nil
}

func (w HTMLInputElement[T]) value(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLInputElement"), slog.String("Method", "value"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLInputElement[T]) setValue(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLInputElement"), slog.String("Method", "setValue"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}
