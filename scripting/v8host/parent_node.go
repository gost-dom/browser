package v8host

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	v8 "github.com/gost-dom/v8go"
)

func (w *parentNodeV8Wrapper) getNodesAndInstance(
	cbCtx *argumentHelper,
) (i dom.ParentNode, nodes []dom.Node, err error) {
	args := cbCtx.consumeRest()
	nodes = make([]dom.Node, len(args))
	for idx, a := range args {
		if nodes[idx], err = w.decodeNodeOrText(cbCtx.ScriptCtx(), a); err != nil {
			return
		}
	}
	i, err = abstraction.As[dom.ParentNode](cbCtx.Instance())
	return
}

func (w *parentNodeV8Wrapper) append(cbCtx *argumentHelper) (v *v8.Value, err error) {
	if i, n, err := w.getNodesAndInstance(cbCtx); err == nil {
		err = i.Append(n...)
	}
	return
}

func (w *parentNodeV8Wrapper) prepend(cbCtx *argumentHelper) (res *v8.Value, err error) {
	if i, n, err := w.getNodesAndInstance(cbCtx); err == nil {
		err = i.Prepend(n...)
	}
	return
}

func (w *parentNodeV8Wrapper) replaceChildren(cbCtx *argumentHelper) (res *v8.Value, err error) {
	if i, n, err := w.getNodesAndInstance(cbCtx); err == nil {
		err = i.ReplaceChildren(n...)
	}
	return
}
