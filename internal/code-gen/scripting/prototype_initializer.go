package scripting

import (
	"fmt"

	"github.com/dave/jennifer/jen"
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
	init := PrototypeInitializer{i.WrapperStruct}
	return g.StatementList(
		g.Raw(jen.Func().
			Add(jen.Id(i.InitializerName())).
			Add(jen.Index(jen.Id("T").Id("any"))).
			Params(jen.Id("jsClass").Add(jsClass.Generate())).
			Block(init.Body().Generate()),
		),
	).Generate()
}

func (i PrototypeInitializer) Body() g.Generator {
	return g.StatementList(
		i.CreatePrototypeInitializerBody(),
		i.MixinsGenerator(),
	)
}

func (i PrototypeInitializer) CreatePrototypeInitializerBody() g.Generator {
	class := class{g.NewValue("jsClass")}
	return g.StatementList(
		i.InstallFunctionHandlers(class),
		i.InstallAttributeHandlers(class),
		renderIf(i.Data.RunCustomCode,
			g.NewValue(fmt.Sprintf("%sCustomInitializer", i.IdlName())).Call(class),
		),
	)
}

func (i PrototypeInitializer) MixinsGenerator() g.Generator {
	result := g.StatementList()
	for _, mixin := range i.Data.Includes() {
		// Note: This excludes mixins not known in this spec.
		// Could the mixing come from another spec, then this will skip
		// something we may want.
		// Not a problem yet ...
		if _, included := i.Data.Spec.DomSpec.Interfaces[mixin.Name]; included {
			result.Append(
				g.NewValue(fmt.Sprintf("Initialize%s", mixin.Name)).Call(g.Id("jsClass")),
			)
		}
	}
	return result
}
func (i PrototypeInitializer) InstallFunctionHandlers(class class) g.Generator {
	renderedAny := false
	stmts := g.StatementList()
	for op := range i.Data.WrapperFunctionsToInstall() {
		cb := g.Id(op.CallbackMethodName())
		stmts.Append(class.CreateOperation(op.Name, cb))
		renderedAny = true
	}
	if renderedAny {
		return stmts
	} else {
		return g.Noop
	}
}

func (i PrototypeInitializer) InstallAttributeHandlers(class class) g.Generator {
	stmts := g.StatementList()
	for op := range i.Data.AttributesToInstall() {
		stmts.Append(i.InstallAttributeHandler(op, class))
	}
	return stmts
}

func (i PrototypeInitializer) attributeOptions(attr model.ESAttribute) []g.Generator {
	var res []g.Generator
	if attr.Spec.LegacyUnforgeable {
		res = append(res, jsLegacyUnforgeable())

	}
	return res
}

func (i PrototypeInitializer) InstallAttributeHandler(
	op model.ESAttribute,
	class class,
) g.Generator {
	getter := op.Getter
	setter := op.Setter
	if getter == nil {
		return g.Noop
	}
	getterFn := g.Id(getter.CallbackMethodName())
	setterFn := g.Nil
	if setter != nil {
		setterFn = g.Id(setter.CallbackMethodName())
	}

	res := g.StatementList(
		class.CreateAttribute(op.Name, getterFn, setterFn, i.attributeOptions(op)...),
	)
	if op.Spec.Stringifier {
		res.Append(
			class.CreateOperation("toString", getterFn),
		)
	}
	return res
}
