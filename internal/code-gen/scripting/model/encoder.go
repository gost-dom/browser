package model

import (
	"fmt"

	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

// Given only an IDL type specification, evaluate a function that can _encode_
// this value.
//
// Without any knowledge of the Go type, this will have to make a qualified
// guess, e.g., an IDL DOMString is represented as Go string. If the guess is
// wrong, the generated code will fail to compile.
func EncoderForIdlType(t idl.Type) g.Value {
	// Note, this is _currently_ only used for iterator protocol implementation,
	// but a lot of duplication exists in this library
	switch t.Name {
	case "DOMString", "USVString":
		return encodeString
	case "ByteString":
		return encodeByteString
	}
	return g.NewValue(fmt.Sprintf("encode%s", t.Name))
}
