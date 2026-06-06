package cssomview1

import js "github.com/gost-dom/browser/scripting/internal/js"

func Element_getBoundingClientRect[T any](
	cbCtx js.CallbackContext[T],
) (res js.Value[T], err error) {
	return createDummyBoundingRect(cbCtx), nil
}

func Element_getClientRects[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return createDummyClientRects(cbCtx), nil
}
