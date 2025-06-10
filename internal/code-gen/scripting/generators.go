package scripting

import (
	"github.com/dave/jennifer/jen"
	g "github.com/gost-dom/generators"
)

type Transformer interface {
	Transform(g.Generator) g.Generator
}

type TransformerFunc func(g.Generator) g.Generator

func (f TransformerFunc) Transform(gen g.Generator) g.Generator { return f(gen) }

func returnNilCommaErr(err g.Generator) g.Generator {
	return g.Return(g.Nil, err)
}

func renderIf(condition bool, gen g.Generator) g.Generator {
	if condition {
		return gen
	}
	return g.Noop
}

func renderIfElse(condition bool, gen g.Generator, elseGen g.Generator) g.Generator {
	if condition {
		return gen
	} else {
		return elseGen
	}
}

func addLinesBetweenElements(gs []g.Generator) []g.Generator {
	l := len(gs)
	if l <= 1 {
		return gs
	}
	for i, gg := range gs {
		gs[i] = g.Raw(jen.Line().Add(gg.Generate()))
	}
	gs = append(gs, g.Line)
	return gs
}
