package wrappers

import "fmt"

type V8WrapperStructGenerators struct{}

func (g V8WrapperStructGenerators) WrapperStructTypeName(interfaceName string) string {
	return fmt.Sprintf("%sV8Wrapper", lowerCaseFirstLetter(interfaceName))
}

func (g V8WrapperStructGenerators) EmbedName(data ESConstructorData) string {
	return "nodeV8WrapperBase"
}
