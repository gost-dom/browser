package html

import (
	"github.com/gost-dom/browser/html"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w History[T]) decodeHistoryState(
	cbCtx js.Scope[T],
	val js.Value[T],
) (html.HistoryState, error) {
	return html.HistoryState(cbCtx.JSONStringify(val)), nil
}

func (w History[T]) toHistoryState(
	cbCtx js.CallbackContext[T],
	val htmlinterfaces.HistoryState,
) (js.Value[T], error) {
	return cbCtx.JSONParse(string(val))
}
