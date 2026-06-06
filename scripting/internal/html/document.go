package html

import js "github.com/gost-dom/browser/scripting/internal/js"

func Document_defaultView[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	// TODO: Implement a WindowProxy class. But for now, this will do
	return cbCtx.GlobalThis(), nil
}
