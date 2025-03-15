package htmlelements

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/customrules/typerule"
	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

/* -------- IdlInterface -------- */

type IdlInterface struct {
	Name           string
	Inherits       string
	HasStringifier bool
	Rules          customrules.InterfaceRule
	Attributes     []IdlInterfaceAttribute
	Operations     []IdlInterfaceOperation
	Includes       []IdlInterfaceInclude
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

	for _, i := range i.Includes {
		fields = append(fields, generators.Id(i.Name))
	}

	if i.HasStringifier {
		fields = append(fields, generators.NewTypePackage("Stringer", "fmt"))
	}

	for _, a := range i.Attributes {
		getterName := UpperCaseFirstLetter(a.Name)
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
		if o.Stringifier && o.Name == "" {
			continue
		}
		name := o.Name
		opRules := i.Rules.Operations[name]
		if !o.Static {
			args := make([]generators.Generator, len(o.Arguments))
			for i, a := range o.Arguments {
				argRules, hasArgRules := opRules.Arguments[a.Name]
				if hasArgRules {
					args[i] = IdlType(argRules.Type)
				} else {
					args[i] = IdlType(a.Type)
				}
				if a.Variadic {
					args[i] = generators.Raw(
						jen.Id(a.Name).Add(jen.Op("...").Add(args[i].Generate())),
					)
				}

				if i < len(o.Arguments)-1 {
					nextArg := o.Arguments[i+1]
					if nextArg.Optional {
						fields = append(fields, generators.Raw(
							jen.Id(UpperCaseFirstLetter(name)).
								Params(generators.ToJenCodes(args[0:i+1])...).
								Add(o.ReturnParams())))
						name = name + UpperCaseFirstLetter(nextArg.Name)
					}
				}
			}

			if opRules.DocComments != "" {
				fields = append(fields, generators.Raw(jen.Comment(opRules.DocComments)))
			}
			fields = append(fields, generators.Raw(
				jen.Id(UpperCaseFirstLetter(name)).
					Params(generators.ToJenCodes(args)...).
					Add(o.ReturnParams()),
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

// OutputType describes a type to generate. The Default value is from Web IDL
// specifications, and the Override is custom configuration.
type OutputType struct {
	Default  IdlType
	Override *typerule.TypeRule
}

func (t OutputType) Generate() *jen.Statement {
	if t.Override != nil {
		return t.Override.Generate()
	}
	return t.Default.Generate()
}

type IdlInterfaceOperation struct {
	idl.Operation
	ReturnType     IdlType
	ReturnOverride *typerule.TypeRule
	HasError       bool
}

// ReturnParams return multiple parameters for an operation's return types.
// The return values can include a bool for methods like GetAttribute, that
// return (string, bool), indicating if the attribute was found. If hasError is
// true, an error return type will be added as well.
func (o IdlInterfaceOperation) ReturnParams() *jen.Statement {
	result := make([]generators.Generator, 1, 3)
	if o.ReturnOverride != nil {
		result[0] = o.ReturnOverride
	} else {
		s := o.ReturnType
		result[0] = s
		if s.Nullable && !s.Nillable() {
			result = append(result, generators.Id("bool"))
		}
		if o.HasError {
			result = append(result, generators.Id("error"))
		}
	}
	return jen.Params(generators.ToJenCodes(result)...)
}

/* -------- IdlInterfaceInclude -------- */

type IdlInterfaceInclude struct{ idl.Interface }
