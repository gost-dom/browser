package v8engine

import (
	"github.com/gost-dom/v8go"
)

type V8Error struct {
	*v8Value
	exception *v8go.Exception
	error
}

func newV8Error(ctx *V8ScriptContext, err error) *V8Error {
	exc := v8go.NewError(ctx.iso(), err.Error())
	val := &v8Value{ctx, exc.Value}
	return &V8Error{val, exc, err}
}
