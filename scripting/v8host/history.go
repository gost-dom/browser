package v8host

import (
	"github.com/gost-dom/browser/html"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
)

func (w historyV8Wrapper) defaultDelta() int  { return 0 }
func (w historyV8Wrapper) defaultUrl() string { return "" }

func (w historyV8Wrapper) decodeHistoryState(
	ctx jsCallbackContext,
	val jsValue,
) (html.HistoryState, error) {
	return html.HistoryState(ctx.ValueFactory().JSONStringify(val)), nil
}

func (w historyV8Wrapper) toHistoryState(
	cbCtx jsCallbackContext,
	val htmlinterfaces.HistoryState,
) (jsValue, error) {
	return cbCtx.ValueFactory().JSONParse(string(val))
}
