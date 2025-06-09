package scripting

import (
	"github.com/dave/jennifer/jen"
	. "github.com/gost-dom/code-gen/scripting/model"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
)

type Generator = generators.Generator

// PrototypeWrapperGenerator generates code to create a JavaScript prototype
// that wraps an internal Go type.
type PrototypeWrapperGenerator struct {
	Data ESConstructorData
}

func (gen PrototypeWrapperGenerator) Generate() *jen.Statement {
	wrapper := WrapperStruct(gen)

	return generators.StatementList(
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

func (gen PrototypeWrapperGenerator) OperationCallbacks() Generator {
	wrapper := WrapperStruct(gen)
	list := generators.StatementList()
	for op := range gen.Data.OperationCallbackInfos() {
		list.Append(
			generators.Line,
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
				generators.Line,
				callbacks.AttributeGetter(attr),
			)
		}
		if attr.Setter != nil && !attr.Setter.CustomImplementation {
			list.Append(
				generators.Line,
				callbacks.AttributeSetter(attr),
			)
		}
	}
	return list
}

type CtxTransformer interface {
	TransformCtx(CallbackContext) Generator
}
