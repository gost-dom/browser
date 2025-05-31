package v8host

import (
	"github.com/gost-dom/browser/html"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	v8 "github.com/gost-dom/v8go"
)

func (w historyV8Wrapper) defaultDelta() int  { return 0 }
func (w historyV8Wrapper) defaultUrl() string { return "" }

func (w historyV8Wrapper) decodeHistoryState(
	ctx jsCallbackContext,
	val jsValue,
) (html.HistoryState, error) {
	r, err := v8.JSONStringify(ctx.ScriptCtx().v8ctx, assertV8Value(val).v8Value())
	return html.HistoryState(r), err
}

func (w historyV8Wrapper) toHistoryState(
	cbCtx jsCallbackContext,
	val htmlinterfaces.HistoryState,
) (jsValue, error) {
	v, err := v8.JSONParse(cbCtx.ScriptCtx().v8ctx, string(val))
	return newV8Value(cbCtx.iso(), v), err
}
