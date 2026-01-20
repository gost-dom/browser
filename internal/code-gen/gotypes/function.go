package gotypes

import "github.com/dave/jennifer/jen"

type GoFunction struct {
	Name    string
	Package string
}

func (f GoFunction) Generate() *jen.Statement {
	if f.Package == "" {
		return jen.Id(f.Name)
	}
	return jen.Qual(f.Package, f.Name)
}

func (f GoFunction) IsZero() bool { return f == GoFunction{} }
