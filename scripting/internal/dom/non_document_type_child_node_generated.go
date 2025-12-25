// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type NonDocumentTypeChildNode[T any] struct{}

func NewNonDocumentTypeChildNode[T any](scriptHost js.ScriptEngine[T]) *NonDocumentTypeChildNode[T] {
	return &NonDocumentTypeChildNode[T]{}
}

func (wrapper NonDocumentTypeChildNode[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w NonDocumentTypeChildNode[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("previousElementSibling", w.previousElementSibling, nil)
	jsClass.CreatePrototypeAttribute("nextElementSibling", w.nextElementSibling, nil)
}

func (w NonDocumentTypeChildNode[T]) previousElementSibling(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousElementSibling()
	return codec.EncodeEntity(cbCtx, result)
}

func (w NonDocumentTypeChildNode[T]) nextElementSibling(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextElementSibling()
	return codec.EncodeEntity(cbCtx, result)
}
