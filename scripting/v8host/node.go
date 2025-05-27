package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func (n nodeV8Wrapper) textContent(cbCtx *argumentHelper) (jsValue, error) {
	i, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}

	r := i.TextContent()
	return cbCtx.ReturnWithValueErr(v8.NewValue(n.iso(), r))
}

func (n nodeV8Wrapper) setTextContent(cbCtx *argumentHelper) (jsValue, error) {
	i, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	arg := cbCtx.consumeValue()
	i.SetTextContent(arg.String())
	return cbCtx.ReturnWithValue(nil)
}

func (n nodeV8Wrapper) nodeType(cbCtx *argumentHelper) (jsValue, error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	return cbCtx.ReturnWithValueErr(v8.NewValue(n.scriptHost.iso, int32(instance.NodeType())))
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
