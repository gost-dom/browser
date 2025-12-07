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

const (
	domErrorInvalid int = iota
	domErrorSyntaxError
	domErrorInvalidCharacter
	domErrorNotFound
)

// Deprecated: Use ErrDom
var ErrDOM = DOMError{}

var ErrDom = DOMError{}

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
	msg  string
	code int
}

func (e DOMError) Error() string { return e.msg }

func (e DOMError) Is(err error) bool {
	_, ok := err.(DOMError)
	return ok
}

type NotImplementedError error

func newDomError(msg string) error {
	return DOMError{msg, domErrorInvalid}
}

func newDomErrorCode(msg string, code int) error {
	return DOMError{msg, code}
}

// Deprecated: Prefer using Errors.Is(err, ErrDom)
func IsDOMError(err error) bool {
	return errors.Is(err, ErrDOM)
}

// Deprecated: Will be removed
func IsNotImplementedError(err error) bool {
	return false
}

// Deprecated: Will be removed
func IsSyntaxError(err error) bool {
	return errors.Is(err, ErrSyntax)
}

// Deprecated: Will be removed
func IsInvalidCharacterError(err error) bool {
	e, ok := err.(DOMError)
	return ok && e.code == domErrorInvalidCharacter
}

func newSyntaxError(msg string) error { return fmt.Errorf("%w: %s", ErrSyntax, msg) }
func newInvalidCharacterError(msg string) error {
	return fmt.Errorf("%w: %s", ErrInvalidCharacter, msg)
}
