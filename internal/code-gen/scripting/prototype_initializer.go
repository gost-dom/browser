package scripting

import (
	"github.com/dave/jennifer/jen"
	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
)

// PrototypeInitializer generates a function to initialize a JavaScript
// prototype. Initialize means setting content attributes for functions, and
// accessor attributes for IDL attributes.
type PrototypeInitializer struct {
	WrapperStruct
}

func (i PrototypeInitializer) Generate() *jen.Statement {
	receiver := g.NewValue("w")
	class := g.NewValue("jsClass")
	wrapperType := i.WrapperStructType()
	return g.FunctionDefinition{
		Name:     "installPrototype",
		Receiver: g.FunctionArgument{Name: receiver, Type: wrapperType},
		Args:     g.Arg(class, jsClass),
		Body:     i.Body(receiver),
	}.Generate()
}

func (i PrototypeInitializer) Body(receiver g.Value) g.Generator {
	return g.StatementList(
		i.CreatePrototypeInitializerBody(receiver),
		i.MixinsGenerator(),
	)
}

func (i PrototypeInitializer) CreatePrototypeInitializerBody(
	receiver g.Value,
) g.Generator {
	class := class{g.NewValue("jsClass")}
	return g.StatementList(
		i.InstallFunctionHandlers(receiver, class),
		i.InstallAttributeHandlers(receiver, class),
	)
}

func (i PrototypeInitializer) MixinsGenerator() g.Generator {
	result := g.StatementList()
	for _, mixin := range i.Data.Includes() {
		wrapperName := LowerCaseFirstLetter(mixin.Name)
		// Note: This excludes mixins not known in this spec.
		// Could the mixing come from another spec, then this will skip
		// something we may want.
		// Not a problem yet ...
		if _, included := i.Data.Spec.DomSpec.Interfaces[mixin.Name]; included {
			result.Append(
				g.NewValue("w").Field(wrapperName).Field("installPrototype").Call(
					g.Id("jsClass")),
			)
		}
	}
	return result
}
func (i PrototypeInitializer) InstallFunctionHandlers(
	receiver g.Value, class class,
) g.Generator {
	renderedAny := false
	stmts := g.StatementList()
	for op := range i.Data.WrapperFunctionsToInstall() {
		cb := receiver.Field(op.CallbackMethodName())
		stmts.Append(class.CreateOperation(op.Name, cb))
		renderedAny = true
	}
	if renderedAny {
		return stmts
	} else {
		return g.Noop
	}
}

func (i PrototypeInitializer) InstallAttributeHandlers(
	receiver g.Value,
	class class,
) g.Generator {
	stmts := g.StatementList()
	for op := range i.Data.AttributesToInstall() {
		stmts.Append(i.InstallAttributeHandler(op, receiver, class))
	}
	return stmts
}

func (i PrototypeInitializer) InstallAttributeHandler(
	op model.ESAttribute,
	receiver g.Value,
	class class,
) g.Generator {
	getter := op.Getter
	setter := op.Setter
	if getter == nil {
		return g.Noop
	}
	getterFn := receiver.Field(getter.CallbackMethodName())
	setterFn := g.Nil
	if setter != nil {
		setterFn = receiver.Field(setter.CallbackMethodName())
	}

	res := g.StatementList(
		class.CreateAttribute(op.Name, getterFn, setterFn),
	)
	if op.Spec.Stringifier {
		res.Append(
			class.CreateOperation("toString", getterFn),
		)
	}
	return res
}
