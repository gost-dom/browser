package codec

import (
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/promise"
	"github.com/gost-dom/browser/internal/types"
	"github.com/gost-dom/browser/scripting/internal/js"
)

// getJSInstance gets the JavaScript object that wraps a specific Go object. If
// a wrapper already has been created, that wrapper is returned; otherwise a new
// object is created with the correct prototype configured.
func EncodeEntity[T any](scope js.Scope[T], e entity.ObjectIder) (js.Value[T], error) {
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

func EncodeEntityScopedWithPrototype[T any](
	scope js.Scope[T],
	e entity.ObjectIder,
	protoName string,
) (js.Value[T], error) {
	if e == nil {
		return scope.Null(), nil
	}

	if cached, ok := scope.GetValue(e); ok {
		return cached, nil
	}

	prototype := scope.Constructor(protoName)
	value, err := prototype.NewInstance(e)
	if err == nil {
		scope.SetValue(e, value)
	}
	return value, err
}

func EncodeBoolean[T any](s js.CallbackScope[T], b bool) (js.Value[T], error) {
	return s.NewBoolean(b), nil
}
func EncodeInt[T any](s js.CallbackScope[T], i int) (js.Value[T], error) {
	return s.NewInt32(int32(i)), nil
}

func EncodeString[T any](scope js.Scope[T], s string) (js.Value[T], error) {
	return scope.NewString(s), nil
}

func EncodeByteString[T any](scope js.Scope[T], s types.ByteString) (js.Value[T], error) {
	return scope.NewString(string(s)), nil
}

func EncodeNullableString[T any](scope js.CallbackScope[T], s *string) (js.Value[T], error) {
	if s != nil {
		return scope.NewString(*s), nil
	}
	return EncodeNull(scope)
}
func EncodeOptionalString[T any](
	scope js.CallbackScope[T],
	s string,
) (js.Value[T], error) {
	if s == "" {
		return EncodeNull(scope)
	}
	return scope.NewString(s), nil
}
func EncodeNillableString[T any](
	scope js.CallbackScope[T],
	s string,
	hasValue bool,
) (js.Value[T], error) {
	if hasValue {
		return scope.NewString(s), nil
	}
	return EncodeNull(scope)
}

func EncodeNull[T any](s js.CallbackScope[T]) (js.Value[T], error) {
	return s.Null(), nil
}

// EncodeConstrucedValue is a simple helper for JS constructor callbacks to
// store the constructed Go value in the JavaScript object, and possibly cache
// it with the script context.
func EncodeConstrucedValue[T any](s js.CallbackScope[T], val any) (js.Value[T], error) {
	// TODO: Figure out if this function should survive
	s.This().SetNativeValue(val)
	if e, ok := val.(entity.ObjectIder); ok {
		s.SetValue(e, s.This())
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
			} else {
				p.Reject(js.ToJsError(c, err))
			}
			return nil
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
	return EncodePromiseFunc(scope, func() (js.Value[T], error) {
		res := <-prom
		if res.Err == nil {
			return encoder(scope, res.Value)
		} else {
			return nil, res.Err
		}
	})
}
