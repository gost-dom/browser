package scripting

import g "github.com/gost-dom/generators"

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
