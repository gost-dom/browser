package js

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/internal/log"
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
	decoders ...func(Scope[U], Value[U]) (T, error),
) (result T, err error) {
	value, _ := args.ConsumeArg()
	if IsUndefined(value) && defaultValue != nil {
		return defaultValue(), nil
	} else {
		errs := make([]error, len(decoders))
		var lastErr error
		errCount := 0
		for i, parser := range decoders {
			result, errs[i] = parser(args, value)
			if errs[i] == nil {
				return
			}
			errCount++
			lastErr = errs[i]
		}
		if errCount == 1 {
			return result, fmt.Errorf("%s: %w", name, lastErr)
		}
		err = args.NewTypeError(fmt.Sprintf("invalid argument: %v", value))
		args.Logger().Warn("Error parsing JS argument", "name", name, log.ErrAttr(errors.Join(errs...)))
		return
	}
}

func IsUndefined[T any](v Value[T]) bool { return v == nil || v.IsUndefined() }

func ConsumeRestArguments[T, U any](
	args CallbackContext[U],
	name string,
	decoders ...func(Scope[U], Value[U]) (T, error),
) (results []T, err error) {
	errs := make([]error, len(decoders))
outer:
	for arg, ok := args.ConsumeArg(); ok; arg, ok = args.ConsumeArg() {
		for i, parser := range decoders {
			var result T
			result, errs[i] = parser(args, arg)
			if errs[i] == nil {
				results = append(results, result)
				continue outer
			}
		}
		err = errors.Join(errs...)
		if err != nil {
			err = fmt.Errorf("argument: %s: %w", name, err)
		}
		return
	}
	return
}

func ConsumeOptionalArg[T, U any](
	cbCtx CallbackContext[T],
	name string,
	decoders ...func(Scope[T], Value[T]) (U, error),
) (result U, found bool, err error) {
	value, _ := cbCtx.ConsumeArg()
	if IsUndefined(value) {
		return
	}
	found = true
	errs := make([]error, len(decoders))
	for i, parser := range decoders {
		result, errs[i] = parser(cbCtx, value)
		if errs[i] == nil {
			return
		}
	}
	// TODO: This should eventually become a TypeError in JS
	err = fmt.Errorf("tryParseArg: %s: %w", name, errors.Join(errs...))
	return
}
