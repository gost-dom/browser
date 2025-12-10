package gosterror

import "fmt"

// AnyError wraps any value as a valid go [error] value. While errors originating
// from Go code will always be instances of error, in JavaScript, any value can
// be an error.
//
// When an error is generated in JavaScript code, and not representable directly
// as an error in Go, AnyError will represent the value.
type AnyError struct{ Reason any }

// Error implements the error interface
func (e AnyError) Error() string {
	return fmt.Sprintf("aborted: reason: %v", e.Reason)
}

// Unwrap supports error checking and casting behaviour of the [errors] package.
// Returns Reason if it is an error instance, otherwise it returns nil.
func (e AnyError) Unwrap() error {
	err, _ := e.Reason.(error)
	return err
}
