package v8engine

import "github.com/gost-dom/v8go"

type v8Array struct {
	v8Value
	Object *v8go.Object
}

func newV8Array(ctx *V8ScriptContext, o *v8go.Object) jsArray {
	return &v8Array{v8Value{ctx, o.Value}, o}
}

func (a *v8Array) Push(v jsValue) error {
	push, err := a.Object.Get("push")
	if err != nil {
		return err
	}
	f, err := push.AsFunction()
	if err != nil {
		return err
	}
	_, err = f.Call(toV8Value(&a.v8Value), toV8Value(v))
	return err
}
