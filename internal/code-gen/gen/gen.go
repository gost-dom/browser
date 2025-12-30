package gen

import g "github.com/gost-dom/generators"

func NewlineBefore(gen g.Generator) g.Generator {
	return g.Raw(g.Line.Generate().Add(gen.Generate()))
}
