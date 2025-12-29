package model

import (
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/webref/idl"
)

func IdlNameToGoName(s string) string {
	return internal.IdlNameToGoName(s)
}

func IsDefined(t idl.Type) bool { return t.Name != "undefined" }
