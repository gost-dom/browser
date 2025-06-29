package streams

import (
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w ReadableStream[T]) decodeObject(
	cbCtx js.CallbackContext[T], v js.Value[T],
) (string, error) {
	var err error
	if v != nil && v.Boolean() {
		return "", codec.CallbackErrorf(cbCtx,
			"godt-dom/scripting/streams: ReadableStream underlyingSource not yet supported",
		)
	}
	return "", err
}

func (w ReadableStream[T]) decodeQueuingStrategy(
	cbCtx js.CallbackContext[T], v js.Value[T],
) ([]string, error) {
	if v != nil && v.Boolean() {
		return nil, codec.CallbackErrorf(cbCtx,
			"gost-dom/scripting/streams: ReadableStream strategy not yet supported",
		)
	}
	return nil, nil
}

func (w ReadableStream[T]) CreateInstance(
	cbCtx js.CallbackContext[T], _x string, _y ...string,
) (js.Value[T], error) {
	return codec.EncodeCallbackErrorf(cbCtx,
		"gost-dom/scripting/streams: ReadableStream constructor not yet supported",
	)
}
