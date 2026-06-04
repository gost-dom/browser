package html

import (
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func CreateMessageChannel[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return nil, nil
}

func encodeMessagePort[T any](
	cbCtx js.CallbackContext[T],
	ch htmlinterfaces.MessagePort,
) (js.Value[T], error) {
	return codec.EncodeEntityScopedWithPrototype(cbCtx, ch, "MessagePort")
}
