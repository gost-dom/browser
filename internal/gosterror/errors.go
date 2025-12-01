package gosterror

import (
	"fmt"
)

// Represents a type that should become a TypeError in JavaScript. This is
// rarely created in Go code, but can happen through invalid combination of
// options. E.g., MutationObserver.Observe(options) will throw a TypeError if
// the options doesn't include one of ChildList, Attributes, or CharacterData.
type TypeError struct{ Message string }

func (e TypeError) Error() string {
	return fmt.Sprintf("TypeError: %s", e.Message)
}

func (e TypeError) Is(err error) bool {
	_, is := err.(TypeError)
	return is
}

func NewTypeError(msg string) error {
	return TypeError{msg}
}

var ErrTypeError = NewTypeError("TypeError")

func First(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
