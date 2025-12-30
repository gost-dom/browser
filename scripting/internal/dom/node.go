package dom

import (
	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func encodeNodeType[T any](cbCtx js.CallbackContext[T], value dom.NodeType) (js.Value[T], error) {
	return codec.EncodeInt(cbCtx, int(value))
}

func decodeGetRootNodeOptions[T any](
	_ js.Scope[T],
	value js.Value[T],
) (dom.GetRootNodeOptions, error) {
	return dom.GetRootNodeOptions(value.Boolean()), nil
}
