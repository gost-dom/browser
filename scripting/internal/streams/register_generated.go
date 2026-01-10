// This file is generated. Do not edit.

package streams

import js "github.com/gost-dom/browser/scripting/internal/js"

func ConfigureWindowRealm[T any](e js.ScriptEngine[T]) {
	InitializeReadableStream(js.CreateClass(e, "ReadableStream", "", ReadableStreamConstructor))
	InitializeReadableStreamBYOBReader(js.CreateClass(e, "ReadableStreamBYOBReader", "", ReadableStreamBYOBReaderConstructor))
	InitializeReadableStreamDefaultReader(js.CreateClass(e, "ReadableStreamDefaultReader", "", ReadableStreamDefaultReaderConstructor))
}
