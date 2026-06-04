package cssomview1

import js "github.com/gost-dom/browser/scripting/internal/js"

func createDummyBoundingRect[T any](cbCtx js.Scope[T]) (js.Value[T], error) {
	obj := cbCtx.NewObject()
	obj.Set("x", cbCtx.NewNumber(0))
	obj.Set("y", cbCtx.NewNumber(0))
	obj.Set("width", cbCtx.NewNumber(0))
	obj.Set("height", cbCtx.NewNumber(0))
	obj.Set("top", cbCtx.NewNumber(0))
	obj.Set("right", cbCtx.NewNumber(0))
	obj.Set("bottom", cbCtx.NewNumber(0))
	obj.Set("left", cbCtx.NewNumber(0))
	return obj, nil
}
