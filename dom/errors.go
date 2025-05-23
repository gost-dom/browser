// This defines some error constructors and code to check if they are of the
// right type. This could easily change to facilitate better interop with
// relevant error types in v8, and custom ES error classes. However, for now,
// this supports writing the code with clearly identifiable error types that
// correspond to error types defined in the specification.

package dom

import (
	"errors"
	"fmt"
)

// just a helper to avoid implementing Error() on all types
type browserError struct {
	base      string
	msg       string
	errorType int
}

const (
	errorTypeDOMError int = iota
	errorTypeSyntaxError
	errorTypeNotImplementedError
)

const (
	domErrorSyntaxError int = iota
	domErrorInvalidCharacter
	domErrorNotFound
)

var ErrDOM = DOMError{}

// ErrSyntax is returned when adding an empty string to a [DOMTokenList]. This
// corresponds to a SyntaxError in JavaScript. This is a special case of a
// [DOMException]
//
// [DOMException]: https://developer.mozilla.org/en-US/docs/Web/API/DOMException
var ErrSyntax = fmt.Errorf("%w: syntax error", ErrDOM)

// ErrSyntax is returned when adding a token containing whitespace to a
// [DOMTokenList]. This corresponds to a SyntaxError in JavaScript. This is a
// special case of a [DOMException]
//
// [DOMException]: https://developer.mozilla.org/en-US/docs/Web/API/DOMException
var ErrInvalidCharacter = fmt.Errorf("%w: invalid character", ErrDOM)

// DOMError corresponds to [DOMException] in JavaScript
//
// [DOMException]: https://developer.mozilla.org/en-US/docs/Web/API/DOMException
type DOMError struct {
	error
	code int
}

type NotImplementedError error

func newBrowserError(base string, msg string, errorType int) error {
	return browserError{base, msg, errorType}
}

func (e browserError) Error() string {
	return fmt.Sprintf("%s: %s", e.base, e.msg)
}

func newDomError(msg string) error {
	return DOMError{newBrowserError("DOMError", msg, errorTypeDOMError), -1}
}

func newDomErrorCode(msg string, code int) error {
	return DOMError{newBrowserError("DOMError", msg, errorTypeDOMError), code}
}

func IsDOMError(err error) bool {
	_, ok := err.(DOMError)
	return ok
}

func isBrowserErrorOfType(err error, errorType int) bool {
	browserError, ok := err.(browserError)
	return ok && browserError.errorType == errorType
}

func IsNotImplementedError(err error) bool {
	return isBrowserErrorOfType(err, errorTypeNotImplementedError)
}

func IsSyntaxError(err error) bool {
	return errors.Is(err, ErrSyntax)
	// e, ok := err.(DOMError)
	// return ok && e.code == domErrorSyntaxError
}

func IsInvalidCharacterError(err error) bool {
	e, ok := err.(DOMError)
	return ok && e.code == domErrorInvalidCharacter
}

func newSyntaxError(msg string) error { return fmt.Errorf("%w: %s", ErrSyntax, msg) }
func newInvalidCharacterError(msg string) error {
	return fmt.Errorf("%w: %s", ErrInvalidCharacter, msg)
}
