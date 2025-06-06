package v8host

import (
	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (n nodeV8Wrapper[T]) textContent(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	i, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}

	r := i.TextContent()
	return codec.EncodeString(cbCtx, r)
}

func (n nodeV8Wrapper[T]) setTextContent(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	i, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	arg, _ := cbCtx.ConsumeArg()
	i.SetTextContent(arg.String())
	return cbCtx.ReturnWithValue(nil)
}

func (n nodeV8Wrapper[T]) nodeType(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodeInt(cbCtx, int(instance.NodeType()))
}

func (n nodeV8Wrapper[T]) decodeGetRootNodeOptions(
	_ js.CallbackContext[T],
	value js.Value[T],
) (dom.GetRootNodeOptions, error) {
	return dom.GetRootNodeOptions(value.Boolean()), nil
}

func (n nodeV8Wrapper[T]) defaultGetRootNodeOptions() dom.GetRootNodeOptions {
	return false
}

func (w nodeV8Wrapper[T]) defaultboolean() bool {
	return false
}
