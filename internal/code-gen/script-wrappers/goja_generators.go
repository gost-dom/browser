package wrappers

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
)

var (
	gojaFc      = g.Raw(jen.Qual(gojaSrc, "FunctionCall"))
	gojaValue   = g.Raw(jen.Qual(gojaSrc, "Value"))
	gojaObj     = g.Raw(jen.Op("*").Qual(gojaSrc, "Object"))
	gojaRuntime = g.Raw(jen.Op("*").Qual(gojaSrc, "Runtime"))
	flagTrue    = g.Raw(jen.Qual(gojaSrc, "FLAG_TRUE"))
)

type GojaNamingStrategy struct {
	ESConstructorData
}

func (s GojaNamingStrategy) PrototypeWrapperBaseName() string {
	return fmt.Sprintf("%sWrapper", s.Name())
}

func (s GojaNamingStrategy) PrototypeWrapperTypeName() string {
	return LowerCaseFirstLetter(s.PrototypeWrapperBaseName())
}

func (s GojaNamingStrategy) PrototypeWrapperConstructorName() string {
	return fmt.Sprintf("new%s", s.PrototypeWrapperBaseName())
}

func (s GojaNamingStrategy) ReceiverName() string {
	return "w" // data.Receiver
}

type GojaTargetGenerators struct{}

func (gen GojaTargetGenerators) PlatformInfoArg() g.Generator { return g.Id("c") }

// CreateConstructor has no effect for Goja. It's currently based on a system
// that it automatically creates the constructors based on whether or not they
// call the ingerface
func (gen GojaTargetGenerators) CreateHostInitializer(ESConstructorData) g.Generator {
	return g.Noop
}

func (gen GojaTargetGenerators) CreateConstructorCallback(ESConstructorData) g.Generator {
	return g.Noop
}

func (gen GojaTargetGenerators) CreateInitFunction(data ESConstructorData) g.Generator {
	naming := GojaNamingStrategy{data}
	return g.FunctionDefinition{
		Name: "init",
		Body: g.NewValue("installClass").
			Call(
				g.Lit(data.Name()),
				g.Lit(data.Inheritance),
				g.Id(naming.PrototypeWrapperConstructorName()),
			),
	}
}

// CreatePrototypeInitializer creates the "initializePrototype" method, which
// sets all the properties on the prototypes on this class.
func (gen GojaTargetGenerators) CreatePrototypeInitializer(
	data ESConstructorData,
	body g.Generator,
) g.Generator {
	naming := GojaNamingStrategy{data}
	receiver := g.NewValue(naming.ReceiverName())
	prototype := g.NewValue("prototype")

	return g.FunctionDefinition{
		Receiver: g.FunctionArgument{
			Name: receiver,
			Type: g.Id(naming.PrototypeWrapperTypeName()),
		},
		Name: "initializePrototype",
		Args: g.Arg(prototype, gojaObj).Arg(g.Id("vm"), gojaRuntime),
		Body: body,
	}
}

func (gen GojaTargetGenerators) CreatePrototypeInitializerBody(
	data ESConstructorData,
) g.Generator {
	naming := GojaNamingStrategy{data}
	receiver := g.NewValue(naming.ReceiverName())
	vm := receiver.Field("ctx").Field("vm")
	prototype := g.NewValue("prototype")
	body := g.StatementList()
	for op := range data.WrapperFunctionsToInstall() {
		body.Append(prototype.Field("Set").Call(g.Lit(op.Name), receiver.Field(op.Name)))
	}

	for a := range data.AttributesToInstall() {
		var getter, setter g.Generator
		if a.Getter != nil {
			getter = vm.Field("ToValue").Call(receiver.Field(a.Getter.Name))
		} else {
			getter = g.Nil
		}
		if a.Setter != nil {
			setter = vm.Field("ToValue").Call(receiver.Field(a.Setter.Name))
		} else {
			setter = g.Nil
		}
		body.Append(
			prototype.Field("DefineAccessorProperty").
				Call(g.Lit(a.Name), getter, setter, flagTrue, flagTrue),
		)
	}

	return body
}

