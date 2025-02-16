package wrappers

import (
	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/generators"
)

type Generator = generators.Generator

type PlatformWrapperStructGenerators interface {
	WrapperStructType(interfaceName string) Generator
	WrapperStructConstructorName(interfaceName string) string
	WrapperStructConstructorRetType(interfaceName string) Generator
	EmbeddedType(wrappedType Generator) Generator
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
	// Create a struct definition, and it's constructor, that must contain the
	// methods acting as callback for prototype functions, including the
	// constructor.
	CreateWrapperStruct(ESConstructorData) Generator
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
	// CreateMethodCallback generates the function to be called when
	// JavaScript code calls a method on an instance.
	CreateMethodCallback(ESConstructorData, ESOperation) Generator

	WrapperStructGenerators() PlatformWrapperStructGenerators
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
		list.Append(g.Platform.CreateMethodCallback(data, op))
	}
	return list
}
