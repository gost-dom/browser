package wrappers

import (
	"github.com/dave/jennifer/jen"
	gen "github.com/gost-dom/generators"
)

type WrapperStructGenerator struct {
	Platform TargetGenerators
	Data     ESConstructorData
}

func (g WrapperStructGenerator) Generate() *jen.Statement {
	return gen.StatementList(
		g.TypeGenerator(),
		g.ConstructorGenerator(),
		gen.Line,
	).Generate()
}

func (g WrapperStructGenerator) WrappedType() Generator {
	idlInterfaceName := g.Data.Name()
	return gen.NewTypePackage(idlInterfaceName, g.Data.GetInternalPackage())
}

func (g WrapperStructGenerator) TypeGenerator() Generator {
	structGens := g.Platform.WrapperStructGenerators()

	idlInterfaceName := g.Data.Name()
	includes := g.Data.Includes()
	wrapperStruct := gen.NewStruct(structGens.WrapperStructType(idlInterfaceName))
	wrapperStruct.Embed(structGens.EmbeddedType(g.WrappedType()))

	for _, i := range includes {
		wrapperStruct.Field(
			gen.Id(lowerCaseFirstLetter(i.Name)),
			gen.Type{Generator: structGens.WrapperStructType(i.Name)}.Pointer(),
		)
	}
	return wrapperStruct
}

func (g WrapperStructGenerator) ConstructorGenerator() Generator {
	structGens := g.Platform.WrapperStructGenerators()

	idlInterfaceName := g.Data.Name()
	constructorName := structGens.WrapperStructConstructorName(idlInterfaceName)
	hostArg := structGens.HostArg()
	hostType := structGens.HostType()

	return gen.FunctionDefinition{
		Name:     constructorName,
		Args:     gen.Arg(hostArg, hostType),
		RtnTypes: gen.List(structGens.WrapperStructConstructorRetType(idlInterfaceName)),
		Body:     gen.Return(g.Platform.CreateWrapperStruct(g.Data)),
	}
}
