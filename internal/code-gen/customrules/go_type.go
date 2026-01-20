package customrules

import (
	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/gotypes"
	g "github.com/gost-dom/generators"
)

type GoType = gotypes.GoType
type GoFunction = gotypes.GoFunction

type GoTypeGenerator GoType

func (t GoTypeGenerator) Generate() *jen.Statement {
	var res g.Type
	if t.Package == "" {
		res = g.NewType(t.Name)
	} else {
		res = g.NewTypePackage(t.Name, t.Package)
	}
	if t.Pointer {
		res = res.Pointer()
	}
	return res.Generate()
}
