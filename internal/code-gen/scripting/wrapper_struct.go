package scripting

import (
	"github.com/dave/jennifer/jen"
	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/scripting/model"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
	gen "github.com/gost-dom/generators"
)

var scriptHost = g.NewValue("scriptHost")

type WrapperStruct struct {
	Data model.ESConstructorData
}

// TypeDefinition renders the actual struct type definition
func (ws WrapperStruct) TypeDefinition() g.Generator {
	return g.StatementList(
		ws.TypeGenerator(),
		ws.ConstructorGenerator(),
	)
}

func (ws WrapperStruct) IdlName() string { return ws.Data.Name() }

func (g WrapperStruct) WrapperStructTypeForName(name string) generators.Type {
	return generators.Type{
		Generator: generators.Raw(jen.Id(name).Types(jen.Id("T"))),
	}
}

func (g WrapperStruct) WrapperStructType() generators.Type {
	return g.WrapperStructTypeForName(g.IdlName())
}

func (g WrapperStruct) TypeGenerator() g.Generator {
	includes := g.Data.Includes()
	wrapperStruct := gen.NewStruct(
		generators.Raw(jen.Id(g.IdlName()).Types(jen.Id("T").Any())),
	)

	for _, i := range includes {
		wrapperStruct.Field(
			gen.Id(LowerCaseFirstLetter(i.Name)),
			generators.Raw(jen.Op("*").Add(g.WrapperStructTypeForName(i.Name).Generate())),
		)
	}
	return wrapperStruct
}

func (wrapper WrapperStruct) ConstructorGenerator() g.Generator {
	idlInterfaceName := wrapper.Data.Name()
	constructorName := ConstructorNameForInterface(idlInterfaceName)
	hostArg := g.Id("scriptHost")

	return gen.Raw(
		jen.Func().Id(constructorName).
			Types(jen.Id("T").Any()).
			Params(
				hostArg.Generate().Add(jsScriptEngine.Generate()),
			).
			Params(
				wrapper.WrapperStructType().Pointer().Generate(),
			).
			Block(wrapper.Body().Generate()))
}

func (ws WrapperStruct) WrapperStructTypeRetDef() g.Type {
	return generators.Type{
		Generator: generators.Raw(jen.Id(ws.IdlName()).Types(jen.Id("T"))),
	}
}

func (ws WrapperStruct) Body() g.Generator {
	includes := ws.Data.Includes()
	fieldInitializers := make([]g.Generator, 0)
	for _, i := range includes {
		includeConstructorName := ConstructorNameForInterface(i.Name)
		fieldInitializers = append(
			fieldInitializers,
			generators.NewValue(includeConstructorName).Call(scriptHost),
		)
	}
	fieldInitializers = addLinesBetweenElements(fieldInitializers)

	wrapperType := ws.WrapperStructTypeRetDef()
	return generators.Return(wrapperType.CreateInstance(fieldInitializers...).Reference())
}

func (ws WrapperStruct) PlatformInfoArg() g.Generator { return generators.Id("info") }

func (ws WrapperStruct) Callbacks() CallbackMethods {
	return CallbackMethods{ws}
}
