// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	log "github.com/gost-dom/browser/internal/log"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"log/slog"
)

type ChildNode[T any] struct{}

func NewChildNode[T any](scriptHost js.ScriptEngine[T]) *ChildNode[T] {
	return &ChildNode[T]{}
}

func (wrapper ChildNode[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w ChildNode[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("before", w.before)
	jsClass.CreatePrototypeMethod("after", w.after)
	jsClass.CreatePrototypeMethod("replaceWith", w.replaceWith)
	jsClass.CreatePrototypeMethod("remove", w.remove)
}

func (w ChildNode[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "ChildNode"), slog.String("Method", "Constructor"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w ChildNode[T]) before(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "ChildNode"), slog.String("Method", "before"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "ChildNode.before: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ChildNode[T]) after(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "ChildNode"), slog.String("Method", "after"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "ChildNode.after: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ChildNode[T]) replaceWith(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "ChildNode"), slog.String("Method", "replaceWith"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "ChildNode.replaceWith: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ChildNode[T]) remove(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "ChildNode"), slog.String("Method", "remove"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[dom.ChildNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Remove()
	return nil, nil
}
