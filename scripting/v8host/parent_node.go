package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w *parentNodeV8Wrapper) getNodesAndInstance(
	cbCtx *v8CallbackContext,
) (i dom.ParentNode, nodes []dom.Node, err error) {
	args := cbCtx.consumeRest()
	nodes = make([]dom.Node, len(args))
	for idx, a := range args {
		if nodes[idx], err = w.decodeNodeOrText(cbCtx, &v8Value{cbCtx.iso(), a}); err != nil {
			return
		}
	}
	i, err = js.As[dom.ParentNode](cbCtx.Instance())
	return
}

func (w *parentNodeV8Wrapper) append(cbCtx *v8CallbackContext) (jsValue, error) {
	if instance, nodes, err := w.getNodesAndInstance(cbCtx); err == nil {
		if err = instance.Append(nodes...); err != nil {
			return cbCtx.ReturnWithError(err)
		}
	}
	return cbCtx.ReturnWithValue(nil)
}

func (w *parentNodeV8Wrapper) prepend(cbCtx *v8CallbackContext) (jsValue, error) {
	if instance, nodes, err := w.getNodesAndInstance(cbCtx); err == nil {
		if err = instance.Prepend(nodes...); err != nil {
			return cbCtx.ReturnWithError(err)
		}
	}
	return cbCtx.ReturnWithValue(nil)
}

func (w *parentNodeV8Wrapper) replaceChildren(cbCtx *v8CallbackContext) (jsValue, error) {
	if instance, nodes, err := w.getNodesAndInstance(cbCtx); err == nil {
		if err = instance.ReplaceChildren(nodes...); err != nil {
			return cbCtx.ReturnWithError(err)
		}
	}
	return cbCtx.ReturnWithValue(nil)
}
