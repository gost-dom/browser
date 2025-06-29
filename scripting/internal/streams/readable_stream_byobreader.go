// This file is generated. Do not edit.

package streams

import (
	"errors"

	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w ReadableStreamBYOBReader[T]) CreateInstance(
	cbCtx js.CallbackContext[T], _x string,
) (js.Value[T], error) {
	return nil, errors.New(
		"gost-dom/scripting/streams: ReadableStreamBYOBReader constructor not yet supported",
	)
}

func (w ReadableStreamBYOBReader[T]) decodeReadableStream(ctx js.CallbackContext[T], v js.Value[T]) (string, error) {
	if v != nil && v.Boolean() {
		return "", codec.CallbackErrorf(ctx,
			"gost-dom/scripting/streams: ReadableStreamBYOBReader readableStream not yet supported",
		)
	}
	return "", nil
}
