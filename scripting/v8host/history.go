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

func (w historyV8Wrapper) decodeAny(
	ctx *V8ScriptContext,
	val *v8.Value,
) (html.HistoryState, error) {
	r, err := v8.JSONStringify(ctx.v8ctx, val)
	return html.HistoryState(r), err
}

func (w historyV8Wrapper) toHistoryState(
	cbCtx *argumentHelper,
	val htmlinterfaces.HistoryState,
) (*v8.Value, error) {
	return v8.JSONParse(cbCtx.ScriptCtx().v8ctx, string(val))
}
