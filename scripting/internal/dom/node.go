package dom

import (
	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (n Node[T]) textContent(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	i, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}

	r := i.TextContent()
	return codec.EncodeString(cbCtx, r)
}

func (n Node[T]) setTextContent(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	i, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	arg, _ := cbCtx.ConsumeArg()
	i.SetTextContent(arg.String())
	return nil, nil
}

func encodeNodeType[T any](cbCtx js.CallbackContext[T], value dom.NodeType) (js.Value[T], error) {
	return codec.EncodeInt(cbCtx, int(value))
}

func decodeGetRootNodeOptions[T any](
	_ js.Scope[T],
	value js.Value[T],
) (dom.GetRootNodeOptions, error) {
	return dom.GetRootNodeOptions(value.Boolean()), nil
}

func (n Node[T]) defaultGetRootNodeOptions() dom.GetRootNodeOptions {
	return false
}

func (w Node[T]) defaultboolean() bool {
	return false
}
