package typerule

import (
	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/generators"
)

// TypeRule overrides a default type, e.g., return value, attribute type, or
// argument value
type TypeRule struct {
	Name    string // Type name in Go
	Package string // Package. If empty, same package as output
	Pointer bool   // Whether to generate pointer types for struct types
}

var Bool = &TypeRule{Name: "bool"}

func (r TypeRule) Generate() *jen.Statement {
	var res generators.Type
	if r.Package == "" {
		res = generators.NewType(r.Name)
	} else {
		res = generators.NewTypePackage(r.Name, r.Package)
	}
	if r.Pointer {
		return res.Pointer().Generate()
	}
	return res.Generate()
}
