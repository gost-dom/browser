package v8gen

import (
	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/script-wrappers/model"
	. "github.com/gost-dom/code-gen/script-wrappers/model"
	g "github.com/gost-dom/generators"
)

// The ConstructorBuilder is the function that creates the ES constructor
// itself, i.e. starts with a new function template, installs prototypes on the
// template, etc.
type ConstructorBuilder struct {
	v8Iso
	FT           v8FunctionTemplate
	Proto        v8PrototypeTemplate
	InstanceTmpl v8InstanceTemplate
	Wrapper      WrapperInstance
}

func NewConstructorBuilder() ConstructorBuilder {
	return ConstructorBuilder{
		v8Iso:        v8Iso{g.NewValue("iso")},
		FT:           v8FunctionTemplate{g.NewValue("ft")},
		Proto:        v8PrototypeTemplate{g.NewValue("prototypeTmpl")},
		InstanceTmpl: v8InstanceTemplate{g.NewValue("instanceTmpl")},
		Wrapper:      WrapperInstance{g.NewValue("wrapper")},
	}
}

func (builder ConstructorBuilder) NewFunctionTemplateOfWrappedMethod(name string) g.Generator {
	return builder.NewFunctionTemplate(builder.Wrapper.Method(name))
}

type PrototypeInstaller struct {
	Ft       v8FunctionTemplate
	Proto    v8PrototypeTemplate
	Wrapper  WrapperInstance
	Host     g.Generator
	Data     model.ESConstructorData
	Receiver g.Value
}

func (i PrototypeInstaller) Generate() *jen.Statement {
	class := jsClass(g.NewValue("jsClass"))
	return g.StatementList(
		i.InstallFunctionHandlers(i.Data, class),
		i.InstallAttributeHandlers(i.Host, i.Data),
	).Generate()
}

func (b PrototypeInstaller) InstallFunctionHandlers(
	data ESConstructorData, class jsClass,
) g.Generator {
	renderedAny := false
	stmts := g.StatementList(
		g.Assign(class, newJSClass(b.Host, b.Ft)),
	)
	for _, op := range data.Operations {
		if op.MethodCustomization.Ignored {
			continue
		}
		cb := b.Receiver.Field(op.CallbackMethodName())
		stmts.Append(class.CreatePrototypeMethod(op.Name, cb))
		renderedAny = true
	}
	if renderedAny {
		return stmts
	} else {
		return g.Noop
	}
}

func wrapCallback(host, callback g.Generator) g.Generator {
	return g.NewValue("wrapV8Callback").Call(host, callback)
}

func (builder PrototypeInstaller) InstallAttributeHandlers(
	host g.Generator,
	data ESConstructorData,
) g.Generator {
	length := len(data.Attributes)
	if length == 0 {
		return g.Noop
	}
	generators := make([]g.Generator, 1, length+1)
	generators[0] = g.Assign(g.NewValue("prototypeTmpl"), builder.Ft.GetPrototypeTemplate())
	for op := range data.AttributesToInstall() {
		generators = append(generators, builder.InstallAttributeHandler(host, op))
	}
	return g.StatementList(generators...)
}

func (builder PrototypeInstaller) InstallAttributeHandler(
	host g.Generator,
	op ESAttribute,
) g.Generator {
	wrapper := builder.Wrapper
	getter := op.Getter
	setter := op.Setter
	if getter == nil {
		return g.Noop
	}
	getterFt := wrapCallback(host, wrapper.Field(getter.CallbackMethodName()))
	setterFt := g.Nil
	if setter != nil {
		setterFt = wrapCallback(host, wrapper.Field(setter.CallbackMethodName()))
	}

	generator := builder.Proto.SetAccessorProperty(
		op.Name,
		g.WrapLine(getterFt),
		g.WrapLine(setterFt),
		g.WrapLine(v8None),
	)
	if op.Spec.Stringifier {
		return g.StatementList(generator,
			builder.InstallFunction(host, "toString", getter.CallbackMethodName()),
		)
	} else {
		return generator
	}
}
