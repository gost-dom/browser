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

func (g GojaTargetGenerators) Name() string { return "goja" }

func (gen GojaTargetGenerators) ConstructorCallbackEnabled() bool { return false }

func (gen GojaTargetGenerators) Host(receiver g.Generator) g.Generator {
	return g.ValueOf(receiver).Field("ctx")
}

func (gen GojaTargetGenerators) PlatformInfoArg() g.Generator { return g.Id("c") }

// CreateConstructor has no effect for Goja. It's currently based on a system
// that it automatically creates the constructors based on whether or not they
// call the ingerface
func (gen GojaTargetGenerators) CreateHostInitializer(model.ESConstructorData) g.Generator {
	return g.Noop
}

func (gen GojaTargetGenerators) CreateConstructorCallbackBody(
	data model.ESConstructorData,
	cbCtx wrappers.CallbackContext,
) g.Generator {
	return g.Return(cbCtx.ReturnWithTypeError("Goja constructor not yet implemented"))
}

func (gen GojaTargetGenerators) CreateIllegalConstructorCallback(
	model.ESConstructorData, wrappers.CallbackContext,
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
	ctx := receiver.Field("ctx")
	prototype := g.NewValue("prototype")
	body := g.StatementList()
	for op := range data.WrapperFunctionsToInstall() {
		body.Append(
			prototype.Field("Set").Call(g.Lit(op.Name), wrapCallback(ctx, receiver.Field(op.Name))),
		)
	}

	for a := range data.AttributesToInstall() {
		var getter, setter g.Generator
		if a.Getter != nil {
			getter = wrapCallback(ctx, receiver.Field(a.Getter.CallbackMethodName()))
		} else {
			getter = g.Nil
		}
		if a.Setter != nil {
			setter = wrapCallback(ctx, receiver.Field(a.Setter.Name))
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

func wrapCallback(ctx, callback g.Generator) g.Generator {
	return g.NewValue("wrapCallback").Call(ctx, callback)
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
	naming := GojaNamingStrategy{data}
	receiver := g.NewValue(naming.ReceiverName())
	instance := g.NewValue("instance")
	return g.StatementList(
		gen.getInstance(cbCtx, data, instance),
		gen.ConvertResult(op, data, receiver, cbCtx, eval(instance)),
	)
}

func (gen GojaTargetGenerators) getInstance(
	cbCtx wrappers.CallbackContext,
	data model.ESConstructorData,
	instance g.Generator,
) g.Generator {
	err := g.Id("instErr")
	return g.StatementList(
		g.AssignMany(
			g.List(instance, err),
			wrappers.As.TypeParam(data.WrappedType()).Call(cbCtx.GetInstance()),
		),
		wrappers.IfError(err, wrappers.TransformerFunc(func(err g.Generator) g.Generator {
			return g.Return(g.Nil, err)
		})),
	)
}

func (gen GojaTargetGenerators) CreateAttributeSetter(
	data model.ESConstructorData,
	op model.ESOperation,
	cbCtx wrappers.CallbackContext,
	updateValue func(g.Generator, g.Generator) g.Generator,
) g.Generator {
	naming := GojaNamingStrategy{data}
	receiver := g.NewValue(naming.ReceiverName())
	instance := g.NewValue("instance")
	readArgs := g.StatementList()
	argNames := make([]g.Generator, len(op.Arguments))
	for i, a := range op.Arguments {
		argNames[i] = g.Id(a.Name)
		value := g.Raw(cbCtx.Generate().Dot("Argument").Call(jen.Lit(i)))
		converter := fmt.Sprintf("decode%s", a.IdlArg.Type.Name)
		readArgs.Append(g.Assign(argNames[i], receiver.Field(converter).Call(value)))
	}
	return g.StatementList(
		gen.getInstance(cbCtx, data, instance),
		readArgs,
		gen.ConvertResult(op, data, receiver, cbCtx, updateValue(instance, argNames[0])),
	)
}

func (gen GojaTargetGenerators) CreateMethodCallbackBody(
	data model.ESConstructorData,
	op model.ESOperation,
	cbCtx wrappers.CallbackContext,
) g.Generator {
	naming := GojaNamingStrategy{data}
	receiver := g.NewValue(naming.ReceiverName())
	instance := g.NewValue("instance")
	readArgs := g.StatementList()
	argNames := make([]g.Generator, len(op.Arguments))
	for i, a := range op.Arguments {
		argNames[i] = g.Id(a.Name)
		value := g.Raw(cbCtx.Generate().Dot("Argument").Call(jen.Lit(i)))
		converter := fmt.Sprintf("decode%s", a.IdlArg.Type.Name)
		readArgs.Append(g.Assign(argNames[i], receiver.Field(converter).Call(value)))
	}
	return g.StatementList(
		// g.Assign(instance, receiver.Field("getInstance").Call(callArgument)),
		gen.getInstance(cbCtx, data, instance),
		readArgs,
		gen.ConvertResult(op, data, receiver, cbCtx,
			instance.Field(UpperCaseFirstLetter(op.Name)).Call(argNames...),
		),
	)
}

func (gen GojaTargetGenerators) ConvertResult(
	op model.ESOperation,
	data model.ESConstructorData,
	receiver g.Value,
	cbCtx wrappers.CallbackContext,
	evaluate g.Generator,
) g.Generator {
	list := g.StatementList()
	if op.HasResult() {
		converter := op.Encoder(false, receiver, cbCtx, data)
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
		list.Append(
			g.Return(cbCtx.ReturnWithValueErr(converter.Call(op.RetValues(data)...))),
		)
	} else {
		if op.GetHasError() {
			list.Append(
				g.Assign(g.Id("err"), evaluate),
				panicOnNotNil(g.Id("err")),
			)

		} else {
			list.Append(evaluate)
			list.Append(g.Return(g.Nil, g.Nil))
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

func (g GojaTargetGenerators) WrapperStructTypeDef(interfaceName string) generators.Type {
	return generators.NewType(fmt.Sprintf("%sWrapper", LowerCaseFirstLetter(interfaceName)))
}

func (g GojaTargetGenerators) WrapperStructTypeRetDef(interfaceName string) generators.Type {
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

func (g GojaTargetGenerators) EmbeddedTypeConstructor(wrappedType g.Generator) g.Generator {
	return generators.NewValue("newBaseInstanceWrapper").TypeParam(wrappedType)
}

func (g GojaTargetGenerators) HostArg() g.Generator {
	return generators.Id("instance")
}

func (g GojaTargetGenerators) HostType() g.Generator {
	return generators.NewType("GojaContext").Pointer()
}

func (g GojaTargetGenerators) CallbackMethodArgs() g.FunctionArgumentList {
	callArgument := generators.Id("cbCtx")
	return generators.Arg(callArgument, gojaCbCtx)
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
