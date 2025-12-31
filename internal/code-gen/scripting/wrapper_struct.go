package scripting

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/scripting/model"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
)

var scriptHost = g.NewValue("scriptHost")

type WrapperStruct struct {
	Data model.ESConstructorData
}

func (ge WrapperStruct) InitializerName() string {
	return fmt.Sprintf("Initialize%s", ge.IdlName())
}

func (ge WrapperStruct) Initializer() g.Generator {
	return g.Id(ge.InitializerName())
}

func (ge WrapperStruct) SpecName() string { return ge.Data.SpecName() }

// TypeDefinition renders the actual struct type definition
func (ge WrapperStruct) TypeDefinition() g.Generator {
	return g.StatementList(
		ge.TypeGenerator(),
		ge.ConstructorGenerator(),
	)
}

func (ge WrapperStruct) IdlName() string { return ge.Data.Name() }

func (ge WrapperStruct) WrapperStructTypeForName(name string) generators.Type {
	return generators.Type{
		Generator: generators.Raw(jen.Id(name).Types(jen.Id("T"))),
	}
}

func (ge WrapperStruct) WrapperStructType() generators.Type {
	return ge.WrapperStructTypeForName(ge.IdlName())
}

func (ge WrapperStruct) TypeGenerator() g.Generator {
	includes := ge.Data.Includes()
	wrapperStruct := g.NewStruct(
		generators.Raw(jen.Id(ge.IdlName()).Types(jen.Id("T").Any())),
	)

	for _, i := range includes {
		wrapperStruct.Field(
			g.Id(LowerCaseFirstLetter(i.Name)),
			generators.Raw(ge.WrapperStructTypeForName(i.Name).Generate()),
		)
	}
	return wrapperStruct
}

func Initializer(d model.ESConstructorData) g.Generator {
	ws := WrapperStruct{d}
	return ws.Initializer()
}

func (ge WrapperStruct) ConstructorGenerator() g.Generator {
	idlInterfaceName := ge.Data.Name()
	constructorName := ConstructorNameForInterface(idlInterfaceName)
	hostArg := g.Id("scriptHost")

	return g.Raw(
		jen.Func().Id(constructorName).
			Types(jen.Id("T").Any()).
			Params(
				hostArg.Generate().Add(jsScriptEngine.Generate()),
			).
			Params(
				ge.WrapperStructType().Generate(),
			).
			Block(ge.Body().Generate()))
}

func (ge WrapperStruct) WrapperStructTypeRetDef() g.Type {
	return generators.Type{
		Generator: generators.Raw(jen.Id(ge.IdlName()).Types(jen.Id("T"))),
	}
}

func (ge WrapperStruct) Body() g.Generator {
	includes := ge.Data.Includes()
	fieldInitializers := make([]g.Generator, 0)
	for _, i := range includes {
		includeConstructorName := ConstructorNameForInterface(i.Name)
		fieldInitializers = append(
			fieldInitializers,
			generators.NewValue(includeConstructorName).Call(scriptHost),
		)
	}
	fieldInitializers = addLinesBetweenElements(fieldInitializers)

	wrapperType := ge.WrapperStructTypeRetDef()
	return generators.Return(wrapperType.CreateInstance(fieldInitializers...))
}

func (ge WrapperStruct) PlatformInfoArg() g.Generator { return generators.Id("info") }

func (ge WrapperStruct) Callbacks() CallbackMethods {
	return CallbackMethods{ge}
}
