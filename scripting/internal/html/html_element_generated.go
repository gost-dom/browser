// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"log/slog"
)

type HTMLElement[T any] struct {
	htmlOrSVGElement *HTMLOrSVGElement[T]
}

func NewHTMLElement[T any](scriptHost js.ScriptEngine[T]) *HTMLElement[T] {
	return &HTMLElement[T]{NewHTMLOrSVGElement(scriptHost)}
}

func (wrapper HTMLElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLElement[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("click", w.click)
	w.htmlOrSVGElement.installPrototype(jsClass)
}

func (w HTMLElement[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLElement"), slog.String("Method", "Constructor"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLElement[T]) click(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "HTMLElement"), slog.String("Method", "click"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[html.HTMLElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Click()
	return nil, nil
}
