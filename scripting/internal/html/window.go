package html

import (
	"errors"

	"github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w *Window[T]) window(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func (w *Window[T]) history(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	win, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return cbCtx.Constructor("History").NewInstance(win.History())
}

func (w *Window[T]) self(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func (w *Window[T]) parent(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func (w *Window[T]) opener(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.Null(), nil
}

func (w *Window[T]) setOpener(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return nil, errors.New("Not implemented")
}

func (w *Window[T]) toLocation(cbCtx js.CallbackContext[T], l html.Location) (js.Value[T], error) {
	if l == nil {
		return cbCtx.Null(), nil
	}
	return codec.EncodeEntity(cbCtx, l)
}
