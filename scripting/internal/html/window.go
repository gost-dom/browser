package html

import (
	"github.com/gost-dom/browser/html"
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
