package v8engine

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/gost-dom/browser/url"
	"github.com/gost-dom/v8go"
)

// V8Module represents a compiled ECMAScript module
type V8Module struct {
	ctx      *V8ScriptContext
	module   *v8go.Module
	logger   *slog.Logger
	location *url.URL
}

func (mod V8Module) Run() error {
	p, err := mod.module.Evaluate(mod.ctx.v8ctx)
	if err != nil {
		err = fmt.Errorf("gost-dom/v8engine: run module: %w", err)
	} else {
		if _, err = awaitPromise(mod.ctx.v8ctx, p); err != nil {
			err = fmt.Errorf("gost-dom/v8engine: awaiting module: %w", err)
		}
	}
	mod.ctx.tick()
	return err
}

func (mod V8Module) Eval() (any, error) {
	return nil, mod.Run()
}

func awaitPromise(ctx *v8go.Context, val *v8go.Value) (*v8go.Value, error) {
	timeout := time.After(time.Second)
	resolve := make(chan *v8go.Value, 1)
	reject := make(chan error, 1)

	if !val.IsPromise() {
		return val, nil
	}
	p, _ := val.AsPromise()

	p.Then(func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		args := info.Args()
		if len(args) > 0 {
			resolve <- args[0]
			return args[0]
		} else {
			resolve <- nil
			return nil
		}
	}, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		args := info.Args()
		if len(args) > 0 && args[0] != nil {
			val := args[0]
			if val.IsNativeError() {
				exc, _ := val.AsException()
				reject <- exc
			} else {
				reject <- fmt.Errorf("Non-Error rejection: %s", val.String())
			}
		} else {
			reject <- errors.New("Suspicious reject call. No argument provided")
		}
		return nil
	})
	ctx.PerformMicrotaskCheckpoint()

	select {
	case <-timeout:
		return nil, errors.New("Timeout waiting for promise")
	case val := <-resolve:
		return val, nil
	case err := <-reject:
		return nil, err
	}
}
