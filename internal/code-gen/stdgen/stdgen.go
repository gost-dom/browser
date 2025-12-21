package stdgen

import (
	g "github.com/gost-dom/generators"
)

func ErrorsJoin(args ...g.Generator) g.Generator {
	return g.NewValuePackage("Join", "errors").Call(args...)
}

func ErrorsNew(arg g.Generator) g.Generator {
	return g.NewValuePackage("New", "errors").Call(arg)
}
