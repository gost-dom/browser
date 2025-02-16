package v8host

import (
	"github.com/gost-dom/browser/dom"
	v8 "github.com/tommie/v8go"
)

func (w *parentNodeV8Wrapper) getNodes(
	info *v8.FunctionCallbackInfo,
) (parentNode dom.ParentNode, nodes []dom.Node, err error) {
	ctx := w.scriptHost.mustGetContext(info.Context())
	args := info.Args()
	nodes = make([]dom.Node, len(args))
	for i, a := range args {
		if nodes[i], err = w.decodeNodeOrText(ctx, a); err != nil {
			return nil, nil, err
		}
	}
	i, err := w.getInstance(info)
	return i, nodes, err
}

func (w *parentNodeV8Wrapper) append(info *v8.FunctionCallbackInfo) (res *v8.Value, err error) {
	if i, nodes, err := w.getNodes(info); err != nil {
		i.Append(nodes...)
	}
	return
}

func (w *parentNodeV8Wrapper) prepend(info *v8.FunctionCallbackInfo) (res *v8.Value, err error) {
	if i, nodes, err := w.getNodes(info); err != nil {
		i.Prepend(nodes...)
	}
	return
}

func (w *parentNodeV8Wrapper) replaceChildren(
	info *v8.FunctionCallbackInfo,
) (res *v8.Value, err error) {
	if i, nodes, err := w.getNodes(info); err != nil {
		i.ReplaceChildren(nodes...)
	}
	return
}
