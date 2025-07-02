package codec

import (
	"fmt"

	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/internal/js"
)

// EncodeCallbackErrorf create callback return values with a specific error
// message, and logs the error to the logger associated with the script context.
//
// While _uncaught_ JavaScript errors would be logged by default, if the error
// is caught by JavaScript, the error message will not be logged.
//
// This is particularly valuable for not-implemented methods, as JavaScript code
// will fail perfectly valid assumptions about the function are violated.
func EncodeCallbackErrorf[T any](
	ctx js.CallbackContext[T],
	format string,
	a ...any,
) (js.Value[T], error) {
	return nil, CallbackErrorf(ctx, format, a...)
}

func CallbackErrorf[T any](
	ctx js.CallbackContext[T],
	format string,
	a ...any,
) error {
	err := fmt.Errorf(format, a...)
	ctx.Logger().Error("JS Callback", log.ErrAttr(err))
	return err
}

// UnsupportedOptionErrorf is used to detect when script code uses options or
// arguments that are not yet supported by Gost-DOM.
//
// The primary usage is to identify when a failing test case is failing because
// of missing functionality in Gost-DOM, rather than an error in the system
// under test.
//
// The value argument represents a JavaScript function argument keyed option
// value. If it is not null or undefined, an error is logged and returned. The
// name of the api is specified in the webAPI, and the unsupported method or
// option is specified in key.
func UnsupportedOptionErrorf[T any](
	ctx js.CallbackContext[T],
	value js.Value[T],
	webAPI string,
	key string,
) error {
	if js.IsNullish(value) {
		return nil
	}
	return CallbackErrorf(ctx,
		"gost-dom/scripting/%s: %s: not yet supported",
		webAPI, key,
	)
}
