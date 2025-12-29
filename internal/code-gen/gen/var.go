package gen

import (
	"github.com/dave/jennifer/jen"
	g "github.com/gost-dom/generators"
)

func Var(id g.Generator) g.Generator {
	return g.Raw(jen.Var())
}
