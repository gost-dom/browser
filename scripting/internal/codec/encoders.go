package codec

import (
	"errors"

	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/promise"
	"github.com/gost-dom/browser/scripting/internal/js"
)

// getJSInstance gets the JavaScript object that wraps a specific Go object. If
// a wrapper already has been created, that wrapper is returned; otherwise a new
// object is created with the correct prototype configured.
func EncodeEntity[T any](cbCtx js.Scope[T], e entity.ObjectIder) (js.Value[T], error) {
	return EncodeEntityScoped(cbCtx, e)
}

// TODO: Embed scope in CallbackScope, so only one function is necessary
func EncodeEntityScoped[T any](scope js.Scope[T], e entity.ObjectIder) (js.Value[T], error) {
	if e == nil {
		return scope.Null(), nil
	}

	if cached, ok := scope.GetValue(e); ok {
		return cached, nil
	}

	prototypeName := LookupJSPrototype(e)
	prototype := scope.Constructor(prototypeName)
	value, err := prototype.NewInstance(e)
	if err == nil {
		scope.SetValue(e, value)
	}
	return value, err
}

func EncodeBoolean[T any](cbCtx js.CallbackScope[T], b bool) (js.Value[T], error) {
	return cbCtx.NewBoolean(b), nil
}
func EncodeInt[T any](cbCtx js.CallbackScope[T], i int) (js.Value[T], error) {
	return cbCtx.NewInt32(int32(i)), nil
}

// TODO: Embed scope in CallbackScope, so only one function is necessary
func EncodeStringScoped[T any](cbCtx js.Scope[T], s string) (js.Value[T], error) {
	return cbCtx.NewString(s), nil
}

func EncodeString[T any](cbCtx js.CallbackScope[T], s string) (js.Value[T], error) {
	return cbCtx.NewString(s), nil
}

func EncodeNullableString[T any](
	cbCtx js.CallbackScope[T],
	s *string,
) (js.Value[T], error) {
	if s != nil {
		return cbCtx.NewString(*s), nil
	}
	return EncodeNull(cbCtx)
}
func EncodeNillableString[T any](
	cbCtx js.CallbackScope[T],
	s string,
	hasValue bool,
) (js.Value[T], error) {
	if hasValue {
		return cbCtx.NewString(s), nil
	}
	return EncodeNull(cbCtx)
}

func EncodeNull[T any](cbCtx js.CallbackScope[T]) (js.Value[T], error) {
	return cbCtx.Null(), nil
}

// EncodeConstrucedValue is a simple helper for JS constructor callbacks to
// store the constructed Go value in the JavaScript object, and possibly cache
// it with the script context.
func EncodeConstrucedValue[T any](cbCtx js.CallbackScope[T], val any) (js.Value[T], error) {
	// TODO: Figure out if this function should survive
	cbCtx.This().SetNativeValue(val)
	if e, ok := val.(entity.ObjectIder); ok {
		cbCtx.SetValue(e, cbCtx.This())
	}
	return nil, nil
}

// EncodePromiseFunc returnes a JavaScript Promise that will settle with the result
// of running function f. Function f must be safe to run concurrently, as it
// will execute in a separate goroutine.
//
// The promise will not settile immediately after f finishes, but will be
// deferred to run on the "main loop" that the embedder controls.
func EncodePromiseFunc[T any](
	c js.Scope[T],
	f func() (js.Value[T], error),
) (js.Value[T], error) {
	p := c.NewPromise()
	e := c.Clock().BeginEvent()
	go func() {
		r, err := f()
		e.AddEvent(func() error {
			if err == nil {
				p.Resolve(r)
				return nil
			} else {
				jsErr, err := EncodeError(c, err)
				p.Reject(jsErr)
				return err
			}
		})
	}()
	return p, nil
}

type Encoder[T, U any] = func(js.Scope[T], U) (js.Value[T], error)

// EncodePromise converts a [promise.Promise] value to a JavaScript Promise
// value, using encoder to convert the native fulfilled value to a JavaScript
// value.
//
// The returned Promise will not settile immediately after a value is received
// from prom, but will be deferred to run on the "main loop" that the embedder
// controls.
func EncodePromise[T, U any](
	scope js.Scope[T],
	prom promise.Promise[U],
	encoder Encoder[T, U],
) (js.Value[T], error) {
	p := scope.NewPromise()
	e := scope.Clock().BeginEvent()
	go func() {
		res := <-prom
		e.AddEvent(func() error {
			err := res.Err
			var val js.Value[T]
			if err == nil {
				val, err = encoder(scope, res.Value)
			}
			if err == nil {
				p.Resolve(val)
				return nil
			} else {
				errVal, err := EncodeError(scope, err)
				if errVal == nil {
					errVal = scope.Undefined()
				}
				p.Reject(errVal)
				return err
			}
		})
	}()
	return p, nil
}

func EncodeError[T any](scope js.Scope[T], err error) (js.Value[T], error) {
	if err == nil {
		return nil, nil
	}
	var errAny promise.ErrAny
	if errors.As(err, &errAny) {
		if v, ok := errAny.Reason.(js.Value[T]); ok {
			return v, nil
		}
	}
	return scope.NewError(err), nil
}
