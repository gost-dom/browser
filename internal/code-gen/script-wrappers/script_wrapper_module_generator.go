package wrappers

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
)

type Generator = generators.Generator

// Create a struct definition, and it's constructor, that must contain the
// methods acting as callback for prototype functions, including the
// constructor.
type PlatformWrapperStructGenerators interface {
	WrapperStructType(interfaceName string) generators.Type
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
	// CreateConstructorCallback generates the function to be called whan
	// JavaScript code constructs an instance.
	CreateConstructorCallback(ESConstructorData) Generator

	CreateMethodCallbackBody(ESConstructorData, ESOperation) Generator
	WrapperStructGenerators() PlatformWrapperStructGenerators
	ReturnError(Generator) Generator
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
		PrototypeInitializerGenerator(g),
		g.Platform.CreateConstructorCallback(g.Data),
		g.MethodCallbacks(g.Data),
	)

	return list.Generate()
}

func (g PrototypeWrapperGenerator) MethodCallbacks(data ESConstructorData) Generator {
	list := generators.StatementList()
	for op := range data.WrapperFunctionsToGenerate() {
		list.Append(
			generators.Line,
			MethodCallback{data, op, g.Platform},
		)
	}
	return list
}

type MethodCallback struct {
	data     ESConstructorData
	op       ESOperation
	platform TargetGenerators
}

func (c MethodCallback) Generate() *jen.Statement {
	typeGenerators := c.platform.WrapperStructGenerators()
	receiver := generators.Id("w")
	return generators.FunctionDefinition{
		Receiver: generators.FunctionArgument{
			Name: receiver,
			Type: typeGenerators.WrapperStructType(c.data.Name()),
		},
		Name:     c.op.CallbackMethodName(),
		Args:     typeGenerators.CallbackMethodArgs(), // generators.Arg(generators.Id("info"), v8FunctionCallbackInfoPtr),
		RtnTypes: typeGenerators.CallbackMethodRetTypes(),
		Body:     MethodCallbackBody{c.data, c.op, receiver, c.platform},
	}.Generate()
}

type MethodCallbackBody struct {
	data     ESConstructorData
	op       ESOperation
	receiver g.Generator
	platform TargetGenerators
}

func (c MethodCallbackBody) Generate() *jen.Statement {
	debug := g.NewValuePackage("Debug", packagenames.Log).Call(
		g.ValueOf(c.receiver).Field("logger").Call(c.platform.PlatformInfoArg()),
		g.Lit(fmt.Sprintf("V8 Function call: %s.%s", c.data.Name(), c.op.Name)))

	if c.op.NotImplemented {
		errMsg := fmt.Sprintf(
			"%s.%s: Not implemented. Create an issue: %s",
			c.data.Name(),
			c.op.Name,
			packagenames.ISSUE_URL,
		)
		return g.StatementList(
			debug,
			c.platform.ReturnError(
				(g.Lit(errMsg)))).Generate()
		// jen.Qual("errors", "New").Call(jen.Lit(errMsg))))
		// g.Return(g.Nil, g.Raw(jen.Qual("errors", "New").Call(jen.Lit(errMsg))))).Generate()
	}
	return c.platform.CreateMethodCallbackBody(c.data, c.op).Generate()
}
