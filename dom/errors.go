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
	domErrorInvalid             int = 0
	index_size_err                  = 1
	domstring_size_err              = 2
	hierarchy_request_err           = 3
	wrong_document_err              = 4
	invalid_character_err           = 5
	no_data_allowed_err             = 6
	no_modification_allowed_err     = 7
	not_found_err                   = 8
	not_supported_err               = 9
	inuse_attribute_err             = 10
	invalid_state_err               = 11
	syntax_err                      = 12
	invalid_modification_err        = 13
	namespace_err                   = 14
	invalid_access_err              = 15
	validation_err                  = 16
	type_mismatch_err               = 17
	security_err                    = 18
	network_err                     = 19
	abort_err                       = 20
	url_mismatch_err                = 21
	quota_exceeded_err              = 22
	timeout_err                     = 23
	invalid_node_type_err           = 24
	data_clone_err                  = 25
)

// Deprecated: Use ErrDom
var ErrDOM = DOMError{}

var ErrDom = DOMError{}

// ErrSyntax is returned when adding an empty string to a [DOMTokenList]. This
// corresponds to a SyntaxError in JavaScript. This is a special case of a
// [DOMException]
//
// [DOMException]: https://developer.mozilla.org/en-US/docs/Web/API/DOMException
var ErrSyntax = DOMError{Code: syntax_err}

// ErrSyntax is returned when adding a token containing whitespace to a
// [DOMTokenList]. This corresponds to a SyntaxError in JavaScript. This is a
// special case of a [DOMException]
//
// [DOMException]: https://developer.mozilla.org/en-US/docs/Web/API/DOMException
var ErrInvalidCharacter = DOMError{Code: invalid_character_err}

// DOMError corresponds to [DOMException] in JavaScript
//
// [DOMException]: https://developer.mozilla.org/en-US/docs/Web/API/DOMException
type DOMError struct {
	msg  string
	Code int
}

func (e DOMError) Error() string { return e.msg }

func (e DOMError) Is(target error) bool {
	other, ok := target.(DOMError)
	if ok && other.Code > 0 {
		ok = other.Code == e.Code
	}
	return ok
}

// Deprecated: Unused, will be removed
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
	return ok && e.Code == invalid_character_err
}

func newSyntaxError(msg string) error {
	return DOMError{
		msg:  fmt.Sprintf("SyntaxError: %s", msg),
		Code: syntax_err,
	}
}

func newInvalidCharacterError(msg string) error {
	return DOMError{
		msg:  msg,
		Code: invalid_character_err,
	}
}
