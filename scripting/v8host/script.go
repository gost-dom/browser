package v8host

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/v8go"
)

type V8Script struct {
	ctx    *V8ScriptContext
	script string
}

func (s V8Script) Run() error {
	_, err := s.ctx.runScript(s.script)
	return err
}

func (s V8Script) Eval() (any, error) {
	var (
		result  *v8go.Value
		promise *v8go.Promise
		err     error
	)
	result, err = s.ctx.runScript(s.script)

	logResult(s.ctx.host.logger, result, err)
	if err != nil {
		return nil, err
	}
	if promise, err = result.AsPromise(); err == nil {
		if promise.State() != v8go.Pending {
			result, err = waitForPromise(s.ctx.v8ctx, promise)
		}
	}
	return v8ValueToGoValue(result)
}

func logResult(logger *slog.Logger, result *v8go.Value, err error) {
	if logger != nil {
		logger.Debug("Script executed",
			slog.Any("result", result),
			log.ErrAttr(err))
	}
}

func v8ValueToGoValue(result *v8go.Value) (any, error) {
	if result == nil {
		return nil, nil
	}
	if result.IsBoolean() {
		return result.Boolean(), nil
	}
	if result.IsInt32() {
		return result.Int32(), nil
	}
	if result.IsString() {
		return result.String(), nil
	}
	if result.IsNull() {
		return nil, nil
	}
	if result.IsUndefined() {
		return nil, nil
	}
	if o, err := result.AsObject(); err == nil {
		if c := o.InternalFieldCount(); c > 0 {
			f := o.GetInternalField(0)
			if f.IsExternal() {
				return f.ExternalHandle().Value(), nil
			}
		}
	}
	if result.IsArray() {
		obj, _ := result.AsObject()
		length, err := obj.Get("length")
		l := length.Uint32()
		errs := make([]error, l+1)
		errs[0] = err
		result := make([]any, l)
		for i := uint32(0); i < l; i++ {
			val, err := obj.GetIdx(i)
			if err == nil {
				result[i], err = v8ValueToGoValue(val)
			}
			errs[i+1] = err
		}
		return result, errors.Join(errs...)
	}
	return nil, fmt.Errorf("Value not yet supported: %v", *result)
}

func waitForPromise(ctx *v8go.Context, p *v8go.Promise) (*v8go.Value, error) {
	timeout := time.After(time.Second)
	resolve := make(chan *v8go.Value)
	reject := make(chan error)

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
	go ctx.PerformMicrotaskCheckpoint()

	select {
	case <-timeout:
		return nil, errors.New("Timeout waiting for promise")
	case val := <-resolve:
		return val, nil
	case err := <-reject:
		return nil, err
	}
}
