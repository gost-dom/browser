package v8engine

import (
	"errors"

	"github.com/gost-dom/v8go"
)

type v8Function struct {
	v8Value
	v8fn *v8go.Function
}

func newFunction(ctx *V8ScriptContext, ft *v8go.FunctionTemplate) *v8Function {
	f := ft.GetFunction(ctx.v8ctx)
	return &v8Function{*newV8Value(ctx, f.Value), f}
}

func (f v8Function) Call(this jsObject, args ...jsValue) (jsValue, error) {
	v8Args := make([]v8go.Valuer, len(args))
	for i, a := range args {
		v8Args[i] = toV8Value(a)
	}
	var res jsValue
	v, err := f.v8fn.Call(assertV8Object(this).Object, v8Args...)
	err = errors.Join(err, f.ctx.tick())
	if err == nil {
		res = newV8Value(f.ctx, v)
	}
	return res, err
}
