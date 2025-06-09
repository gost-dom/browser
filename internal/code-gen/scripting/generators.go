package scripting

import g "github.com/gost-dom/generators"

type Transformer interface {
	Transform(g.Generator) g.Generator
}

type TransformerFunc func(g.Generator) g.Generator

func (f TransformerFunc) Transform(gen g.Generator) g.Generator { return f(gen) }
