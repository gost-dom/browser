package v8host

import (
	"github.com/gost-dom/browser/html"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w historyV8Wrapper[T]) defaultDelta() int  { return 0 }
func (w historyV8Wrapper[T]) defaultUrl() string { return "" }

func (w historyV8Wrapper[T]) decodeHistoryState(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) (html.HistoryState, error) {
	return html.HistoryState(cbCtx.ValueFactory().JSONStringify(val)), nil
}

func (w historyV8Wrapper[T]) toHistoryState(
	cbCtx js.CallbackContext[T],
	val htmlinterfaces.HistoryState,
) (js.Value[T], error) {
	return cbCtx.ValueFactory().JSONParse(string(val))
}
