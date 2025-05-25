package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func (n nodeV8Wrapper) textContent(cbCtx *argumentHelper) (*v8.Value, error) {
	i, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}

	r := i.TextContent()
	return v8.NewValue(n.iso(), r)
}

func (n nodeV8Wrapper) setTextContent(cbCtx *argumentHelper) (*v8.Value, error) {
	i, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	arg := cbCtx.consumeValue()
	i.SetTextContent(arg.String())
	return nil, nil
}

func (n nodeV8Wrapper) nodeType(cbCtx *argumentHelper) (*v8.Value, error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return v8.NewValue(n.scriptHost.iso, int32(instance.NodeType()))
}

func (n nodeV8Wrapper) decodeGetRootNodeOptions(
	ctx *V8ScriptContext,
	value *v8.Value,
) (dom.GetRootNodeOptions, error) {
	return dom.GetRootNodeOptions(value.Boolean()), nil
}

func (n nodeV8Wrapper) defaultGetRootNodeOptions() dom.GetRootNodeOptions {
	return false
}

func (w nodeV8Wrapper) defaultboolean() bool {
	return false
}
