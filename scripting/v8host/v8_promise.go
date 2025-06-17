package v8host

import (
	"fmt"

	"github.com/gost-dom/v8go"
)

type v8Promise struct {
	jsValue
	*v8go.PromiseResolver
}

func newV8Promise(ctx *V8ScriptContext) v8Promise {
	promise, err := v8go.NewPromiseResolver(ctx.v8ctx)
	if err != nil {
		panic(fmt.Sprintf("newV8Promise: %v", err))
	}
	return v8Promise{newV8Value(ctx, promise.GetPromise().Value), promise}
}

func (p v8Promise) Resolve(val jsValue) { p.PromiseResolver.Resolve(val.Self().Value) }
func (p v8Promise) Reject(val jsValue)  { p.PromiseResolver.Reject(val.Self().Value) }
