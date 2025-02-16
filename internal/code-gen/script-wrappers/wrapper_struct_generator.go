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
		g.Platform.CreateWrapperStruct(g.Data),
	).Generate()
}

func (g WrapperStructGenerator) TypeGenerator() Generator {
	idlInterfaceName := g.Data.Name()

	structGens := g.Platform.WrapperStructGenerators()
	includes := g.Data.Includes()
	innerType := gen.NewTypePackage(idlInterfaceName, g.Data.GetInternalPackage())
	embedName := structGens.EmbedName(g.Data)
	wrapperStruct := gen.NewStruct(gen.Id(structGens.WrapperStructTypeName(idlInterfaceName)))
	wrapperStruct.Embed(gen.NewType(embedName).TypeParam(innerType))

	for _, i := range includes {
		wrapperStruct.Field(
			gen.Id(lowerCaseFirstLetter(i.Name)),
			gen.NewType(structGens.WrapperStructTypeName(i.Name)).Pointer(),
		)
	}
	return wrapperStruct
}
