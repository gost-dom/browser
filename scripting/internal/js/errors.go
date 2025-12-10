package js

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/gosterror"
	"github.com/gost-dom/browser/internal/log"
)

// HandleJSCallbackError is to be called when calling into a JS callback function
// results in an error. Argument cbType represent the type of callback, e.g.,
// event handler, mutation observer, interval, etc.
func HandleJSCallbackError[T any](scope Scope[T], cbType string, err error) {
	scope.Logger().Error("Callback error", "callback-type", cbType, log.ErrAttr(err))
	scope.Window().DispatchEvent(event.NewErrorEvent(err))
}

// Error is a JavaScript representation of a go error instance.
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
	if res, ok := err.(Error[T]); ok {
		return res
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
	cls := s.Constructor("DOMException")
	if cls == nil {
		err = errors.Join(err, errors.New("DOMException ctor not found"), constants.ErrGostDomBug)
		*res = s.NewError(err)
		return
	}
	obj, ctorErr := cls.NewInstance(e)
	if ctorErr != nil {
		err = errors.Join(err, fmt.Errorf("DOMException: %w", ctorErr), constants.ErrGostDomBug)
		*res = s.NewError(err)
		return
	}
	var val Value[T] = obj
	*res = s.NewValueError(val, err)

	return
}
