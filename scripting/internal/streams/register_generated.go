// This file is generated. Do not edit.

package streams

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](e js.ScriptEngine[T]) {
	js.RegisterClass(e, "ReadableStream", "", NewReadableStream, ReadableStreamConstructor)
	js.RegisterClass(e, "ReadableStreamBYOBReader", "", NewReadableStreamBYOBReader, ReadableStreamBYOBReaderConstructor)
	js.RegisterClass(e, "ReadableStreamDefaultReader", "", NewReadableStreamDefaultReader, ReadableStreamDefaultReaderConstructor)
}
