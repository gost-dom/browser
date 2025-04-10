package stdgen

import (
	"github.com/gost-dom/code-gen/packagenames"
	g "github.com/gost-dom/generators"
)

func LogDebug(args ...g.Generator) g.Generator {
	return g.NewValuePackage("Debug", packagenames.Log).Call(args...)
}
