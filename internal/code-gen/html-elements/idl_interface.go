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
		fields = append(fields, o)
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

// Represents an operation specified on an IDL interface. The type is itself a
// [generators.Generator] for generating the method, potentially multiple
// methods if the method is overloaded in the source.
type IdlInterfaceOperation struct {
	IdlOperation idl.Operation
	Arguments    []IdlInterfaceOperationArgument
	ReturnType   IdlType
	Rules        customrules.OperationRule
}

func (o IdlInterfaceOperation) Stringifier() bool { return o.IdlOperation.Stringifier }
func (o IdlInterfaceOperation) Name() string      { return o.IdlOperation.Name }
func (o IdlInterfaceOperation) Static() bool      { return o.IdlOperation.Static }

func (o IdlInterfaceOperation) Generate() *jen.Statement {
	if o.Stringifier() && o.Name() == "" {
		return nil
	}
	name := o.Name()
	opRules := o.Rules
	result := generators.StatementList()
	if !o.Static() {
		args := make([]generators.Generator, 0, len(o.Arguments))
		for i, a := range o.Arguments {
			if a.Ignore() {
				continue
			}
			var arg generators.Generator = IdlType(a.Type())
			if a.Rules.OverridesType() {
				arg = IdlType(a.Rules.Type)
			}
			if a.Variadic() {
				arg = generators.Raw(
					jen.Id(a.Name()).Add(jen.Op("...").Add(arg.Generate())),
				)
			}
			args = append(args, arg)

			if i < len(o.Arguments)-1 {
				nextArg := o.Arguments[i+1]
				if nextArg.Optional() {
					result.Append(generators.Raw(
						jen.Id(UpperCaseFirstLetter(name)).
							Params(generators.ToJenCodes(args)...).
							Add(o.ReturnParams())))
					name = name + UpperCaseFirstLetter(nextArg.Name())
				}
			}
		}

		if opRules.DocComments != "" {
			result.Append(generators.Raw(jen.Comment(opRules.DocComments)))
		}
		result.Append(generators.Raw(
			jen.Id(UpperCaseFirstLetter(name)).
				Params(generators.ToJenCodes(args)...).
				Add(o.ReturnParams()),
		))
	}
	return result.Generate()
}

func (o IdlInterfaceOperation) HasError() bool {
	return o.Rules.HasError
}

// ReturnParams return multiple parameters for an operation's return types.
// The return values can include a bool for methods like GetAttribute, that
// return (string, bool), indicating if the attribute was found. If hasError is
// true, an error return type will be added as well.
func (o IdlInterfaceOperation) ReturnParams() *jen.Statement {
	result := make([]generators.Generator, 1, 3)
	if o.Rules.ReturnType != nil {
		result[0] = o.Rules.ReturnType
	} else {
		s := o.ReturnType
		result[0] = s
		if s.Nullable && !s.Nillable() {
			result = append(result, generators.Id("bool"))
		}
		if o.HasError() {
			result = append(result, generators.Id("error"))
		}
	}
	return jen.Params(generators.ToJenCodes(result)...)
}

/* -------- IdlInterfaceInclude -------- */

type IdlInterfaceInclude struct{ idl.Interface }

type IdlInterfaceOperationArgument struct {
	Argument idl.Argument
	Rules    customrules.ArgumentRule
}

func (a IdlInterfaceOperationArgument) Name() string   { return a.Argument.Name }
func (a IdlInterfaceOperationArgument) Type() idl.Type { return a.Argument.Type }
func (a IdlInterfaceOperationArgument) Variadic() bool { return a.Argument.Variadic }
func (a IdlInterfaceOperationArgument) Optional() bool { return a.Argument.Optional }
func (a IdlInterfaceOperationArgument) Ignore() bool   { return a.Rules.Ignore }
