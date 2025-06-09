package scripting

import (
	"github.com/dave/jennifer/jen"
	g "github.com/gost-dom/generators"
)

// HostInitializer renders a method on the JavaScript mappings type that
// initializes a script host with the necessary methods and attributes on
// prototype and instance objects.
type HostInitializer struct {
	WrapperStruct
}

func (i HostInitializer) Generate() *jen.Statement {
	classArg := g.NewValue("jsClass")
	receiver := g.NewValue("wrapper")
	return g.StatementList(
		g.FunctionDefinition{
			Name:     "Initialize", // prototypeFactoryFunctionName(data),
			Receiver: g.FunctionArgument{Name: receiver, Type: i.WrapperStructType()},
			Args:     g.Arg(classArg, v8Class),
			Body:     i.body(receiver, classArg),
		}).Generate()
}

func (i HostInitializer) body(receiver g.Value, classArg g.Generator) g.Generator {
	return g.StatementList(
		receiver.Field("installPrototype").Call(classArg),
		renderIf(i.Data.RunCustomCode,
			receiver.Field("CustomInitializer").Call(classArg),
		),
	)
}
