package scripting

import (
	"github.com/dave/jennifer/jen"
	. "github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
)

// PrototypeWrapperGenerator generates code to create a JavaScript prototype
// that wraps an internal Go type.
type PrototypeWrapperGenerator struct {
	Data ESConstructorData
}

func (gen PrototypeWrapperGenerator) Generate() *jen.Statement {
	wrapper := WrapperStruct(gen)

	return g.StatementList(
		renderIf(!gen.Data.Spec.SkipWrapper, wrapper.TypeDef()),
		g.Line,
		HostInitializer{wrapper},
		g.Line,
		PrototypeInitializer{wrapper},
		g.Line,
		renderIf(!gen.Data.Spec.SkipConstructor, wrapper.Callbacks().ConstructorCallback()),
		gen.OperationCallbacks(),
		gen.AttributeCallbacks(),
	).Generate()
}

func (gen PrototypeWrapperGenerator) OperationCallbacks() g.Generator {
	wrapper := WrapperStruct(gen)
	list := g.StatementList()
	for op := range gen.Data.OperationCallbackInfos() {
		list.Append(
			g.Line,
			wrapper.Callbacks().MethodCallback(op),
		)
	}
	return list
}

func (gen PrototypeWrapperGenerator) AttributeCallbacks() g.Generator {
	list := g.StatementList()
	wrapper := WrapperStruct(gen)
	callbacks := wrapper.Callbacks()
	for _, attr := range gen.Data.Attributes {
		if attr.Getter != nil && !attr.Getter.CustomImplementation {
			list.Append(
				g.Line,
				callbacks.AttributeGetter(attr),
			)
		}
		if attr.Setter != nil && !attr.Setter.CustomImplementation {
			list.Append(
				g.Line,
				callbacks.AttributeSetter(attr),
			)
		}
	}
	return list
}
