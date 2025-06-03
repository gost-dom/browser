package v8gen

import (
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
	v8Iso
	Proto   v8PrototypeTemplate
	Wrapper WrapperInstance
}

func (builder PrototypeInstaller) InstallFunctionHandlers(
	host g.Generator,
	data ESConstructorData,
) g.Generator {
	generators := make([]g.Generator, 0, len(data.Operations))
	for _, op := range data.Operations {
		if !op.MethodCustomization.Ignored {
			generators = append(generators,
				builder.InstallFunction(host, op.Name, op.CallbackMethodName()),
			)
		}
	}
	return g.StatementList(generators...)
}

func (builder PrototypeInstaller) InstallFunction(
	host g.Generator,
	name, cbMethod string,
) g.Generator {
	return builder.Proto.Set(name, wrapCallback(host, builder.Wrapper.Field(cbMethod)))
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
	generators[0] = g.Line
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
