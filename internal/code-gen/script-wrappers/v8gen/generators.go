package v8gen

import (
	"fmt"

	. "github.com/gost-dom/code-gen/internal"
	wrappers "github.com/gost-dom/code-gen/script-wrappers"
	"github.com/gost-dom/code-gen/script-wrappers/model"
	. "github.com/gost-dom/code-gen/script-wrappers/model"
	"github.com/gost-dom/code-gen/stdgen"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"

	"github.com/dave/jennifer/jen"
)

var scriptHost = g.NewValue("scriptHost")

type V8NamingStrategy struct{ ESConstructorData }

func (s V8NamingStrategy) Receiver() string { return "w" }
func (s V8NamingStrategy) PrototypeWrapperBaseName() string {
	return fmt.Sprintf("%sV8Wrapper", s.Name())
}

func (s V8NamingStrategy) PrototypeWrapperName() string {
	return LowerCaseFirstLetter(s.PrototypeWrapperBaseName())
}

func prototypeFactoryFunctionName(data ESConstructorData) string {
	return fmt.Sprintf("create%sPrototype", data.IdlInterface.Name)
}

func CreateV8ConstructorBody(data ESConstructorData) g.Generator {
	naming := V8NamingStrategy{data}
	builder := NewConstructorBuilder()
	constructor := v8FunctionTemplate{g.NewValue("constructor")}

	createWrapperFunction := g.NewValue(fmt.Sprintf("new%s", naming.PrototypeWrapperBaseName()))

	statements := g.StatementList(
		builder.v8Iso.Assign(scriptHost.Field("iso")),
		g.Assign(builder.Wrapper, createWrapperFunction.Call(scriptHost)),
		g.Assign(constructor, builder.NewFunctionTemplateOfWrappedMethod("Constructor")),
		g.Line,
		g.Assign(builder.InstanceTmpl, constructor.GetInstanceTemplate()),
		builder.InstanceTmpl.SetInternalFieldCount(1),
		g.Line,
		builder.Wrapper.Field("installPrototype").Call(constructor.GetPrototypeTemplate()),
		g.Line,
	)
	if data.RunCustomCode {
		statements.Append(
			g.Raw(jen.Id("wrapper").Dot("CustomInitialiser").Call(jen.Id("constructor"))),
		)
	}
	statements.Append(g.Return(constructor))
	return statements
}

func CreateV8ConstructorWrapperBody(
	data ESConstructorData,
	cbCtx wrappers.CallbackContext,
) g.Generator {
	naming := V8NamingStrategy{data}
	var readArgsResult V8ReadArguments
	op := *data.Constructor
	readArgsResult = ReadArguments(data, op, cbCtx)
	cbInfo := g.NewValue("info")
	statements := g.StatementList(
		readArgsResult)
	receiver := g.NewValue(naming.Receiver())
	baseFunctionName := "CreateInstance"
	var CreateCall = func(functionName string, argnames []g.Generator, op ESOperation) g.Generator {
		return g.StatementList(
			g.Return(
				receiver.Field(functionName).Call(append([]g.Generator{cbCtx.Context(),
					cbInfo.Field("This").Call()},
					argnames...)...,
				),
			))
	}
	statements.Append(
		CreateV8WrapperMethodInstanceInvocations(
			data,
			op,
			cbCtx,
			baseFunctionName,
			readArgsResult.Args,
			nil,
			CreateCall,
			false,
		),
	)
	return statements
}

func CreateV8WrapperMethodInstanceInvocations(
	prototype ESConstructorData,
	op ESOperation,
	cbCtx wrappers.CallbackContext,
	baseFunctionName string,
	args []V8ReadArg,
	instanceErr g.Generator,
	createCallInstance func(string, []g.Generator, ESOperation) g.Generator,
	extraError bool,
) g.Generator {
	statements := g.StatementList()
	missingArgsConts := fmt.Sprintf("%s.%s: Missing arguments", prototype.Name(), op.Name)
	for i := len(args); i >= 0; i-- {
		functionName := baseFunctionName
		for j, arg := range args {
			if j < i {
				if arg.Argument.OptionalInGo() {
					functionName += IdlNameToGoName(arg.Argument.Name)
				}
			}
		}
		currentArgs := args[0:i]
		ei := i
		if extraError {
			ei++
		}
		errNames := make([]g.Generator, 0, i+1)
		if instanceErr != nil {
			errNames = append(errNames, instanceErr)
		}
		for _, a := range currentArgs {
			errNames = append(errNames, a.ErrName)
		}

		callArgs := make([]g.Generator, i)
		for idx, a := range currentArgs {
			arg := a.ArgName
			if a.Argument.CustomRule.Variadic {
				arg = g.Raw(arg.Generate().Op("..."))
			}
			callArgs[idx] = arg
		}
		callInstance := createCallInstance(functionName, callArgs, op)
		if i > 0 {
			arg := args[i-1].Argument
			statements.Append(g.StatementList(
				g.IfStmt{
					Condition: g.Raw(cbCtx.Generate().Dot("noOfReadArguments").Op(">=").Lit(i)),
					Block: g.StatementList(
						wrappers.ReturnOnAnyError(errNames),
						callInstance,
					),
				}))
			if !arg.OptionalInGo() {
				statements.Append(
					g.Return(
						g.Nil,
						stdgen.ErrorsNew(g.Lit(missingArgsConts)),
					),
				)
				break
			}
		} else {
			statements.Append(wrappers.ReturnOnAnyError(errNames))
			statements.Append(callInstance)
		}
	}
	return statements
}

