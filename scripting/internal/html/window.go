package html

import (
	"errors"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func Window_window[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func Window_history[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	win, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return cbCtx.Constructor("History").NewInstance(win.History())
}

func Window_self[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func Window_parent[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.This(), nil
}

func Window_opener[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.Null(), nil
}

func Window_setOpener[T any](_ js.CallbackContext[T]) (js.Value[T], error) {
	return nil, errors.New("Not implemented")
}
