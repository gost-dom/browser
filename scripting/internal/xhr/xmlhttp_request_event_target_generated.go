// This file is generated. Do not edit.

package xhr

import (
	log "github.com/gost-dom/browser/internal/log"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"log/slog"
)

type XMLHttpRequestEventTarget[T any] struct{}

func NewXMLHttpRequestEventTarget[T any](scriptHost js.ScriptEngine[T]) *XMLHttpRequestEventTarget[T] {
	return &XMLHttpRequestEventTarget[T]{}
}

func (wrapper XMLHttpRequestEventTarget[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w XMLHttpRequestEventTarget[T]) installPrototype(jsClass js.Class[T]) {}

func (w XMLHttpRequestEventTarget[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "XMLHttpRequestEventTarget"), slog.String("Method", "Constructor"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}
