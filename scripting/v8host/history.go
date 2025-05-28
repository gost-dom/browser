package v8host

import (
	"github.com/gost-dom/browser/html"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	v8 "github.com/gost-dom/v8go"
)

func (w historyV8Wrapper) defaultDelta() int {
	return 0
}

func (w historyV8Wrapper) defaultUrl() string {
	return ""
}

func (w historyV8Wrapper) decodeHistoryState(
	ctx jsCallbackContext,
	val jsValue,
) (html.HistoryState, error) {
	r, err := v8.JSONStringify(ctx.ScriptCtx().v8ctx, assertV8Value(val).v8Value())
	return html.HistoryState(r), err
}

func (w historyV8Wrapper) toHistoryState(
	cbCtx *v8CallbackContext,
	val htmlinterfaces.HistoryState,
) (jsValue, error) {
	return cbCtx.ReturnWithValueErr(v8.JSONParse(cbCtx.ScriptCtx().v8ctx, string(val)))
}
