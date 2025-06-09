package scripting

import (
	"github.com/dave/jennifer/jen"
	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/scripting/model"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
	gen "github.com/gost-dom/generators"
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
		Args:     g.Arg(class, v8Class),
		Body:     i.Body(receiver),
	}.Generate()
}

func (g PrototypeInitializer) Body(receiver g.Value) generators.Generator {
	return generators.StatementList(
		g.CreatePrototypeInitializerBody(receiver),
		g.MixinsGenerator(),
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

func (g PrototypeInitializer) MixinsGenerator() generators.Generator {
	result := gen.StatementList()
	for _, mixin := range g.Data.Includes() {
		wrapperName := LowerCaseFirstLetter(mixin.Name)
		// Note: This excludes mixins not known in this spec.
		// Could the mixing come from another spec, then this will skip
		// something we may want.
		// Not a problem yet ...
		if _, included := g.Data.Spec.DomSpec.Interfaces[mixin.Name]; included {
			result.Append(
				gen.NewValue("w").Field(wrapperName).Field("installPrototype").Call(
					gen.Id("jsClass")),
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
	for _, op := range i.Data.Operations {
		if op.MethodCustomization.Ignored {
			continue
		}
		cb := receiver.Field(op.CallbackMethodName())
		stmts.Append(class.CreatePrototypeMethod(op.Name, cb))
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
			class.CreatePrototypeMethod("toString", getterFn),
		)
	}
	return res
}
