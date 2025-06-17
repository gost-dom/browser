package htmlelements

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/customrules/typerule"
	"github.com/gost-dom/code-gen/idltransform"
	. "github.com/gost-dom/code-gen/internal"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

/* -------- IdlInterface -------- */

type IdlInterface struct {
	SpecName       string
	Name           string
	Inherits       string
	HasStringifier bool // Whether the interface contains a stringifier
	Rules          customrules.InterfaceRule
	Attributes     []IdlInterfaceAttribute
	Operations     []IdlInterfaceOperation
	Includes       []IdlInterfaceInclude
	IterableTypes  []idltransform.IdlType
	TargetPkg      string
}

func (i IdlInterface) Generate() *jen.Statement {
	fields := make(
		[]g.Generator,
		0,
		2*len(i.Attributes)+1, // Make room for getters and setters
	)
	if i.Inherits != "" {
		fields = append(fields, idltransform.TypeGen(i.Inherits, i.TargetPkg))
	}

	for _, incl := range i.Includes {
		fields = append(fields, g.Id(incl.Name))
	}

	if i.HasStringifier {
		fields = append(fields, g.NewTypePackage("Stringer", "fmt"))
	}

	for _, a := range i.Attributes {
		getterName := UpperCaseFirstLetter(a.Name)
		fields = append(fields, InterfaceFunction{
			Name:     getterName,
			RetTypes: g.List(a.GetterType(i.SpecName)),
		})
		if !a.ReadOnly {
			setterName := fmt.Sprintf("Set%s", getterName)
			fields = append(fields, InterfaceFunction{
				Name: setterName,
				Args: g.List(a.Type),
			})
		}
	}
	for _, o := range i.Operations {
		fields = append(fields, o)
	}
	switch len(i.IterableTypes) {
	case 0:
		{
		}
	case 1:
		fields = append(fields, iterator(i.IterableTypes[0]))
	case 2:
		fields = append(fields, iterator2(i.IterableTypes[0], i.IterableTypes[1]))
	default:
		panic("codegen: more than two iterable types found. Web idl spec allow up to two")
	}
	return jen.Type().Add(jen.Id(i.Name)).Interface(g.ToJenCodes(fields)...)
}

func iterator(t idltransform.IdlType) g.Generator {
	return g.Raw(
		jen.Id("All").Params().Qual("iter", "Seq2").Index(t.Generate()),
	)
}
func iterator2(k, v idltransform.IdlType) g.Generator {
	return g.Raw(
		jen.Id("All").Params().Qual("iter", "Seq2").Types(k.Generate(), v.Generate()),
	)
}

/* -------- IdlInterfaceAttribute -------- */

type IdlInterfaceAttribute struct {
	Name     string
	Type     idltransform.IdlType
	ReadOnly bool
}

func (a IdlInterfaceAttribute) GetterType(spec string) g.Generator {
	rules := customrules.GetSpecRules(spec)
	if rule, found := rules[a.Type.Name]; found {
		if rule.OutputType == customrules.OutputTypeStruct {
			return g.Raw(jen.Op("*").Add(a.Type.Generate()))
		}
	}
	return a.Type
}

/* -------- IdlInterfaceOperation -------- */

// OutputType describes a type to generate. The Default value is from Web IDL
// specifications, and the Override is custom configuration.
type OutputType struct {
	Default  idltransform.IdlType
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
	ReturnType   idltransform.IdlType
	Rules        customrules.OperationRule
	Target       string
}

func (o IdlInterfaceOperation) Stringifier() bool { return o.IdlOperation.Stringifier }
func (o IdlInterfaceOperation) Name() string      { return o.IdlOperation.Name }
func (o IdlInterfaceOperation) Static() bool      { return o.IdlOperation.Static }

func (o IdlInterfaceOperation) newIdlType(t idl.Type) idltransform.IdlType {
	return idltransform.IdlType{Type: t, TargetPackage: o.Target}
}

func (o IdlInterfaceOperation) Generate() *jen.Statement {
	if o.Stringifier() && o.Name() == "" || o.Static() {
		return nil
	}
	opRules := o.Rules
	name := o.Name()
	result := g.StatementList()
	args := make([]g.Generator, 0, len(o.Arguments))
	for i, a := range o.Arguments {
		if a.Ignore() {
			continue
		}
		var arg g.Generator = o.newIdlType(a.Type())
		if a.Rules.OverridesType() {
			arg = o.newIdlType(a.Rules.Type)
		}
		if a.Variadic() {
			arg = g.Raw(jen.Op("...").Add(arg.Generate()))
		}
		args = append(args, arg)

		if i < len(o.Arguments)-1 {
			nextArg := o.Arguments[i+1]
			argRule := o.Rules.Arguments[nextArg.Name()]
			if nextArg.Optional() && !argRule.ZeroAsDefault {
				result.Append(InterfaceFunction{
					Name:     UpperCaseFirstLetter(name),
					Args:     args,
					RetTypes: o.ReturnParams(),
				})
				name = name + UpperCaseFirstLetter(nextArg.Name())
			}
		}
	}

	if opRules.DocComments != "" {
		result.Append(g.Raw(jen.Comment(opRules.DocComments)))
	}
	result.Append(InterfaceFunction{
		Name:     UpperCaseFirstLetter(name),
		Args:     args,
		RetTypes: o.ReturnParams(),
	})

	return result.Generate()
}

func (o IdlInterfaceOperation) HasError() bool {
	return o.Rules.HasError
}

// ReturnParams return multiple parameters for an operation's return types.
// The return values can include a bool for methods like GetAttribute, that
// return (string, bool), indicating if the attribute was found. If hasError is
// true, an error return type will be added as well.
func (o IdlInterfaceOperation) ReturnParams() []g.Generator {
	result := make([]g.Generator, 1, 3)
	s := o.ReturnType
	result[0] = s
	if s.Nullable && !s.Nillable() {
		result = append(result, g.Id("bool"))
	}
	if o.HasError() {
		result = append(result, g.Id("error"))
	}
	return result
}

/* -------- IdlInterfaceInclude -------- */

type IdlInterfaceInclude struct{ idl.Interface }

type IdlInterfaceOperationArgument struct {
	Argument idl.Argument
	Rules    customrules.ArgumentRule
}

func (a IdlInterfaceOperationArgument) Name() string   { return a.Argument.Name }
func (a IdlInterfaceOperationArgument) Type() idl.Type { return a.Argument.Type }
func (a IdlInterfaceOperationArgument) Variadic() bool {
	return a.Argument.Variadic || a.Rules.Variadic
}

func (a IdlInterfaceOperationArgument) Optional() bool { return a.Argument.Optional && !a.Variadic() }
func (a IdlInterfaceOperationArgument) Ignore() bool   { return a.Rules.Ignore }

// InterfaceFunction renders a method in an interface definition.
type InterfaceFunction struct {
	Name     string
	Args     []g.Generator
	RetTypes []g.Generator
}

func (f InterfaceFunction) Generate() *jen.Statement {
	return jen.Id(f.Name).
		Params(g.ToJenCodes(f.Args)...).    // Arguments
		Params(g.ToJenCodes(f.RetTypes)...) // Return values
}
