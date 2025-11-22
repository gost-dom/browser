package sobekhost

import (
	"errors"

	"github.com/grafana/sobek"
)

type sobekModule struct {
	ctx    *scriptContext
	record sobek.ModuleRecord
}

func (m sobekModule) Eval() (any, error) {
	return nil, errors.New("Not implemented")
}

func (m sobekModule) Run() error {
	m.ctx.logger().Debug("Evaluate module", "vm", m.ctx.vm)
	// TODO: Handle promise return value

	// if err := m.record.Link(); err != nil {
	// 	return err
	// }
	p := m.record.Evaluate(m.ctx.vm)
	if p.State() != sobek.PromiseStateFulfilled {
		return p.Result().Export().(error)
	}
	return nil
}

// func awaitPromise(ctx *scriptContext, prom *sobek.Promise) (js.Value[jsTypeParam], error) {
// 	ctx.vm.ToValue(prom)
// 	timeout := time.After(time.Second)
// 	resolve := make(chan js.Value[jsTypeParam], 1)
// 	reject := make(chan error, 1)
//
// 	if val == nil {
// 		return nil, nil
// 	}
//
// 	obj, ok := val.AsObject()
// 	if !ok {
// 		return val, nil
// 	}
// 	then, err := obj.Get("then")
// 	if err != nil {
// 		return nil, err
// 	}
// 	fn, ok := then.AsFunction()
// 	if !ok {
// 		return val, nil
// 	}
// 	fn.Call(
// 		obj,
// 		newValue(ctx, wrapJSCallback(
// 			ctx,
// 			func(cb js.CallbackContext[jsTypeParam]) (js.Value[jsTypeParam], error) {
// 				arg, ok := cb.ConsumeArg()
// 				if ok {
// 					resolve <- arg
// 				} else {
// 					resolve <- nil
// 				}
// 				return nil, nil
// 			},
// 		)),
// 		newValue(ctx, wrapJSCallback(
// 			ctx,
// 			func(cb js.CallbackContext[jsTypeParam]) (js.Value[jsTypeParam], error) {
// 				val, ok := cb.ConsumeArg()
// 				if !ok {
// 					reject <- fmt.Errorf("Promise error with no value")
// 				} else {
// 					// TODO: if val is an error wrapping a native Go error, extract
// 					// the error
// 					reject <- fmt.Errorf("Promise error: %v", val)
// 				}
// 				return nil, nil
// 			})),
// 	)
//
// 	// p, _ := val.AsPromise()
// 	//
// 	// p.Then(func(info *v8go.FunctionCallbackInfo) *v8go.Value {
// 	// }, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
// 	// 	args := info.Args()
// 	// 	if len(args) > 0 && args[0] != nil {
// 	// 		val := args[0]
// 	// 		if val.IsNativeError() {
// 	// 			exc, _ := val.AsException()
// 	// 			reject <- exc
// 	// 		} else {
// 	// 			reject <- fmt.Errorf("Non-Error rejection: %s", val.String())
// 	// 		}
// 	// 	} else {
// 	// 		reject <- errors.New("Suspicious reject call. No argument provided")
// 	// 	}
// 	// 	return nil
// 	// })
// 	// ctx.PerformMicrotaskCheckpoint()
//
// 	select {
// 	case <-timeout:
// 		return nil, errors.New("Timeout waiting for promise")
// 	case val := <-resolve:
// 		return val, nil
// 	case err := <-reject:
// 		return nil, err
// 	}
// }
