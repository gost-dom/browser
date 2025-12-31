package dom

import (
	"errors"
	"fmt"

	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type domException[T any] struct{}

func newDOMException[T any](_ js.ScriptEngine[T]) *domException[T] { return &domException[T]{} }

func domExceptionConstructor[T any](info js.CallbackContext[T]) (js.Value[T], error) {
	return nil, nil
}

func InitializeDomException[T any](jsClass js.Class[T]) {
	jsClass.CreateAttribute("code", domException_code, nil)
	jsClass.CreateAttribute("name", domException_name, nil)
	// TODO: Check if the message shouldn't have been Error.prototype.message
	jsClass.CreateAttribute("message", domException_message, nil)
}

// decodeThisError retrieves the wrapped error object from the `this` object
func decodeThisError[T any, E error](cbCtx js.CallbackContext[T], res *E) (err error) {
	var val error
	val, err = js.As[error](cbCtx.Instance())
	if err != nil {
		return
	}
	if !errors.As(val, res) {
		err = cbCtx.NewTypeError(fmt.Sprintf("Value does not implement target type: %T", *res))
		return
	}
	return
}

func domException_code[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var domErr dom.DOMError
	if err = decodeThisError(cbCtx, &domErr); err == nil {
		res = cbCtx.NewInt32(int32(domErr.Code))
	}
	return
}

func domException_name[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var domErr dom.DOMError
	if err = decodeThisError(cbCtx, &domErr); err == nil {
		res = cbCtx.NewString(domErr.Code.String())
	}
	return
}

func domException_message[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var domErr dom.DOMError
	if err = decodeThisError(cbCtx, &domErr); err == nil {
		res = cbCtx.NewString(domErr.Message)
	}
	return
}
