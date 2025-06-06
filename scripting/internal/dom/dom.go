package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w ParentNodeV8Wrapper[T]) decodeNodeOrText(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) (dom.Node, error) {
	if val.IsString() {
		return cbCtx.Scope().Window().Document().CreateTextNode(val.String()), nil
	}
	return codec.DecodeNode(cbCtx, val)
}
