package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w ParentNode[T]) decodeNodeOrText(
	cbCtx js.Scope[T],
	val js.Value[T],
) (dom.Node, error) {
	if val.IsString() {
		return cbCtx.Window().Document().CreateTextNode(val.String()), nil
	}
	return codec.DecodeNode(cbCtx, val)
}
