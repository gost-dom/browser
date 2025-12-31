package variable

import (
	"github.com/dave/jennifer/jen"
	g "github.com/gost-dom/generators"
)

type Var struct {
	Name g.Generator
	Type g.Generator
}

func New(opts ...VarOption) Var {
	var res Var
	for _, opt := range opts {
		opt(&res)
	}
	return res
}

func (v Var) Generate() *jen.Statement {
	res := jen.Var()
	if v.Name != nil {
		res.Add(v.Name.Generate())
	}
	if v.Type != nil {
		res.Add(v.Type.Generate())
	}
	return res
}

type VarOption func(*Var)

func NameStr(n string) VarOption {
	return func(v *Var) { v.Name = g.Id(n) }
}

func Name(n g.Generator) VarOption {
	return func(v *Var) { v.Name = n }
}
func Type(t g.Generator) VarOption {
	return func(v *Var) { v.Type = t }
}
