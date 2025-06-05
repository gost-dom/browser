package wrappers

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/packagenames"
	. "github.com/gost-dom/code-gen/script-wrappers/model"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
)

type Generator = generators.Generator

// Create a struct definition, and it's constructor, that must contain the
// methods acting as callback for prototype functions, including the
// constructor.
type PlatformWrapperStructGenerators interface {
	WrapperStructType(interfaceName string) generators.Type
	WrapperStructTypeDef(interfaceName string) generators.Type
	WrapperStructTypeRetDef(interfaceName string) generators.Type
	WrapperStructConstructorName(interfaceName string) string
	WrapperStructConstructorRetType(interfaceName string) Generator
	EmbeddedType(wrappedType Generator) Generator
	EmbeddedTypeConstructor(wrappedType Generator) generators.Value
	CallbackMethodArgs() generators.FunctionArgumentList
	CallbackMethodRetTypes() []generators.Generator
	HostArg() Generator
	HostType() Generator
}

type TargetGenerators interface {
	Name() string
	Host(g.Generator) g.Generator
	// CreateInitFunction generates an init function intended to register that a
	// class should be created. This doesn't _create_ the class, as that
	// requires a host created at runtime. So this is a registration that _when_
	// a host is created, this class must be added to global scope, optionally
	// with a subclass.
	CreateInitFunction(ESConstructorData) Generator
	// CreateHostInitializer creates the function that will register the class
	// in the host's global scope.
	CreateHostInitializer(ESConstructorData) Generator
	// CreatePrototypeInitializer creates the "initializePrototype" method, which
	// sets the properties on the prototype object, both data properties for
	// methods, and accessor properties for attributes.
	CreatePrototypeInitializer(ESConstructorData, Generator) Generator
	CreatePrototypeInitializerBody(ESConstructorData) Generator
	ConstructorCallbackEnabled() bool
	// CreateConstructorCallback generates the function to be called whan
	// JavaScript code constructs an instance.
	CreateConstructorCallbackBody(ESConstructorData, CallbackContext) Generator
	CreateIllegalConstructorCallback(ESConstructorData, CallbackContext) Generator

	CreateMethodCallbackBody(ESConstructorData, ESOperation, CallbackContext) Generator
	CreateAttributeGetter(
		ESConstructorData,
		ESOperation,
		CallbackContext,
		func(Generator) Generator,
	) Generator
	CreateAttributeSetter(
		ESConstructorData,
		ESOperation,
		CallbackContext,
		func(Generator, Generator) Generator,
	) Generator
	WrapperStructGenerators() PlatformWrapperStructGenerators

	// Generate a return statement with an error messages. Goja is
	// non-idiomatic, and the return value is converted to a panic. V8go is
	// idiomatic, and this generates a return nil, errors.New(msg)
	ReturnErrMsg(Generator) Generator
	PlatformInfoArg() Generator
}

// PrototypeWrapperGenerator generates code to create a JavaScript prototype
// that wraps an internal Go type.
type PrototypeWrapperGenerator struct {
	Platform TargetGenerators
	Data     ESConstructorData
}

func (g PrototypeWrapperGenerator) Generate() *jen.Statement {
	list := generators.StatementList()

	if !g.Data.IdlInterface.Mixin {
		list.Append(
			g.Platform.CreateInitFunction(g.Data),
			generators.Line,
		)
	}
	if !g.Data.Spec.SkipWrapper {
		list.Append(WrapperStructGenerator(g))
	}

	list.Append(
		g.Platform.CreateHostInitializer(g.Data),
		generators.Line,
		PrototypeInitializerGenerator(g),
		g.Constructor(),
		g.CreateOperationCallbacks(g.Data),
		g.CreateAttributeCallbacks(g.Data),
	)

	return list.Generate()
}

func (gen PrototypeWrapperGenerator) Constructor() g.Generator {
	if gen.Data.Spec.SkipConstructor || !gen.Platform.ConstructorCallbackEnabled() {
		return g.Noop
	}
	receiver := generators.Id("w")
	return g.StatementList(
		g.Line,
		MethodCallback{
			"Constructor", receiver, gen.Data, nil, gen.Platform,
			ConstructorCallbackBody{receiver, gen.Data, gen.Platform},
		},
	)
}

func (g PrototypeWrapperGenerator) CreateOperationCallbacks(data ESConstructorData) Generator {
	list := generators.StatementList()
	receiver := generators.Id("w")
	for op := range data.OperationCallbackInfos() {
		list.Append(
			generators.Line,
			MethodCallback{
				op.CallbackMethodName(),
				receiver,
				data,
				&op,
				g.Platform,
				MethodCallbackBody{receiver, data, op, g.Platform},
			},
		)
	}
	return list
}

func (gen PrototypeWrapperGenerator) CreateAttributeCallbacks(data ESConstructorData) Generator {
	list := generators.StatementList()
	receiver := generators.Id("w")
	for _, attr := range data.Attributes {
		if attr.Getter != nil && !attr.Getter.CustomImplementation {
			list.Append(
				generators.Line,
				MethodCallback{
					attr.Getter.CallbackMethodName(),
					receiver,
					data,
					attr.Getter,
					gen.Platform,
					AttributeGetterCallbackBody{
						receiver, data, *attr.Getter, gen.Platform,
					},
				},
			)
		}
		if attr.Setter != nil && !attr.Setter.CustomImplementation {
			list.Append(
				generators.Line,
				MethodCallback{
					attr.Setter.CallbackMethodName(),
					receiver,
					data,
					attr.Setter,
					gen.Platform,
					AttributeSetterCallbackBody{
						receiver, data, *attr.Setter, gen.Platform,
					},
				},
			)
		}
	}
	return list
}