type V8InstanceInvocation struct {
	Name     string
	Args     []g.Generator
	Op       ESOperation
	Instance *g.Value
	Receiver WrapperInstance
}

type V8InstanceInvocationResult struct {
	Generator      g.Generator
	RequireContext bool
}

func (c V8InstanceInvocation) AssignValues(evaluation g.Generator) g.Generator {
	HasError := c.Op.GetHasError()
	HasValue := c.Op.HasResult()
	if !HasError && !HasValue {
		return evaluation
	}
	if HasError && !HasValue {
		return g.Assign(g.Id("callErr"), evaluation)
	}
	if !HasError && HasValue {
		return g.Assign(g.Id("result"), evaluation)
	}
	return g.AssignMany([]g.Generator{
		g.Id("result"),
		g.Id("callErr")}, evaluation)
}

func (c V8InstanceInvocation) ConvertResult(
	ctx g.Value,
	evaluation g.Generator,
) g.Generator {
	hasError := c.Op.GetHasError()
	hasValue := c.Op.HasResult() // != "undefined"

	list := g.StatementListStmt{}
	list.Append(c.AssignValues(evaluation))

	if !hasValue {
		if hasError {
			list.Append(g.Return(g.Nil, g.Id("callErr")))
		} else {
			list.Append(g.Return(g.Nil, g.Nil))
		}
	} else {
		returnValue := c.ConvertReturnValue(ctx, c.Op.RetType)
		if hasError {
			list.Append(g.IfStmt{
				Condition: g.Neq{Lhs: g.Id("callErr"), Rhs: g.Nil},
				Block:     g.Return(g.Nil, g.Id("callErr")),
				Else:      returnValue,
			})
		} else {
			list.Append(returnValue)
		}
	}
	return list
}

func (c V8InstanceInvocation) ConvertReturnValue(ctx g.Value, retType idl.Type) g.Generator {
	if model.IsNodeType(retType.Name) {
		return g.Return(ctx.Field("getInstanceForNode").Call(g.Id("result")))
	} else {
		converter := c.Op.Encoder()
		return g.Return(c.Receiver.Method(converter).Call(ctx, g.Id("result")))
	}
}

func (c V8InstanceInvocation) GetGenerator(ctx g.Value) g.Generator {
	if c.Instance == nil {
		return c.ConvertResult(ctx, g.NewValue(IdlNameToGoName(c.Name)).Call(c.Args...))
	} else {
		return c.ConvertResult(ctx, c.Instance.Method(IdlNameToGoName(c.Name)).Call(c.Args...))
	}
}

func CreateV8IllegalConstructorBody(data ESConstructorData) g.Generator {
	naming := V8NamingStrategy{data}
	return g.Return(g.Nil, g.NewValuePackage("NewTypeError", v8).
		Call(g.NewValue(naming.Receiver()).Field("scriptHost").Field("iso"),
			g.Lit("Illegal Constructor")))
}

type V8ReadArg struct {
	Argument ESOperationArgument
	ArgName  g.Generator
	ErrName  g.Generator
	Index    int
}

type V8ReadArguments struct {
	Args      []V8ReadArg
	Generator g.Generator
}

func (r V8ReadArguments) Generate() *jen.Statement {
	if r.Generator != nil {
		return r.Generator.Generate()
	} else {
		return g.Noop.Generate()
	}
}

func ReadArguments(
	data ESConstructorData,
	op ESOperation,
	cbCtx wrappers.CallbackContext,
) (res V8ReadArguments) {
	naming := V8NamingStrategy{data}
	argCount := len(op.Arguments)
	res.Args = make([]V8ReadArg, 0, argCount)
	statements := g.StatementList()
	receiver := g.NewValue(naming.Receiver())
	for i, arg := range op.Arguments {
		argName := g.Id(wrappers.SanitizeVarName(arg.Name))
		errName := g.Id(fmt.Sprintf("err%d", i+1))
		if arg.Ignore {
			continue
		}
		res.Args = append(res.Args, V8ReadArg{
			Argument: arg,
			ArgName:  argName,
			ErrName:  errName,
			Index:    i,
		})

		var dec = wrappers.DecodersForArg(receiver, arg)

		gConverters := []g.Generator{cbCtx, g.Lit(i)}
		defaultName, hasDefault := arg.DefaultValueInGo()
		if hasDefault {
			gConverters = append(gConverters, g.NewValue(naming.Receiver()).Field(defaultName))
		}
		gConverters = append(gConverters, dec...)
		if hasDefault {
			statements.Append(g.AssignMany(g.List(argName, errName),
				g.NewValue("tryParseArgWithDefault").Call(gConverters...)))
		} else if arg.IdlArg.Type.Nullable {
			statements.Append(g.AssignMany(
				g.List(argName, errName),
				g.NewValue("tryParseArgNullableType").Call(gConverters...)))
		} else {
			statements.Append(g.AssignMany(
				g.List(argName, errName),
				g.NewValue("tryParseArg").Call(gConverters...)))
		}
	}
	res.Generator = statements
	return
}

func GetInstanceAndError(id g.Generator, errId g.Generator, data ESConstructorData) g.Generator {
	naming := V8NamingStrategy{data}
	return g.AssignMany(
		g.List(id, errId),
		g.NewValue(naming.Receiver()).Field("getInstance").Call(g.Id("info")),
	)
}
