package stdgen

import (
	"github.com/gost-dom/code-gen/packagenames"
	g "github.com/gost-dom/generators"
)

func LogDebug(args ...g.Generator) g.Generator {
	return g.NewValuePackage("Debug", packagenames.Log).Call(args...)
}

func ErrorsJoin(args ...g.Generator) g.Generator {
	return g.NewValuePackage("Join", "errors").Call(args...)
}

func ErrorsNew(arg g.Generator) g.Generator {
	return g.NewValuePackage("New", "errors").Call(arg)
}
