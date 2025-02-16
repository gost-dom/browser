package wrappers

import "github.com/dave/jennifer/jen"

// PrototypeInitializerGenerator generates a function to initialize a JavaScript
// prototype. Initialize means setting data attributes for functions, and
// accessor attributes for IDL attributes.
type PrototypeInitializerGenerator struct {
	Platform TargetGenerators
	Data     ESConstructorData
}

func (g PrototypeInitializerGenerator) Generate() *jen.Statement {
	return g.Platform.CreatePrototypeInitializer(
		g.Data,
		g.Platform.CreatePrototypeInitializerBody(g.Data),
	).Generate()

}
