package scripting

import (
	"github.com/dave/jennifer/jen"
	. "github.com/gost-dom/code-gen/internal"
	. "github.com/gost-dom/code-gen/scripting/model"
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
	wrapperStruct := gen.NewStruct(structGens.WrapperStructTypeDef(idlInterfaceName))
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

	if g.Platform.Name() == "goja" {
		return gen.FunctionDefinition{
			Name:     constructorName,
			Args:     gen.Arg(hostArg, hostType),
			RtnTypes: gen.List(structGens.WrapperStructConstructorRetType(idlInterfaceName)),
			Body:     g.Body(),
		}
	} else {

		return gen.Raw(
			jen.Func().Id(constructorName).
				Types(jen.Id("T").Any()).
				Params(
					hostArg.Generate().Add(hostType.Generate()),
				).
				Params(
					structGens.WrapperStructConstructorRetType(idlInterfaceName).Generate(),
				).
				Block(g.Body().Generate()))
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
	fieldInitializers := make([]Generator, 0)
	if embedConstructorName != nil {
		fieldInitializers = append(
			fieldInitializers,
			generators.ValueOf(embedConstructorName).Call(structGens.HostArg()),
		)
	}
	for _, i := range includes {
		includeConstructorName := structGens.WrapperStructConstructorName(i.Name)
		fieldInitializers = append(
			fieldInitializers,
			generators.NewValue(includeConstructorName).Call(scriptHost),
		)
	}
	fieldInitializers = addLinesBetweenElements(fieldInitializers)

	wrapperType := structGens.WrapperStructTypeRetDef(idlInterfaceName)
	return generators.Return(wrapperType.CreateInstance(fieldInitializers...).Reference())
}

func (g WrapperStructGenerator) PlatformInfoArg() Generator { return generators.Id("info") }
