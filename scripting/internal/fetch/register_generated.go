// This file is generated. Do not edit.

package fetch

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "Headers", "", Headers[T]{}.Initialize, HeadersConstructor)
	js.RegisterClass(e, "Request", "", Request[T]{}.Initialize, RequestConstructor)
	js.RegisterClass(e, "Response", "", Response[T]{}.Initialize, ResponseConstructor)
}
