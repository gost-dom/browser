package v8gen

import (
	"fmt"

	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/scripting"
	. "github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"

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

func CreateV8ConstructorBody(data ESConstructorData) g.Generator {
	naming := V8NamingStrategy{data}
	builder := NewConstructorBuilder()
	constructor := v8FunctionTemplate{g.NewValue("constructor")}

	createWrapperFunction := g.NewValue(fmt.Sprintf("new%s", naming.PrototypeWrapperBaseName()))

	statements := g.StatementList(
		g.Assign(builder.Wrapper, createWrapperFunction.Call(scriptHost)),
		g.Assign(constructor, wrapCallback(scriptHost, builder.Wrapper.Field("constructor"))),
		g.Line,
		g.Assign(builder.InstanceTmpl, constructor.GetInstanceTemplate()),
		builder.InstanceTmpl.SetInternalFieldCount(1),
		g.Line,
		g.Assign(builder.Class, newJSClass(scriptHost, constructor)),
		builder.Wrapper.InitializeClass(builder.Class),
		g.Line,
	)
	if data.RunCustomCode {
		statements.Append(
			g.NewValue("wrapper").Field("CustomInitializer").Call(builder.Class),
		)
	}
	statements.Append(g.Return(builder.Class))
	return statements
}

func CreateV8ClassInitializerBody(data ESConstructorData) g.Generator {
	builder := NewConstructorBuilder()
	statements := g.StatementList(
		builder.Wrapper.InitializeClass(builder.Class),
	)
	if data.RunCustomCode {
		statements.Append(
			g.NewValue("wrapper").Field("CustomInitializer").Call(builder.Class),
		)
	}
	return statements
}

func CreateV8ConstructorWrapperBody(
	data ESConstructorData,
	cbCtx scripting.CallbackContext,
) g.Generator {
	op := *data.Constructor
	naming := V8NamingStrategy{data}
	receiver := g.NewValue(naming.Receiver())
	return V8CallbackGenerators{
		Data:     data,
		Op:       op,
		Receiver: receiver,
	}.ConstructorCallback(cbCtx)
}

func CreateV8IllegalConstructorBody(
	data ESConstructorData,
	cbCtx scripting.CallbackContext,
) g.Generator {
	return g.Return(cbCtx.ReturnWithTypeError("Illegal Constructor"))
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
	cbCtx scripting.CallbackContext,
) (res V8ReadArguments) {
	naming := V8NamingStrategy{data}
	argCount := len(op.Arguments)
	res.Args = make([]V8ReadArg, 0, argCount)
	statements := g.StatementList()
	receiver := g.NewValue(naming.Receiver())
	for i, arg := range op.Arguments {
		argName := g.Id(scripting.SanitizeVarName(arg.Name))
		errName := g.Id(fmt.Sprintf("err%d", i+1))
		if arg.Ignore {
			statements.Append(g.NewValue("ignoreArgument").Call(cbCtx))
			continue
		}
		res.Args = append(res.Args, V8ReadArg{
			Argument: arg,
			ArgName:  argName,
			ErrName:  errName,
			Index:    i,
		})

		var dec = scripting.DecodersForArg(receiver, arg)

		defaultName, hasDefault := arg.DefaultValueInGo()
		nullable := arg.IdlArg.Type.Nullable
		parseArgs := []g.Generator{cbCtx, g.Lit(arg.Name)}
		if hasDefault {
			parseArgs = append(parseArgs, g.NewValue(naming.Receiver()).Field(defaultName))
		} else if nullable {
			parseArgs = append(parseArgs, ZeroValue)
		} else {
			parseArgs = append(parseArgs, g.Nil)
		}
		parseArgs = append(parseArgs, dec...)
		statements.Append(g.AssignMany(g.List(argName, errName),
			ConsumeArgument.Call(parseArgs...)))
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
