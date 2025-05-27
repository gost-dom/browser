package v8gen

import (
	"fmt"

	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
)

type V8WrapperStructGenerators struct{}

func (g V8WrapperStructGenerators) WrapperStructType(interfaceName string) generators.Type {
	return generators.NewType(fmt.Sprintf("%sV8Wrapper", LowerCaseFirstLetter(interfaceName)))
}

func (g V8WrapperStructGenerators) WrapperStructConstructorName(interfaceName string) string {
	return fmt.Sprintf("new%sV8Wrapper", interfaceName)
}

func (g V8WrapperStructGenerators) WrapperStructConstructorRetType(
	idlInterfaceName string,
) g.Generator {
	return g.WrapperStructType(idlInterfaceName).Pointer()
}

func (g V8WrapperStructGenerators) EmbeddedType(wrappedType g.Generator) g.Generator {
	return generators.NewType("handleReffedObject").TypeParam(wrappedType)
}

func (g V8WrapperStructGenerators) EmbeddedTypeConstructor(
	wrappedType g.Generator,
) generators.Value {
	return generators.NewValue("newHandleReffedObject").TypeParam(wrappedType)
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
	return generators.List(g.NewType("jsValue"), g.NewType("error"))
}