type AttributeGetterCallback struct {
	data     ESConstructorData
	op       ESOperation
	platform TargetGenerators
	body     Generator
}

func (c AttributeGetterCallback) Generate() *jen.Statement {
	typeGenerators := c.platform.WrapperStructGenerators()
	receiver := generators.Id("w")
	return generators.FunctionDefinition{
		Receiver: generators.FunctionArgument{
			Name: receiver,
			Type: typeGenerators.WrapperStructType(c.data.Name()),
		},
		Name:     c.op.CallbackMethodName(),
		Args:     typeGenerators.CallbackMethodArgs(),
		RtnTypes: typeGenerators.CallbackMethodRetTypes(),
		Body:     c.body,
	}.Generate()

}

type CtxTransformer interface {
	TransformCtx(CallbackContext) Generator
}

type MethodCallback struct {
	name     string
	receiver Generator
	data     ESConstructorData
	op       *ESOperation
	platform TargetGenerators
	body     CtxTransformer
}

func renderIf(condition bool, gen g.Generator) g.Generator {
	if condition {
		return gen
	}
	return g.Noop
}

func (c MethodCallback) Generate() *jen.Statement {
	typeGenerators := c.platform.WrapperStructGenerators()
	cbCtx := NewCallbackContext(g.Id("cbCtx"))
	args := typeGenerators.CallbackMethodArgs()
	argIds := make([]Generator, len(args))
	if len(args) != 1 {
		panic("Arguments passed to ctx must be more flexible")
	}
	for i, arg := range args {
		argIds[i] = arg.Name
	}
	notImplemented := c.op != nil && c.op.NotImplemented

	return generators.FunctionDefinition{
		Receiver: generators.FunctionArgument{
			Name: c.receiver,
			Type: typeGenerators.WrapperStructType(c.data.Name()),
		},
		Name:     c.name,
		Args:     args,
		RtnTypes: typeGenerators.CallbackMethodRetTypes(),
		Body: g.StatementList(
			c.LogCall(cbCtx),
			renderIf(!notImplemented, g.StatementList(
				c.body.TransformCtx(cbCtx),
			)),
			renderIf(notImplemented, c.ReturnNotImplementedError(cbCtx)),
		),
	}.Generate()
}

func (c MethodCallback) ReturnNotImplementedError(cbCtx CallbackContext) g.Generator {
	var name string
	if c.op != nil {
		name = c.op.Name
	}
	errMsg := fmt.Sprintf(
		"%s.%s: Not implemented. Create an issue: %s",
		c.data.Name(), name, packagenames.ISSUE_URL,
	)
	// return c.platform.ReturnErrMsg(g.Lit(errMsg))
	return g.Return(
		cbCtx.ReturnWithError(
			g.NewValuePackage("New", "errors").Call(g.Lit(errMsg))))
}

func (c MethodCallback) LogCall(cbCtx g.Generator) g.Generator {
	return g.ValueOf(cbCtx).Field("Logger").Call().Field("Debug").Call(
		g.Lit(fmt.Sprintf("V8 Function call: %s.%s", c.data.Name(), c.name)))
	// return stdgen.LogDebug(
	// 	g.ValueOf(c.receiver).Field("logger").Call(c.platform.PlatformInfoArg()),
	// 	g.Lit(fmt.Sprintf("V8 Function call: %s.%s", c.data.Name(), c.name)))
}

type MethodCallbackBody struct {
	receiver g.Generator
	data     ESConstructorData
	op       ESOperation
	platform TargetGenerators
}

func (b MethodCallbackBody) TransformCtx(cbCtx CallbackContext) Generator {
	statements := g.StatementList()

	statements.Append(
		b.platform.CreateMethodCallbackBody(b.data, b.op, cbCtx),
	)
	return statements
}

type AttributeGetterCallbackBody struct {
	receiver g.Generator
	data     ESConstructorData
	op       ESOperation
	platform TargetGenerators
}

func (b AttributeGetterCallbackBody) TransformCtx(cbCtx CallbackContext) Generator {
	statements := g.StatementList()

	statements.Append(
		b.platform.CreateAttributeGetter(b.data, b.op, cbCtx,
			func(instance g.Generator) g.Generator {
				name := IdlNameToGoName(b.op.Name)
				field := g.ValueOf(instance).Field(name)
				if b.data.CustomRule.OutputType == customrules.OutputTypeStruct {
					return field
				} else {
					return field.Call()
				}
			}),
	)
	return statements
}

type AttributeSetterCallbackBody struct {
	receiver g.Generator
	data     ESConstructorData
	op       ESOperation
	platform TargetGenerators
}

func (b AttributeSetterCallbackBody) TransformCtx(cbCtx CallbackContext) Generator {
	statements := g.StatementList()

	statements.Append(
		b.platform.CreateAttributeSetter(
			b.data, b.op, cbCtx,
			func(instance g.Generator, val g.Generator) g.Generator {
				name := IdlNameToGoName(b.op.Name)
				field := g.ValueOf(instance).Field(name)
				if b.data.CustomRule.OutputType == customrules.OutputTypeStruct {
					return g.Reassign(field, val)
				} else {
					return field.Call(val)
				}
			},
		),
	)
	return statements
}

type ConstructorCallbackBody struct {
	receiver g.Generator
	data     ESConstructorData
	platform TargetGenerators
}

func (b ConstructorCallbackBody) TransformCtx(ctx CallbackContext) Generator {
	if b.data.AllowConstructor() {
		return b.platform.CreateConstructorCallbackBody(b.data, ctx)
	} else {
		return g.Return(ctx.IllegalConstructor())
	}
}
