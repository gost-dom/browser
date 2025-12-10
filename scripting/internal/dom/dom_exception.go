package dom

import (
	"errors"

	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type domException[T any] struct{}

func newDOMException[T any](_ js.ScriptEngine[T]) *domException[T] { return &domException[T]{} }

func (w domException[T]) Constructor(info js.CallbackContext[T]) (js.Value[T], error) {
	return nil, nil
}

func (w domException[T]) Initialize(jsClass js.Class[T]) {
	jsClass.CreateInstanceAttribute("code", w.code, nil)
	jsClass.CreateInstanceAttribute("name", w.name, nil)
	// TODO: Check if the message shouldn't have been Error.prototype.message
	jsClass.CreateInstanceAttribute("message", w.message, nil)
}

// decodeThisError retrieves the wrapped error object from the `this` object
func decodeThisError[T any, E error](cbCtx js.CallbackContext[T], res *E) (err error) {
	var val error
	val, err = js.As[error](cbCtx.Instance())
	if err != nil {
		return
	}
	if !errors.As(val, res) {
		err = cbCtx.NewTypeError("Value does not implement DOMException")
		return
	}
	return
}

func (w domException[T]) code(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var domErr dom.DOMError
	if err = decodeThisError(cbCtx, &domErr); err == nil {
		res = cbCtx.NewInt32(int32(domErr.Code))
	}
	return
}

func (w domException[T]) name(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var domErr dom.DOMError
	if err = decodeThisError(cbCtx, &domErr); err == nil {
		res = cbCtx.NewString(domErr.Code.String())
	}
	return
}

func (w domException[T]) message(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var domErr dom.DOMError
	if err = decodeThisError(cbCtx, &domErr); err == nil {
		res = cbCtx.NewString(domErr.Message)
	}
	return
}
