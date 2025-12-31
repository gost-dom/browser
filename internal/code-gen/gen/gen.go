package gen

import (
	"github.com/dave/jennifer/jen"
	g "github.com/gost-dom/generators"
)

func NewlineBefore(gen g.Generator) g.Generator {
	return g.Raw(g.Line.Generate().Add(gen.Generate()))
}

func Not(expr g.Generator) g.Generator {
	return g.Raw(jen.Op("!").Add(expr.Generate()))
}

func AnyConstraint(t g.Generator) g.Generator {
	return g.Raw(t.Generate().Any())
}
