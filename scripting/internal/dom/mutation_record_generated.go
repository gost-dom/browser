// This file is generated. Do not edit.

package dom

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	log "github.com/gost-dom/browser/internal/log"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"log/slog"
)

type MutationRecord[T any] struct{}

func NewMutationRecord[T any](scriptHost js.ScriptEngine[T]) *MutationRecord[T] {
	return &MutationRecord[T]{}
}

func (wrapper MutationRecord[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w MutationRecord[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("type", w.type_, nil)
	jsClass.CreatePrototypeAttribute("target", w.target, nil)
	jsClass.CreatePrototypeAttribute("addedNodes", w.addedNodes, nil)
	jsClass.CreatePrototypeAttribute("removedNodes", w.removedNodes, nil)
	jsClass.CreatePrototypeAttribute("previousSibling", w.previousSibling, nil)
	jsClass.CreatePrototypeAttribute("nextSibling", w.nextSibling, nil)
	jsClass.CreatePrototypeAttribute("attributeName", w.attributeName, nil)
	jsClass.CreatePrototypeAttribute("attributeNamespace", w.attributeNamespace, nil)
	jsClass.CreatePrototypeAttribute("oldValue", w.oldValue, nil)
}

func (w MutationRecord[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "MutationRecord"), slog.String("Method", "Constructor"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w MutationRecord[T]) type_(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "MutationRecord"), slog.String("Method", "type_"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return codec.EncodeString(cbCtx, result)
}

func (w MutationRecord[T]) target(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "MutationRecord"), slog.String("Method", "target"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecord[T]) addedNodes(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "MutationRecord"), slog.String("Method", "addedNodes"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AddedNodes
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecord[T]) removedNodes(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "MutationRecord"), slog.String("Method", "removedNodes"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.RemovedNodes
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecord[T]) previousSibling(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "MutationRecord"), slog.String("Method", "previousSibling"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecord[T]) nextSibling(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "MutationRecord"), slog.String("Method", "nextSibling"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecord[T]) attributeName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "MutationRecord"), slog.String("Method", "attributeName"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AttributeName
	return codec.EncodeNullableString(cbCtx, result)
}

func (w MutationRecord[T]) attributeNamespace(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "MutationRecord"), slog.String("Method", "attributeNamespace"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AttributeNamespace
	return codec.EncodeNullableString(cbCtx, result)
}

func (w MutationRecord[T]) oldValue(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "MutationRecord"), slog.String("Method", "oldValue"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OldValue
	return codec.EncodeNullableString(cbCtx, result)
}
