package wrappers

import (
	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/generators"
	gen "github.com/gost-dom/generators"
)

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
		g.BodyGenerator(),
	).Generate()
}

func (g PrototypeInitializerGenerator) BodyGenerator() generators.Generator {
	return generators.StatementList(
		g.Platform.CreatePrototypeInitializerBody(g.Data),
		g.MixinsGenerator(),
	)
}

func (g PrototypeInitializerGenerator) MixinsGenerator() generators.Generator {
	result := gen.StatementList()
	for _, mixin := range g.Data.IdlInterface.Includes {
		wrapperName := lowerCaseFirstLetter(mixin.Name)
		// Note: This excludes mixins not known in this spec.
		// Could the mixing come from another spec, then this will skip
		// something we may want.
		// Not a problem yet ...
		if _, included := g.Data.Spec.DomSpec.Interfaces[mixin.Name]; included {
			result.Append(
				gen.NewValue("w").Field(wrapperName).Field("installPrototype").Call(
					gen.Id("prototypeTmpl")),
			)
		}
	}
	return result
}
