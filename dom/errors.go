// This defines some error constructors and code to check if they are of the
// right type. This could easily change to facilitate better interop with
// relevant error types in v8, and custom ES error classes. However, for now,
// this supports writing the code with clearly identifiable error types that
// correspond to error types defined in the specification.

package dom

import (
	"errors"
	"fmt"
	"strconv"
)

type domErrorCode int

const (
	domErrorInvalid domErrorCode = 0

	index_size_err              = 1
	domstring_size_err          = 2
	hierarchy_request_err       = 3
	wrong_document_err          = 4
	invalid_character_err       = 5
	no_data_allowed_err         = 6
	no_modification_allowed_err = 7
	not_found_err               = 8
	not_supported_err           = 9
	inuse_attribute_err         = 10
	invalid_state_err           = 11
	syntax_err                  = 12
	invalid_modification_err    = 13
	namespace_err               = 14
	invalid_access_err          = 15
	validation_err              = 16
	type_mismatch_err           = 17
	security_err                = 18
	network_err                 = 19
	abort_err                   = 20
	url_mismatch_err            = 21
	quota_exceeded_err          = 22
	timeout_err                 = 23
	invalid_node_type_err       = 24
	data_clone_err              = 25
)

func (c domErrorCode) String() string {
	switch c {
	case index_size_err:
		return "index_size_err"
	case domstring_size_err:
		return "domstring_size_err"
	case hierarchy_request_err:
		return "hierarchy_request_err"
	case wrong_document_err:
		return "wrong_document_err"
	case invalid_character_err:
		return "invalid_character_err"
	case no_data_allowed_err:
		return "no_data_allowed_err"
	case no_modification_allowed_err:
		return "no_modification_allowed_err"
	case not_found_err:
		return "not_found_err"
	case not_supported_err:
		return "not_supported_err"
	case inuse_attribute_err:
		return "inuse_attribute_err"
	case invalid_state_err:
		return "invalid_state_err"
	case syntax_err:
		return "SyntaxError"
	case invalid_modification_err:
		return "invalid_modification_err"
	case namespace_err:
		return "namespace_err"
	case invalid_access_err:
		return "invalid_access_err"
	case validation_err:
		return "validation_err"
	case type_mismatch_err:
		return "type_mismatch_err"
	case security_err:
		return "security_err"
	case network_err:
		return "network_err"
	case abort_err:
		return "abort_err"
	case url_mismatch_err:
		return "url_mismatch_err"
	case quota_exceeded_err:
		return "quota_exceeded_err"
	case timeout_err:
		return "timeout_err"
	case invalid_node_type_err:
		return "invalid_node_type_err"
	case data_clone_err:
		return "data_clone_err"
	default:
		return strconv.Itoa(int(c))
	}
}

// Deprecated: Use ErrDom
var ErrDOM = DOMError{}

var ErrDom = DOMError{}

// ErrSyntax is returned when adding an empty string to a [DOMTokenList]. This
// corresponds to a SyntaxError in JavaScript. This is a special case of a
// [DOMException]
//
// [DOMException]: https://developer.mozilla.org/en-US/docs/Web/API/DOMException
var ErrSyntax = DOMError{Message: "SyntaxError", Code: syntax_err}

// ErrInvalidCharacter is returned when adding a token containing whitespace to
// a [DOMTokenList]. This corresponds to a SyntaxError in JavaScript. This is a
// special case of a [DOMException]
//
// [DOMException]: https://developer.mozilla.org/en-US/docs/Web/API/DOMException
var ErrInvalidCharacter = DOMError{Code: invalid_character_err}

// DOMError corresponds to [DOMException] in JavaScript
//
// [DOMException]: https://developer.mozilla.org/en-US/docs/Web/API/DOMException
type DOMError struct {
	Message string
	Code    domErrorCode
}

func (e DOMError) Error() string { return e.Message }

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

func newDomErrorCode(msg string, code domErrorCode) error {
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

// Deprecated: Will be removed - use errors.Id(err, ErrInvalidCharacter)
func IsInvalidCharacterError(err error) bool {
	return errors.Is(err, ErrInvalidCharacter)
}

func newSyntaxError(msg string) error {
	return DOMError{
		Message: fmt.Sprintf("SyntaxError: %s", msg),
		Code:    syntax_err,
	}
}

func newInvalidCharacterError(msg string) error {
	return DOMError{
		Message: msg,
		Code:    invalid_character_err,
	}
}
