package v8host

import (
	"github.com/gost-dom/browser/dom"
	v8 "github.com/tommie/v8go"
)

func (w *parentNodeV8Wrapper) getNodesAndInstance(
	info *v8.FunctionCallbackInfo,
) (i dom.ParentNode, nodes []dom.Node, err error) {
	ctx := w.scriptHost.mustGetContext(info.Context())
	args := info.Args()
	nodes = make([]dom.Node, len(args))
	for idx, a := range args {
		if nodes[idx], err = w.decodeNodeOrText(ctx, a); err != nil {
			return
		}
	}
	i, err = w.getInstance(info)
	return
}

func (w *parentNodeV8Wrapper) append(info *v8.FunctionCallbackInfo) (v *v8.Value, err error) {
	if i, n, err := w.getNodesAndInstance(info); err == nil {
		err = i.Append(n...)
	}
	return
}

func (w *parentNodeV8Wrapper) prepend(info *v8.FunctionCallbackInfo) (res *v8.Value, err error) {
	if i, n, err := w.getNodesAndInstance(info); err == nil {
		err = i.Prepend(n...)
	}
	return
}

func (w *parentNodeV8Wrapper) replaceChildren(
	info *v8.FunctionCallbackInfo,
) (res *v8.Value, err error) {
	if i, n, err := w.getNodesAndInstance(info); err == nil {
		err = i.ReplaceChildren(n...)
	}
	return
}
