package v8gen

import (
	wrappers "github.com/gost-dom/code-gen/script-wrappers"
	"github.com/gost-dom/code-gen/script-wrappers/model"
	. "github.com/gost-dom/code-gen/script-wrappers/model"
	"github.com/gost-dom/code-gen/stdgen"
	g "github.com/gost-dom/generators"
)

type V8TargetGenerators struct{}

func (gen V8TargetGenerators) CreateInitFunction(data ESConstructorData) g.Generator {
	return g.FunctionDefinition{
		Name: "init",
		Body: g.NewValue("registerJSClass").Call(
			g.Lit(data.Spec.TypeName),
			g.Lit(data.IdlInterface.Inheritance),
			g.Id(prototypeFactoryFunctionName(data))),
	}
}

func (gen V8TargetGenerators) ReturnErrMsg(errGen g.Generator) g.Generator {
	return g.Return(g.Nil, stdgen.ErrorsNew(errGen))
}

func (gen V8TargetGenerators) WrapperStructGenerators() wrappers.PlatformWrapperStructGenerators {
	return V8WrapperStructGenerators{}
}

func (gen V8TargetGenerators) CreatePrototypeInitializer(
	data ESConstructorData,
	body g.Generator,
) g.Generator {
	naming := V8NamingStrategy{data}
	receiver := g.NewValue(naming.Receiver())
	builder := NewConstructorBuilder()
	return g.FunctionDefinition{
		Name: "installPrototype",
		Receiver: g.FunctionArgument{
			Name: receiver,
			Type: g.Id(naming.PrototypeWrapperName()),
		},
		Args: g.Arg(builder.Proto, v8ObjectTemplatePtr),
		Body: body,
	}
}

func (gen V8TargetGenerators) CreatePrototypeInitializerBody(
	data ESConstructorData,
) g.Generator {
	naming := V8NamingStrategy{data}
	receiver := g.NewValue(naming.Receiver())
	builder := NewConstructorBuilder()
	installer := PrototypeInstaller{
		builder.v8Iso,
		builder.Proto,
		WrapperInstance{g.Value{Generator: receiver}},
	}
	return g.StatementList(
		g.Assign(g.NewValue("iso"), receiver.Field("scriptHost").Field("iso")),
		installer.InstallFunctionHandlers(data),
		installer.InstallAttributeHandlers(data),
	)
}

func (gen V8TargetGenerators) CreateConstructorCallback(data ESConstructorData) g.Generator {
	naming := V8NamingStrategy{data}
	var body g.Generator
	if model.IsNodeType(data.IdlInterface.Name) {
		body = CreateV8IllegalConstructorBody(data)
	} else {
		body = CreateV8ConstructorWrapperBody(data)
	}
	return g.StatementList(
		g.Line,
		g.FunctionDefinition{
			Name: "Constructor",
			Receiver: g.FunctionArgument{
				Name: g.Id(naming.Receiver()),
				Type: g.Id(naming.PrototypeWrapperName()),
			},
			Args:     g.Arg(g.Id("info"), v8FunctionCallbackInfoPtr),
			RtnTypes: g.List(v8Value, g.Id("error")),
			Body:     body,
		},
	)
}

func (gen V8TargetGenerators) CreateAttributeGetter(
	data ESConstructorData,
	op ESOperation,
	eval func(g.Generator) g.Generator,
) g.Generator {
	instance := g.NewValue("instance")
	err := g.NewValue("err")
	ctx := g.NewValue("ctx")
	naming := V8NamingStrategy{data}
	receiver := WrapperInstance{g.NewValue(naming.Receiver())}

	x := V8InstanceInvocation{
		Name:     "",
		Args:     nil,
		Op:       op,
		Instance: &instance,
		Receiver: receiver,
	}
	return g.StatementList(
		V8RequireContext(receiver),
		GetInstanceAndError(instance, err, data),
		wrappers.ReturnIfError(err),
		x.ConvertResult(ctx, eval(instance)),
	)
}

func (gen V8TargetGenerators) CreateAttributeSetter(
	data ESConstructorData,
	op ESOperation,
	set func(g.Generator, g.Generator) g.Generator,
) g.Generator {
	var (
		err      = g.Id("err0")
		err1     = g.Id("err1")
		val      = g.Id("val")
		instance = g.NewValue("instance")
	)

	naming := V8NamingStrategy{data}
	receiver := WrapperInstance{g.NewValue(naming.Receiver())}

	args := append(
		[]g.Generator{g.Id("ctx"), g.Id("info")},
		wrappers.DecodersForArg(receiver, op.Arguments[0])...,
	)
	parsedArg := g.NewValue("parseSetterArg").
		Call(args...)

	return g.StatementList(
		V8RequireContext(receiver),
		GetInstanceAndError(instance, err, data),
		g.AssignMany(g.List(val, err1), parsedArg),

		wrappers.ReturnOnAnyError([]g.Generator{err, err1}),
		set(instance, val),
		g.Return(g.Nil, g.Nil),
	)
}

func (gen V8TargetGenerators) CreateMethodCallbackBody(
	data ESConstructorData,
	op ESOperation,
) g.Generator {
	naming := V8NamingStrategy{data}
	receiver := WrapperInstance{g.NewValue(naming.Receiver())}
	instance := g.NewValue("instance")
	readArgsResult := ReadArguments(data, op)
	err := g.Id("err0")
	if len(op.Arguments) == 0 {
		err = g.Id("err")
	}
	ctx := g.NewValue("ctx")
	requireContext := false
	var CreateCall = func(functionName string, argnames []g.Generator, op ESOperation) g.Generator {
		if op.Name == "toString" {
			functionName = "String"
		}
		requireContext = requireContext || op.HasResult()
		return V8InstanceInvocation{
			Name:     functionName,
			Args:     argnames,
			Op:       op,
			Instance: &instance,
			Receiver: receiver,
		}.GetGenerator(ctx)
	}
	statements := g.StatementList(
		AssignArgs(data, op),
		GetInstanceAndError(instance, err, data),
		readArgsResult,
		CreateV8WrapperMethodInstanceInvocations(
			data,
			op,
			IdlNameToGoName(op.Name),
			readArgsResult.Args,
			err,
			CreateCall,
			true,
		),
	)
	if requireContext {
		statements.Prepend(V8RequireContext(receiver))
	}
	return statements
}

func (gen V8TargetGenerators) CreateHostInitializer(data ESConstructorData) g.Generator {
	return g.FunctionDefinition{
		Name:     prototypeFactoryFunctionName(data),
		Args:     g.Arg(scriptHost, scriptHostPtr),
		RtnTypes: g.List(v8FunctionTemplatePtr),
		Body:     CreateV8ConstructorBody(data),
	}
}

func (gen V8TargetGenerators) PlatformInfoArg() g.Generator { return g.Id("info") }
