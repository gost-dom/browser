package js

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/gosterror"
)

// Error is a "throwable" error. In JavaScript, you can throw any value. This
// includes interface error, making it a valid Go error value as well.
//
// Function ToJsError can convert sn error returned from a native Go function to
// Error[T].
//
// To keep the original go error, the implementing type **must** be able to keep
// an optional error value, that can be obtained through an Unwrap() function.
// Error[T] doesn't include the Unwrap() function as the return type is unknown
// at design time; it can be either error or []error.
type Error[T any] interface {
	Value[T]
	error
}

// ToJsError converts a Go error instance to a value the script engine can
// handle. Passing a nil value will return nil, making the function safe to use
// where it's not known at compile time if there is an error or not. If the
// error is already a js-engine compatible error, it is returned as is.
// TODO: The total list of possible errors needs to be configurable on the
// script engine, and not coupled to the error types known by the core browser.
func ToJsError[T any](s Scope[T], err error) (res Error[T]) {
	if err == nil {
		return nil
	}
	if errors.As(err, &res) {
		return
	}
	if toTypeError(s, err, &res) ||
		toDomException(s, err, &res) ||
		encodeAnyError(s, err, &res) {
		return
	}
	return s.NewError(err)
}

func encodeAnyError[T any](s Scope[T], err error, res *Error[T]) (ok bool) {
	var anyErr gosterror.AnyError
	if ok = errors.As(err, &anyErr); !ok {
		return
	}
	if val, isVal := anyErr.Reason.(Value[T]); isVal {
		*res = s.NewValueError(val, err)
		return
	}
	*res = s.NewError(err)
	return
}

func toTypeError[T any](s Scope[T], err error, res *Error[T]) bool {
	var e gosterror.TypeError
	if !errors.As(err, &e) {
		return false
	}
	*res = s.NewTypeError(err.Error())
	return true
}

func toDomException[T any](s Scope[T], err error, res *Error[T]) (ok bool) {
	var e dom.DOMError
	ok = errors.As(err, &e)
	if !ok {
		return
	}
	if cls := s.Constructor("DOMException"); cls == nil {
		*res = s.NewError(
			errors.Join(err, errors.New("DOMException ctor not found"), constants.ErrGostDomBug),
		)
	} else if obj, ctorErr := cls.NewInstance(e); ctorErr != nil {
		*res = s.NewError(errors.Join(err, fmt.Errorf("DOMException: %w", ctorErr), constants.ErrGostDomBug))
	} else {
		*res = s.NewValueError(obj, err)
	}
	return
}
