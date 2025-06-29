// This file is generated. Do not edit.

package streams

import (
	"github.com/gost-dom/browser/internal/promise"
	"github.com/gost-dom/browser/internal/streams"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w ReadableStreamDefaultReader[T]) CreateInstance(
	cbCtx js.CallbackContext[T], _x string,
) (js.Value[T], error) {
	return codec.EncodeCallbackErrorf(cbCtx,
		"gost-dom/scripting/streams: ReadableStreamDefaultReader constructor not yet supported",
	)
}

func (w ReadableStreamDefaultReader[T]) decodeReadableStream(ctx js.CallbackContext[T], v js.Value[T]) (string, error) {
	if v != nil && v.Boolean() {
		return "", codec.CallbackErrorf(ctx,
			"gost-dom/scripting/streams: ReadableStreamDefaultReader readableStream not yet supported",
		)
	}
	return "", nil
}
func (w ReadableStreamDefaultReader[T]) toPromiseReadableStreamReadResult(
	ctx js.CallbackContext[T], _ promise.Promise[streams.ReadResult]) (js.Value[T], error) {
	return codec.EncodeCallbackErrorf(ctx, "Encode ReadResult not implemented")
}
