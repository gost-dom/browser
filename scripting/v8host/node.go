package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (n nodeV8Wrapper) textContent(cbCtx jsCallbackContext) (jsValue, error) {
	i, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}

	r := i.TextContent()
	return n.toString_(cbCtx, r)
}

func (n nodeV8Wrapper) setTextContent(cbCtx jsCallbackContext) (jsValue, error) {
	i, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	arg, _ := cbCtx.ConsumeArg()
	i.SetTextContent(arg.String())
	return cbCtx.ReturnWithValue(nil)
}

func (n nodeV8Wrapper) nodeType(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return n.toUnsignedShort(cbCtx, int(instance.NodeType()))
}

func (n nodeV8Wrapper) decodeGetRootNodeOptions(
	cbCtx jsCallbackContext,
	value jsValue,
) (dom.GetRootNodeOptions, error) {
	return dom.GetRootNodeOptions(value.Boolean()), nil
}

func (n nodeV8Wrapper) defaultGetRootNodeOptions() dom.GetRootNodeOptions {
	return false
}

func (w nodeV8Wrapper) defaultboolean() bool {
	return false
}
