package typerule

import "github.com/dave/jennifer/jen"

// TypeRule overrides a default type, e.g., return value, attribute type, or
// argument value
type TypeRule struct {
	Name    string // Type name in Go
	Package string // Package. If empty, same package as output
	Pointer bool   // Whether to generate pointer types for struct types
}

var Bool = &TypeRule{Name: "bool"}

func (r TypeRule) Generate() *jen.Statement {
	return jen.Id(r.Name)
}
