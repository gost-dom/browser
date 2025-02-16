package wrappers

import (
	"fmt"

	"github.com/gost-dom/generators"
)

type V8WrapperStructGenerators struct{}

func (g V8WrapperStructGenerators) WrapperStructType(interfaceName string) Generator {
	return generators.Id(fmt.Sprintf("%sV8Wrapper", lowerCaseFirstLetter(interfaceName)))
}

func (g V8WrapperStructGenerators) WrapperStructConstructorName(interfaceName string) string {
	return fmt.Sprintf("new%sV8Wrapper", interfaceName)
}

func (g V8WrapperStructGenerators) WrapperStructConstructorRetType(
	idlInterfaceName string,
) Generator {
	return generators.Type{Generator: g.WrapperStructType(idlInterfaceName)}.Pointer()
}

func (g V8WrapperStructGenerators) EmbeddedType(wrappedType Generator) Generator {
	return generators.NewType("nodeV8WrapperBase").TypeParam(wrappedType)
}

func (g V8WrapperStructGenerators) HostArg() Generator {
	return generators.Id("scriptHost")
}

func (g V8WrapperStructGenerators) HostType() Generator {
	return scriptHostPtr
}
