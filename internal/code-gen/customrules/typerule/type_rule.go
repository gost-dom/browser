package typerule

import "github.com/dave/jennifer/jen"

// TypeRule overrides a default type, e.g., return value, attribute type, or
// argument value
type TypeRule struct {
	Name string
}

var Bool = &TypeRule{Name: "bool"}

func (r TypeRule) Generate() *jen.Statement {
	return jen.Id(r.Name)
}
