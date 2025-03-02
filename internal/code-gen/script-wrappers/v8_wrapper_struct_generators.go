package wrappers

import (
	"fmt"

	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/generators"
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
) Generator {
	return g.WrapperStructType(idlInterfaceName).Pointer()
}

func (g V8WrapperStructGenerators) EmbeddedType(wrappedType Generator) Generator {
	return generators.NewType("handleReffedObject").TypeParam(wrappedType)
}

func (g V8WrapperStructGenerators) EmbeddedTypeConstructor(wrappedType Generator) generators.Value {
	return generators.NewValue("newHandleReffedObject").TypeParam(wrappedType)
}

func (g V8WrapperStructGenerators) HostArg() Generator {
	return generators.Id("scriptHost")
}

func (g V8WrapperStructGenerators) HostType() Generator {
	return scriptHostPtr
}
