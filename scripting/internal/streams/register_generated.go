// This file is generated. Do not edit.

package streams

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "ReadableStream", "", InitializeReadableStream, ReadableStreamConstructor)
	js.RegisterClass(e, "ReadableStreamBYOBReader", "", InitializeReadableStreamBYOBReader, ReadableStreamBYOBReaderConstructor)
	js.RegisterClass(e, "ReadableStreamDefaultReader", "", InitializeReadableStreamDefaultReader, ReadableStreamDefaultReaderConstructor)
}
