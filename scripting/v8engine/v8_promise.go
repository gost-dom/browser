package v8engine

import (
	"fmt"

	"github.com/gost-dom/v8go"
)

type v8Promise struct {
	ctx *V8ScriptContext
	jsValue
	*v8go.PromiseResolver
}

func newV8Promise(ctx *V8ScriptContext) v8Promise {
	promise, err := v8go.NewPromiseResolver(ctx.v8ctx)
	if err != nil {
		panic(fmt.Sprintf("newV8Promise: %v", err))
	}
	return v8Promise{ctx, newV8Value(ctx, promise.GetPromise().Value), promise}
}

func (p v8Promise) Resolve(val jsValue) {
	// TODO: This is a poor attempt at fixing a race condition. If the context
	// has been aborted, the v8 isolate could have been cleaned up, invalidating
	// the value we reject with. But this doesn't fix the problem, only reduce the
	// probability of a panic
	select {
	case <-p.ctx.Context().Done():
	default:
		p.PromiseResolver.Resolve(val.Self().v8Value())
	}
}

func (p v8Promise) Reject(val jsValue) {
	// TODO: This is a poor attempt at fixing a race condition. If the context
	// has been aborted, the v8 isolate could have been cleaned up, invalidating
	// the value we reject with. But this doesn't fix the problem, only reduce the
	// probability of a panic
	select {
	case <-p.ctx.Context().Done():
	default:
		p.PromiseResolver.Reject(val.Self().v8Value())
	}
}
