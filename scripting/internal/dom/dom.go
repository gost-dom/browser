package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func decodeNodeOrText[T any](s js.Scope[T], val js.Value[T]) (dom.Node, error) {
	win, err := codec.GetWindow(s)
	if err != nil {
		return nil, err
	}
	if val.IsString() {
		return win.Document().CreateTextNode(val.String()), nil
	}
	return codec.DecodeNode(s, val)
}
