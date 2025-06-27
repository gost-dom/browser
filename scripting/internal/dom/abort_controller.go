package dom

import (
	"github.com/gost-dom/browser/internal/dom"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type abortControllerWrapper struct {
	*dom.AbortController
}

func (w abortControllerWrapper) Signal() dominterfaces.AbortSignal {
	return w.AbortController.Signal()
}

func (w AbortController[T]) CreateInstance(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	ctrl := dominterfaces.AbortController(abortControllerWrapper{dom.NewAbortController()})
	return codec.EncodeConstrucedValue(cbCtx, ctrl)
}

func (w AbortController[T]) toAbortSignal(
	cbCtx js.CallbackContext[T],
	signal dominterfaces.AbortSignal,
) (js.Value[T], error) {
	return cbCtx.Constructor("AbortSignal").NewInstance(signal)
}
