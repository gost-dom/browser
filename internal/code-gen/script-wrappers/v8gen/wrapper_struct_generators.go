package v8gen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
)

type V8WrapperStructGenerators struct{}

func (g V8WrapperStructGenerators) WrapperStructTypeDef(interfaceName string) generators.Type {
	name := fmt.Sprintf("%sV8Wrapper", interfaceName)
	return generators.Type{
		Generator: generators.Raw(jen.Id(name).Types(jen.Id("T").Any())),
	}
}

func (g V8WrapperStructGenerators) WrapperStructTypeRetDef(interfaceName string) generators.Type {
	name := fmt.Sprintf("%sV8Wrapper", interfaceName)
	return generators.Type{
		Generator: generators.Raw(jen.Id(name).Types(jen.Id("T"))),
	}
}

func (g V8WrapperStructGenerators) WrapperStructType(interfaceName string) generators.Type {
	name := fmt.Sprintf("%sV8Wrapper", interfaceName)
	return generators.Type{
		Generator: generators.Raw(jen.Id(name).Types(jen.Id("T"))),
	}
}

func (g V8WrapperStructGenerators) WrapperStructConstructorName(interfaceName string) string {
	return fmt.Sprintf("New%sV8Wrapper", interfaceName)
}

func (g V8WrapperStructGenerators) WrapperStructConstructorRetType(
	idlInterfaceName string,
) g.Generator {
	name := fmt.Sprintf("%sV8Wrapper", idlInterfaceName)
	return generators.Type{
		Generator: generators.Raw(jen.Id(name).Types(jen.Id("T"))),
	}.Pointer()
}

func (gen V8WrapperStructGenerators) EmbeddedType(wrappedType g.Generator) g.Generator {
	return g.Noop
}

func (gen V8WrapperStructGenerators) EmbeddedTypeConstructor(
	wrappedType g.Generator,
) g.Generator {
	return nil
}

func (g V8WrapperStructGenerators) HostArg() g.Generator {
	return generators.Id("scriptHost")
}

func (g V8WrapperStructGenerators) HostType() g.Generator {
	return scriptHostPtr
}

func (g V8WrapperStructGenerators) CallbackMethodArgs() generators.FunctionArgumentList {
	return generators.Arg(generators.Id("cbCtx"), v8CbCtx)
}

func (gen V8WrapperStructGenerators) CallbackMethodRetTypes() []generators.Generator {
	return generators.List(jsValue, g.NewType("error"))
}
