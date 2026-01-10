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

func Panic(expr g.Generator) g.Generator {
	return g.Raw(jen.Panic(expr.Generate()))
}

// Renders an type parameter with name n without a type constraint, i.e., an
// "any" constraint. If n is nil, the name will be T.
func AnyConstraint(t g.Generator) g.Generator {
	if t == nil {
		t = g.Id("T")
	}
	return g.Raw(t.Generate().Any())
}
