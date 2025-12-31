package streams

import (
	"github.com/gost-dom/browser/internal/promise"
	"github.com/gost-dom/browser/internal/streams"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func CreateReadableStreamDefaultReader[T any](
	cbCtx js.CallbackContext[T], _x string,
) (js.Value[T], error) {
	return codec.EncodeCallbackErrorf(cbCtx,
		"gost-dom/scripting/streams: ReadableStreamDefaultReader constructor not yet supported",
	)
}

func encodePromiseReadableStreamReadResult[T any](
	ctx js.CallbackContext[T], prom promise.Promise[streams.ReadResult]) (js.Value[T], error) {
	return codec.EncodePromise(ctx, prom, encodeReadResult)
}

func encodeReadResult[T any](
	s js.Scope[T], readResult streams.ReadResult,
) (js.Value[T], error) {
	res := s.NewObject()
	res.Set("value", s.NewUint8Array(readResult.Value))
	res.Set("done", s.NewBoolean(readResult.Done))
	return res, nil

}
