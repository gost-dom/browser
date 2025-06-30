// This file is generated. Do not edit.

package streams

import js "github.com/gost-dom/browser/scripting/internal/js"

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "ReadableStream", "", NewReadableStream)
	js.RegisterClass(reg, "ReadableStreamBYOBReader", "", NewReadableStreamBYOBReader)
	js.RegisterClass(reg, "ReadableStreamDefaultReader", "", NewReadableStreamDefaultReader)
}
