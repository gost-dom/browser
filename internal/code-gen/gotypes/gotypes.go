package gotypes

import (
	"github.com/gost-dom/code-gen/packagenames"
)

// GoType represents a go type referenced by attributes or operations - either
// as argument or return types. Name is the type name. If Package is the empty
// string, the type is assumed to belong to the package where the code is
// generated, or a native Go type, such as int or bool. Pointer is for struct
// types only, whether the type is a pointer type.
type GoType struct {
	Name    string
	Package string
	Pointer bool
}

func (t GoType) Zero() bool { return t == GoType{} }

var TimeDuration = GoType{
	Name:    "Duration",
	Package: "time",
}
var TaskHandle = GoType{
	Name:    "TaskHandle",
	Package: packagenames.Clock,
}
