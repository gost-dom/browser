package wrappers

import (
	"github.com/dave/jennifer/jen"
	. "github.com/gost-dom/code-gen/internal"
	. "github.com/gost-dom/code-gen/script-wrappers/model"
	"github.com/gost-dom/generators"
	gen "github.com/gost-dom/generators"
)

var scriptHost = gen.NewValue("scriptHost")

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

func (g WrapperStructGenerator) TypeGenerator() Generator {
	structGens := g.Platform.WrapperStructGenerators()

	idlInterfaceName := g.Data.Name()
	includes := g.Data.Includes()
	wrapperStruct := gen.NewStruct(structGens.WrapperStructType(idlInterfaceName))
	wrapperStruct.Embed(structGens.EmbeddedType(g.Data.WrappedType()))

	for _, i := range includes {
		wrapperStruct.Field(
			gen.Id(LowerCaseFirstLetter(i.Name)),
			structGens.WrapperStructType(i.Name).Pointer(),
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
		Body:     g.Body(),
	}
}

func addLinesBetweenElements(g []Generator) []Generator {
	l := len(g)
	if l <= 1 {
		return g
	}
	for i, gg := range g {
		g[i] = generators.Raw(jen.Line().Add(gg.Generate()))
	}
	g = append(g, generators.Line)
	return g
}

func (g WrapperStructGenerator) Body() Generator {
	structGens := g.Platform.WrapperStructGenerators()
	idlInterfaceName := g.Data.Name()

	innerType := g.Data.WrappedType()
	embedConstructorName := structGens.EmbeddedTypeConstructor(innerType)
	includes := g.Data.Includes()
	fieldInitializers := make([]Generator, len(includes)+1)
	for idx, i := range includes {
		includeConstructorName := structGens.WrapperStructConstructorName(i.Name)
		fieldInitializers[idx+1] = generators.NewValue(includeConstructorName).Call(scriptHost)
	}
	fieldInitializers[0] = embedConstructorName.Call(structGens.HostArg())
	fieldInitializers = addLinesBetweenElements(fieldInitializers)

	wrapperType := structGens.WrapperStructType(idlInterfaceName)
	return generators.Return(wrapperType.CreateInstance(fieldInitializers...).Reference())
}

func (g WrapperStructGenerator) PlatformInfoArg() Generator { return generators.Id("info") }
