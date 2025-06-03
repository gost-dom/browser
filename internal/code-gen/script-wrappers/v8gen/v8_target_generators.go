package v8gen

import (
	wrappers "github.com/gost-dom/code-gen/script-wrappers"
	. "github.com/gost-dom/code-gen/script-wrappers/model"
	"github.com/gost-dom/code-gen/stdgen"
	g "github.com/gost-dom/generators"
)

type V8TargetGenerators struct{}

func (gen V8TargetGenerators) Host(receiver g.Generator) g.Generator {
	return g.ValueOf(receiver).Field("scriptHost")
}

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
		Args: g.Arg(builder.Class, v8Class),
		Body: body,
	}
}

func (gen V8TargetGenerators) CreatePrototypeInitializerBody(
	data ESConstructorData,
) g.Generator {
	naming := V8NamingStrategy{data}
	receiver := g.NewValue(naming.Receiver())
	host := receiver.Field("scriptHost")
	ft := v8FunctionTemplate{g.NewValue("ft")}
	return PrototypeInstaller{
		ft,
		WrapperInstance{g.Value{Generator: receiver}},
		host,
		data,
		receiver,
	}
}

func (gen V8TargetGenerators) CreateConstructorCallbackBody(
	data ESConstructorData,
	cbCtx wrappers.CallbackContext,
) g.Generator {
	return CreateV8ConstructorWrapperBody(data, cbCtx)
}

func (gen V8TargetGenerators) CreateIllegalConstructorCallback(
	data ESConstructorData,
	cbCtx wrappers.CallbackContext,
) g.Generator {
	return CreateV8IllegalConstructorBody(data, cbCtx)
}

func (gen V8TargetGenerators) CreateAttributeGetter(
	data ESConstructorData,
	op ESOperation,
	cbCtx wrappers.CallbackContext,
	eval func(g.Generator) g.Generator,
) g.Generator {
	naming := V8NamingStrategy{data}
	return V8CallbackGenerators{
		data,
		op,
		g.NewValue(naming.Receiver()),
	}.AttributeGetterCallback(
		cbCtx,
		wrappers.TransformerFunc(eval),
	)
}

func (gen V8TargetGenerators) CreateAttributeSetter(
	data ESConstructorData,
	op ESOperation,
	cbCtx wrappers.CallbackContext,
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
		[]g.Generator{cbCtx},
		wrappers.DecodersForArg(receiver, op.Arguments[0])...,
	)
	parsedArg := g.NewValue("parseSetterArg").Call(args...)

	return g.StatementList(
		g.AssignMany(
			g.List(instance, err),
			wrappers.As.TypeParam(data.WrappedType()).Call(cbCtx.GetInstance()),
		),
		g.AssignMany(g.List(val, err1), parsedArg),

		wrappers.IfAnyError(
			[]g.Generator{err, err1},
			wrappers.ReturnTransform(wrappers.TransformerFunc(cbCtx.ReturnWithError)),
		),
		set(instance, val),
		g.Return(cbCtx.ReturnWithValue(g.Nil)),
	)
}

func returnNilCommaErr(err g.Generator) g.Generator {
	return g.Return(g.Nil, err)
}

func (gen V8TargetGenerators) CreateMethodCallbackBody(
	data ESConstructorData,
	op ESOperation,
	cbCtx wrappers.CallbackContext,
) g.Generator {
	receiver := g.NewValue("w")
	return V8CallbackGenerators{data, op, receiver}.OperationCallback(cbCtx)
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
