package scripting

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/scripting/model"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
	gen "github.com/gost-dom/generators"
)

var scriptHost = gen.NewValue("scriptHost")

type WrapperStruct struct {
	Data model.ESConstructorData
}

func (s WrapperStruct) WrapperStructConstructorName(name string) string {
	return fmt.Sprintf("New%s", s.typeNameForType(name))
}

func (g WrapperStruct) TypeDef() gen.Generator {
	return gen.StatementList(
		g.TypeGenerator(),
		g.ConstructorGenerator(),
	)
}

func (g WrapperStruct) IdlName() string { return g.Data.Name() }

func (gen WrapperStruct) typeNameForType(name string) string {
	return fmt.Sprintf("%sV8Wrapper", name)
}

func (gen WrapperStruct) generatedTypeName() string {
	return gen.typeNameForType(gen.IdlName())
}

func (g WrapperStruct) WrapperStructTypeForName(name string) generators.Type {
	return generators.Type{
		Generator: generators.Raw(jen.Id(g.typeNameForType(name)).Types(jen.Id("T"))),
	}
}

func (g WrapperStruct) WrapperStructType() generators.Type {
	return g.WrapperStructTypeForName(g.IdlName())
}

func (g WrapperStruct) TypeGenerator() Generator {
	includes := g.Data.Includes()
	wrapperStruct := gen.NewStruct(
		generators.Raw(jen.Id(g.generatedTypeName()).Types(jen.Id("T").Any())),
	)

	for _, i := range includes {
		wrapperStruct.Field(
			gen.Id(LowerCaseFirstLetter(i.Name)),
			generators.Raw(jen.Op("*").Add(g.WrapperStructTypeForName(i.Name).Generate())),
		)
	}
	return wrapperStruct
}

func (wrapper WrapperStruct) ConstructorGenerator() Generator {
	idlInterfaceName := wrapper.Data.Name()
	constructorName := wrapper.WrapperStructConstructorName(idlInterfaceName)
	hostArg := g.Id("scriptHost")

	return gen.Raw(
		jen.Func().Id(constructorName).
			Types(jen.Id("T").Any()).
			Params(
				hostArg.Generate().Add(scriptEngine.Generate()),
			).
			Params(
				wrapper.WrapperStructType().Pointer().Generate(),
			).
			Block(wrapper.Body().Generate()))
}

func (s WrapperStruct) WrapperStructTypeRetDef() g.Type {
	return generators.Type{
		Generator: generators.Raw(jen.Id(s.generatedTypeName()).Types(jen.Id("T"))),
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

func (g WrapperStruct) Body() Generator {
	includes := g.Data.Includes()
	fieldInitializers := make([]Generator, 0)
	for _, i := range includes {
		includeConstructorName := g.WrapperStructConstructorName(i.Name)
		fieldInitializers = append(
			fieldInitializers,
			generators.NewValue(includeConstructorName).Call(scriptHost),
		)
	}
	fieldInitializers = addLinesBetweenElements(fieldInitializers)

	wrapperType := g.WrapperStructTypeRetDef()
	return generators.Return(wrapperType.CreateInstance(fieldInitializers...).Reference())
}

func (g WrapperStruct) PlatformInfoArg() Generator { return generators.Id("info") }

func (s WrapperStruct) Callbacks() CallbackGenerators {
	return CallbackGenerators{s}
}
