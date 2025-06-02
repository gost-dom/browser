package js

import (
	"errors"
	"fmt"
)

// ConsumeArgument pulls one of the passed arguments and tries to convert it to
// target type T using one of the passed decoders. The return value will be
// taken from the first decode that does not return an error. If no decoder is
// succeeds, an error is returned.
//
// If no more arguments are present, or the next argument is undefined, the
// defaultValue function will be used if not nil; otherwise an error is
// returned.
//
// If the function returns with an error, the name will be used in the error
// message. Otherwise, name has ho effect on the function.
func ConsumeArgument[T, U any](
	args CallbackContext[U],
	name string,
	defaultValue func() T,
	decoders ...func(CallbackContext[U], Value[U]) (T, error),
) (result T, err error) {
	value, _ := args.ConsumeArg()
	if value == nil && defaultValue != nil {
		return defaultValue(), nil
	} else {
		errs := make([]error, len(decoders))
		if value != nil {
			for i, parser := range decoders {
				result, errs[i] = parser(args, value)
				if errs[i] == nil {
					return
				}
			}
		}
		// TODO: This should eventually become a TypeError in JS
		err = fmt.Errorf("tryParseArg: %s: %w", name, errors.Join(errs...))
		return
	}
}
