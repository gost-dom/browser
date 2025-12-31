// This file is generated. Do not edit.

package streams

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "ReadableStream", "", ReadableStream[T]{}.Initialize, ReadableStreamConstructor)
	js.RegisterClass(e, "ReadableStreamBYOBReader", "", ReadableStreamBYOBReader[T]{}.Initialize, ReadableStreamBYOBReaderConstructor)
	js.RegisterClass(e, "ReadableStreamDefaultReader", "", ReadableStreamDefaultReader[T]{}.Initialize, ReadableStreamDefaultReaderConstructor)
}
