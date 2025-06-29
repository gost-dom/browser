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
	err := fmt.Errorf(format, a...)
	ctx.Logger().Error("JS Callback", log.ErrAttr(err))
	return nil, err
}
