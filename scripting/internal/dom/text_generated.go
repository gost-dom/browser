// This file is generated. Do not edit.

package dom

import (
	log "github.com/gost-dom/browser/internal/log"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"log/slog"
)

type Text[T any] struct{}

func NewText[T any](scriptHost js.ScriptEngine[T]) *Text[T] {
	return &Text[T]{}
}

func (wrapper Text[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Text[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("splitText", w.splitText)
	jsClass.CreatePrototypeAttribute("wholeText", w.wholeText, nil)
}

func (w Text[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Text"), slog.String("Method", "Constructor"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Text[T]) splitText(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Text"), slog.String("Method", "splitText"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Text.splitText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Text[T]) wholeText(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Text"), slog.String("Method", "wholeText"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Text.wholeText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
