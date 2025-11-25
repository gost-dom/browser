package v8engine

import (
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
)

type V8Error struct {
	*v8Value
	error
}

func newV8Error(ctx *V8ScriptContext, err error) js.Error[jsTypeParam] {
	exc := v8go.NewError(ctx.iso(), err.Error())
	val := &v8Value{ctx, exc.Value}
	return &V8Error{val, err}
}
