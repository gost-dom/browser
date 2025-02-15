package htmlelements

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

/* -------- IdlInterface -------- */

type IdlInterface struct {
	Name       string
	Inherits   string
	Rules      customrules.InterfaceRule
	Attributes []IdlInterfaceAttribute
	Operations []IdlInterfaceOperation
}

func (i IdlInterface) Generate() *jen.Statement {
	fields := make(
		[]generators.Generator,
		0,
		2*len(i.Attributes)+1,
	) // Make room for getters and setters
	if i.Inherits != "" {
		fields = append(fields, generators.Id(i.Inherits))
	}

	for _, a := range i.Attributes {
		getterName := upperCaseFirstLetter(a.Name)
		fields = append(fields, generators.Raw(
			jen.Id(getterName).Params().Params(a.Type.Generate()),
		))
		if !a.ReadOnly {
			setterName := fmt.Sprintf("Set%s", getterName)
			fields = append(fields, generators.Raw(
				jen.Id(setterName).Params(a.Type.Generate()),
			))
		}
	}
	for _, o := range i.Operations {
		if !o.Static {
			// Todo: Parameters
			// Todo: Customize presence of error
			args := make([]generators.Generator, len(o.Arguments))
			for i, a := range o.Arguments {
				args[i] = IdlType(a.Type)
			}

			var returnTypes *jen.Statement
			if o.HasError {
				returnTypes = jen.Params(o.ReturnType.Generate(), jen.Id("error"))
			} else {
				returnTypes = jen.Params(o.ReturnType.Generate())
			}

			fields = append(fields, generators.Raw(
				jen.Id(upperCaseFirstLetter(o.Name)).
					Params(generators.ToJenCodes(args)...).
					Add(returnTypes),
			))
		}
	}
	return jen.Type().Add(jen.Id(i.Name)).Interface(generators.ToJenCodes(fields)...)
}

/* -------- IdlInterfaceAttribute -------- */

type IdlInterfaceAttribute struct {
	Name     string
	Type     IdlType
	ReadOnly bool
}

/* -------- IdlInterfaceOperation -------- */

type IdlInterfaceOperation struct {
	idl.Operation
	ReturnType IdlType
	HasError   bool
}