func (gen GojaTargetGenerators) ReturnError(errGen g.Generator) g.Generator {
	return g.Raw(jen.Panic(errGen.Generate()))
}

func (gen GojaTargetGenerators) CreateMethodCallbackBody(
	data ESConstructorData,
	op ESOperation,
) g.Generator {
	callArgument := g.Id("c")
	if op.NotImplemented {
		msg := fmt.Sprintf(
			"%s.%s: Not implemented. Create an issue: %s",
			data.Name(),
			op.Name,
			packagenames.ISSUE_URL,
		)
		return g.Raw(jen.Panic(jen.Lit(msg)))
	}
	naming := GojaNamingStrategy{data}
	receiver := g.NewValue(naming.ReceiverName())
	instance := g.NewValue("instance")
	readArgs := g.StatementList()
	argNames := make([]g.Generator, len(op.Arguments))
	for i, a := range op.Arguments {
		argNames[i] = g.Id(a.Name)
		value := g.Raw(callArgument.Generate().Dot("Arguments").Index(jen.Lit(i)))
		converter := fmt.Sprintf("decode%s", a.Type)
		readArgs.Append(g.Assign(argNames[i], receiver.Field(converter).Call(value)))
	}
	list := g.StatementList(
		g.Assign(instance, receiver.Field("getInstance").Call(callArgument)),
		readArgs,
	)
	if op.HasResult() {
		converter := fmt.Sprintf("to%s", idlNameToGoName(op.RetType.TypeName))
		if op.GetHasError() {
			list.Append(
				g.AssignMany(g.List(
					g.Id("result"), g.Id("err")),
					instance.Field(UpperCaseFirstLetter(op.Name)).Call(argNames...),
				),
				panicOnNotNil(g.Id("err")),
			)
		} else {
			list.Append(
				g.Assign(
					g.Id("result"),
					instance.Field(UpperCaseFirstLetter(op.Name)).Call(argNames...),
				),
			)
		}
		list.Append(g.Return(receiver.Field(converter).Call(g.Id("result"))))
	} else {
		if op.GetHasError() {
			list.Append(
				g.Assign(g.Id("err"), instance.Field(UpperCaseFirstLetter(op.Name)).Call(argNames...)),
				panicOnNotNil(g.Id("err")),
			)

		} else {
			list.Append(instance.Field(UpperCaseFirstLetter(op.Name)).Call(argNames...))
		}
	}
	return list
}

func (g GojaTargetGenerators) WrapperStructGenerators() PlatformWrapperStructGenerators {
	return g
}

func (g GojaTargetGenerators) WrapperStructType(interfaceName string) generators.Type {
	return generators.NewType(fmt.Sprintf("%sWrapper", LowerCaseFirstLetter(interfaceName)))
}

func (g GojaTargetGenerators) WrapperStructConstructorName(interfaceName string) string {
	return fmt.Sprintf("new%sWrapper", interfaceName)
}

func (g GojaTargetGenerators) WrapperStructConstructorRetType(string) Generator {
	return generators.Id("wrapper")
}

func (g GojaTargetGenerators) EmbeddedType(wrappedType Generator) Generator {
	return generators.NewType("baseInstanceWrapper").TypeParam(wrappedType)
}

func (g GojaTargetGenerators) EmbeddedTypeConstructor(wrappedType Generator) generators.Value {
	return generators.NewValue("newBaseInstanceWrapper").TypeParam(wrappedType)
}

func (g GojaTargetGenerators) HostArg() Generator {
	return generators.Id("instance")
}

func (g GojaTargetGenerators) HostType() Generator {
	return generators.NewType("GojaContext").Pointer()
}

func (g GojaTargetGenerators) CallbackMethodArgs() generators.FunctionArgumentList {
	callArgument := generators.Id("c")
	return generators.Arg(callArgument, gojaFc)
}

func (g GojaTargetGenerators) CallbackMethodRetTypes() []generators.Generator {
	return []generators.Generator{gojaValue}
}

func panicOnNotNil(lhs g.Generator) g.Generator {
	return g.IfStmt{
		Condition: g.Neq{Lhs: lhs, Rhs: g.Nil},
		Block:     g.Raw(jen.Panic(jen.Id("err"))),
	}
}
