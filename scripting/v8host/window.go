package v8host

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w *windowV8Wrapper[T]) window(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func (w *windowV8Wrapper[T]) history(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	win, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return cbCtx.Scope().Constructor("History").NewInstance(win.History())
}
