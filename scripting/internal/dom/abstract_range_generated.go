// This file is generated. Do not edit.

package dom

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeAbstractRange[T any](jsClass js.Class[T]) {
	jsClass.CreateAttribute("startContainer", AbstractRange_startContainer, nil)
	jsClass.CreateAttribute("startOffset", AbstractRange_startOffset, nil)
	jsClass.CreateAttribute("endContainer", AbstractRange_endContainer, nil)
	jsClass.CreateAttribute("endOffset", AbstractRange_endOffset, nil)
	jsClass.CreateAttribute("collapsed", AbstractRange_collapsed, nil)
}

func AbstractRange_startContainer[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.AbstractRange](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.StartContainer()
	return codec.EncodeEntity(cbCtx, result)
}

func AbstractRange_startOffset[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.AbstractRange](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.StartOffset()
	return codec.EncodeInt(cbCtx, result)
}

func AbstractRange_endContainer[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.AbstractRange](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.EndContainer()
	return codec.EncodeEntity(cbCtx, result)
}

func AbstractRange_endOffset[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.AbstractRange](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.EndOffset()
	return codec.EncodeInt(cbCtx, result)
}

func AbstractRange_collapsed[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.AbstractRange](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Collapsed()
	return codec.EncodeBoolean(cbCtx, result)
}
