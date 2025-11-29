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

type HTMLAnchorElement[T any] struct {
	htmlHyperlinkElementUtils *HTMLHyperlinkElementUtils[T]
}

func NewHTMLAnchorElement[T any](scriptHost js.ScriptEngine[T]) *HTMLAnchorElement[T] {
	return &HTMLAnchorElement[T]{NewHTMLHyperlinkElementUtils(scriptHost)}
}

func (wrapper HTMLAnchorElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLAnchorElement[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("target", w.target, w.setTarget)
	w.htmlHyperlinkElementUtils.installPrototype(jsClass)
}

func (w HTMLAnchorElement[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLAnchorElement"), slog.String("Method", "Constructor"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLAnchorElement[T]) target(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLAnchorElement"), slog.String("Method", "target"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLAnchorElement[T]) setTarget(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLAnchorElement"), slog.String("Method", "setTarget"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err0 := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTarget(val)
	return nil, nil
}
