package gojagen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	. "github.com/gost-dom/code-gen/internal"
	wrappers "github.com/gost-dom/code-gen/script-wrappers"
	"github.com/gost-dom/code-gen/script-wrappers/model"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
)

type GojaNamingStrategy struct {
	model.ESConstructorData
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
func (gen GojaTargetGenerators) CreateHostInitializer(model.ESConstructorData) g.Generator {
	return g.Noop
}

func (gen GojaTargetGenerators) CreateConstructorCallbackBody(
	model.ESConstructorData,
	wrappers.CallbackContext,
) g.Generator {
	return g.Raw(jen.Panic(jen.Lit("Goja constructor not yet implemented")))
}

func (gen GojaTargetGenerators) CreateIllegalConstructorCallback(
	model.ESConstructorData,
) g.Generator {
	return g.Raw(jen.Panic(jen.Lit("Goja constructor not yet implemented")))
}

func (gen GojaTargetGenerators) CreateInitFunction(data model.ESConstructorData) g.Generator {
	naming := GojaNamingStrategy{data}
	return g.FunctionDefinition{
		Name: "init",
		Body: g.NewValue("installClass").
			Call(
				g.Lit(data.Name()),
				g.Lit(data.IdlInterface.Inheritance),
				g.Id(naming.PrototypeWrapperConstructorName()),
			),
	}
}

// CreatePrototypeInitializer creates the "initializePrototype" method, which
// sets all the properties on the prototypes on this class.
func (gen GojaTargetGenerators) CreatePrototypeInitializer(
	data model.ESConstructorData,
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
	data model.ESConstructorData,
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
			getter = vm.Field("ToValue").Call(receiver.Field(a.Getter.CallbackMethodName()))
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

func (gen GojaTargetGenerators) ReturnErrMsg(errGen g.Generator) g.Generator {
	return g.Raw(jen.Panic(errGen.Generate()))
}

func (gen GojaTargetGenerators) CreateAttributeGetter(
	data model.ESConstructorData,
	op model.ESOperation,
	cbCtx wrappers.CallbackContext,
	eval func(g.Generator) g.Generator,
) g.Generator {
	// return gen.CreateMethodCallbackBody(data, op)
	callArgument := g.Id("c")
	naming := GojaNamingStrategy{data}
	receiver := g.NewValue(naming.ReceiverName())
	instance := g.NewValue("instance")
	return g.StatementList(
		g.Assign(instance, receiver.Field("getInstance").Call(callArgument)),
		gen.ConvertResult(op, receiver, eval(instance)),
	)
}

func (gen GojaTargetGenerators) CreateAttributeSetter(
	data model.ESConstructorData,
	op model.ESOperation,
	cbCtx wrappers.CallbackContext,
	updateValue func(g.Generator, g.Generator) g.Generator,
) g.Generator {
	callArgument := g.Id("c")
	naming := GojaNamingStrategy{data}
	receiver := g.NewValue(naming.ReceiverName())
	instance := g.NewValue("instance")
	readArgs := g.StatementList()
	argNames := make([]g.Generator, len(op.Arguments))
	for i, a := range op.Arguments {
		argNames[i] = g.Id(a.Name)
		value := g.Raw(callArgument.Generate().Dot("Arguments").Index(jen.Lit(i)))
		converter := fmt.Sprintf("decode%s", a.IdlArg.Type.Name)
		readArgs.Append(g.Assign(argNames[i], receiver.Field(converter).Call(value)))
	}
	return g.StatementList(
		g.Assign(instance, receiver.Field("getInstance").Call(callArgument)),
		readArgs,
		gen.ConvertResult(op, receiver, updateValue(instance, argNames[0])),
	)
}

func (gen GojaTargetGenerators) CreateMethodCallbackBody(
	data model.ESConstructorData,
	op model.ESOperation,
	cbCtx wrappers.CallbackContext,
) g.Generator {
	callArgument := g.Id("c")
	naming := GojaNamingStrategy{data}
	receiver := g.NewValue(naming.ReceiverName())
	instance := g.NewValue("instance")
	readArgs := g.StatementList()
	argNames := make([]g.Generator, len(op.Arguments))
	for i, a := range op.Arguments {
		argNames[i] = g.Id(a.Name)
		value := g.Raw(callArgument.Generate().Dot("Arguments").Index(jen.Lit(i)))
		converter := fmt.Sprintf("decode%s", a.IdlArg.Type.Name)
		readArgs.Append(g.Assign(argNames[i], receiver.Field(converter).Call(value)))
	}
	return g.StatementList(
		g.Assign(instance, receiver.Field("getInstance").Call(callArgument)),
		readArgs,
		gen.ConvertResult(op, receiver,
			instance.Field(UpperCaseFirstLetter(op.Name)).Call(argNames...),
		),
	)
}

func (gen GojaTargetGenerators) ConvertResult(
	op model.ESOperation,
	receiver g.Value,
	evaluate g.Generator,
) g.Generator {
	list := g.StatementList()
	if op.HasResult() {
		converter := op.Encoder()
		if op.GetHasError() {
			list.Append(
				g.AssignMany(g.List(
					g.Id("result"), g.Id("err")),
					evaluate,
				),
				panicOnNotNil(g.Id("err")),
			)
		} else {
			list.Append(
				g.Assign(g.Id("result"), evaluate),
			)
		}
		list.Append(g.Return(receiver.Field(converter).Call(g.Id("result"))))
	} else {
		if op.GetHasError() {
			list.Append(
				g.Assign(g.Id("err"), evaluate),
				panicOnNotNil(g.Id("err")),
			)

		} else {
			list.Append(evaluate)
			list.Append(g.Return(g.Nil))
		}
	}
	return list
}

func (g GojaTargetGenerators) WrapperStructGenerators() wrappers.PlatformWrapperStructGenerators {
	return g
}

func (g GojaTargetGenerators) WrapperStructType(interfaceName string) generators.Type {
	return generators.NewType(fmt.Sprintf("%sWrapper", LowerCaseFirstLetter(interfaceName)))
}

func (g GojaTargetGenerators) WrapperStructConstructorName(interfaceName string) string {
	return fmt.Sprintf("new%sWrapper", interfaceName)
}

func (g GojaTargetGenerators) WrapperStructConstructorRetType(string) g.Generator {
	return generators.Id("wrapper")
}

func (g GojaTargetGenerators) EmbeddedType(wrappedType g.Generator) g.Generator {
	return generators.NewType("baseInstanceWrapper").TypeParam(wrappedType)
}

func (g GojaTargetGenerators) EmbeddedTypeConstructor(wrappedType g.Generator) g.Value {
	return generators.NewValue("newBaseInstanceWrapper").TypeParam(wrappedType)
}

func (g GojaTargetGenerators) HostArg() g.Generator {
	return generators.Id("instance")
}

func (g GojaTargetGenerators) HostType() g.Generator {
	return generators.NewType("GojaContext").Pointer()
}

func (g GojaTargetGenerators) CallbackMethodArgs() g.FunctionArgumentList {
	callArgument := generators.Id("c")
	return generators.Arg(callArgument, gojaFc)
}

func (g GojaTargetGenerators) CallbackMethodRetTypes() []g.Generator {
	return []generators.Generator{gojaValue}
}

func panicOnNotNil(lhs g.Generator) g.Generator {
	return g.IfStmt{
		Condition: g.Neq{Lhs: lhs, Rhs: g.Nil},
		Block:     g.Raw(jen.Panic(jen.Id("err"))),
	}
}
