package htmlelements

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/idltransform"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/generators"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

// HTMLGeneratorReq specifies what to generate for a specific Web IDL spec. The
// name of the spec is in the SpecName field, and the interface to generate is in
// the Interface field. The Generate... fields specify what to generate.
//
// Note: As more needs for customisation arises, so will the Generate... fields
// also more likely become more complex.
//
// E.g., the URL type is loaded from url.json, and has the interface name URL in
// the idl, so this is specified by:
//
//	HTMLGeneratorReq {
//		InterfaceName: "URL",
//		SpecName:      "url",
//	}
//
// This doesn't include _where_ to generate, e.g., the file name to place the
// code in, and the package name it exists in.
type HTMLGeneratorReq struct {
	// TODO: Shouldn't we extract two separate types? InterfaceName and SpecName
	// are used to lookup the IDL specification; whereas the other properties
	// specifies what to generate. So there's a kind of pipline here
	// Read(intfName, specName) -> Generate(GenStruct, GenCon...) -> Generator

	InterfaceName          string
	SpecName               string
	GenerateStruct         bool
	GenerateConstructor    bool
	GenerateInterface      bool
	GenerateReadonlyStruct bool
	GenerateAttributes     bool
}

/* -------- baseGenerator -------- */

type baseGenerator struct {
	req     HTMLGeneratorReq
	idlType idl.Interface
	type_   g.Type
	rules   customrules.InterfaceRule
}

func CreateGenerator(req HTMLGeneratorReq) (baseGenerator, error) {
	html, err := idl.Load(req.SpecName)
	specRules := customrules.GetSpecRules(req.SpecName)
	return baseGenerator{
		req,
		html.Interfaces[req.InterfaceName],
		g.NewType(toStructName(req.InterfaceName)),
		specRules[req.InterfaceName],
	}, err
}

func (gen baseGenerator) Generate() *jen.Statement {
	list := generators.StatementList()

	if gen.req.GenerateInterface {
		list.Append(gen.GenerateInterface())
	}

	if gen.req.GenerateReadonlyStruct || gen.rules.OutputType == customrules.OutputTypeStruct {
		list.Append(gen.GenerateReadonlyStruct())
	}

	return list.Generate()
}

func (gen baseGenerator) GenerateReadonlyStruct() g.Generator {
	idlInterfaceName := gen.idlType.Name
	if len(gen.idlType.Operations) > 0 {
		panic(
			fmt.Sprintf(
				"baseGenerator.CreateReadonlyStruct: IDL interface has operations; expected only readonly attributes. Interface: %s",
				idlInterfaceName,
			))
	}
	result := g.Struct{
		Name: g.NewType(idlInterfaceName),
	}
	for _, a := range gen.idlType.Attributes {
		if !a.Readonly {
			panic(
				fmt.Sprintf(
					"baseGenerator.CreateReadonlyStruct: IDL interface has writeable attributes. Interface: %s. Attribute: %s",
					idlInterfaceName,
					a.Name,
				),
			)
		}
		field := internal.UpperCaseFirstLetter(string(a.Name))
		result.Field(g.Id(field), idltransform.IdlTypeForStruct(a.Type))
	}
	return result
}

func (gen baseGenerator) GenerateInterface() g.Generator {
	idlInterface := gen.idlType
	attributes := make([]IdlInterfaceAttribute, 0)
	operations := make([]IdlInterfaceOperation, 0)
	includes := make([]IdlInterfaceInclude, len(idlInterface.Includes))
	iterableTypes := make([]idltransform.IdlType, len(idlInterface.IterableTypes))

	interfaces := make([]idl.Interface, 1+len(gen.idlType.Includes))
	interfaces[0] = gen.idlType
	copy(interfaces[1:], gen.idlType.Includes)
	result := IdlInterface{
		Name:     gen.idlType.Name,
		Inherits: gen.idlType.InternalSpec.Inheritance,
		Includes: includes,
		Rules:    gen.rules,
	}

	for idx, i := range gen.idlType.Includes {
		includes[idx] = IdlInterfaceInclude{i}
	}

	for _, a := range gen.idlType.Attributes {
		attributeRule := gen.rules.Attributes[a.Name]
		if attributeRule.NotImplemented {
			continue
		}
		if a.Stringifier {
			result.HasStringifier = true
		}
		attrType := a.Type
		if attributeRule.OverrideType != nil {
			attrType = attributeRule.OverrideType.IdlType()
		}
		attributes = append(attributes, IdlInterfaceAttribute{
			Name:     a.Name,
			Type:     idltransform.IdlType(attrType),
			ReadOnly: a.Readonly,
		})
	}
	for _, o := range gen.idlType.Operations {
		if o.Stringifier {
			result.HasStringifier = true
			if o.Name == "" {
				continue
			}
		}
		operationRule := gen.rules.Operations[o.Name]
		getArg := func(name string) (res customrules.ArgumentRule) {
			if operationRule.Arguments != nil {
				res = operationRule.Arguments[name]
			}
			return
		}
		arguments := make([]IdlInterfaceOperationArgument, len(o.Arguments))
		for i, arg := range o.Arguments {
			arguments[i] = IdlInterfaceOperationArgument{
				Argument: arg,
				Rules:    getArg(arg.Name),
			}
		}
		operations = append(
			operations,
			IdlInterfaceOperation{o, arguments, idltransform.IdlType(o.ReturnType), operationRule},
		)
	}
	for i, t := range idlInterface.IterableTypes {
		iterableTypes[i] = idltransform.IdlType(t)
	}
	result.Attributes = attributes
	result.Operations = operations
	result.IterableTypes = iterableTypes
	return result
}

/* -------- htmlElementGenerator -------- */

type htmlElementGenerator struct {
	baseGenerator
	tagName string
}

func (gen htmlElementGenerator) Generator() g.Generator {
	result := g.StatementList()
	if gen.req.GenerateInterface {
		result.Append(
			gen.GenerateInterface(),
			g.Line,
		)
	}
	if gen.req.GenerateStruct {
		result.Append(gen.GenerateStruct(),
			g.Line,
		)
	}
	if gen.req.GenerateConstructor {
		result.Append(
			gen.GenerateConstructor(),
			g.Line,
		)
	}
	if gen.req.GenerateAttributes {
		result.Append(gen.GenerateAttributes())
	}
	return result
}

func toStructName(name string) string {
	return strings.Replace(name, "HTML", "html", 1)
}

func (gen htmlElementGenerator) GenerateStruct() g.Generator {
	res := g.Struct{Name: g.NewType(toStructName(gen.idlType.Name))}
	res.Embed(g.Id("HTMLElement"))
	// for a := range gen.idlType.Attributes() {
	// 	res.Field(g.Id(idl.SanitizeName(a.Name)), g.Id("string"))
	// }
	return res
}

func (gen htmlElementGenerator) GenerateConstructor() g.Generator {
	res := g.NewValue("result")
	i := g.NewType(gen.idlType.Name)
	t := g.NewType(toStructName(gen.idlType.Name))
	owner := g.Id("ownerDoc")
	return g.FunctionDefinition{
		Name:     fmt.Sprintf("New%s", gen.idlType.Name),
		RtnTypes: g.List(i),
		Args:     g.Arg(owner, g.Id("HTMLDocument")),
		Body: g.StatementList(
			g.Assign(
				res,
				t.CreateInstance(g.NewValue("NewHTMLElement").Call(g.Lit(gen.tagName), owner)).
					Reference(),
			),
			res.Field("SetSelf").Call(res),
			g.Return(res),
		),
	}
}

func (gen htmlElementGenerator) GenerateAttributes() g.Generator {
	result := g.StatementList()
	for _, a := range gen.idlType.Attributes {
		result.Append(IDLAttribute{
			AttributeName: a.Name,
			ReadOnly:      a.Readonly,
			Receiver: Receiver{
				Name: g.NewValue("e"),
				Type: gen.type_.Pointer(),
			},
		})
	}
	return result
}

type FileGeneratorSpec struct {
	OutputFile string
	Package    string
	Generator  g.Generator
}

// GeneratorConfig contains the configuration for which generated files should contain
// which interfaces. The key is a base file name. The system will append
// "_generated.go" to the name before creating the file. The HTMLGeneratorReq
// specifies the IDL source type, as well as what to generate.
type GeneratorConfig map[string]HTMLGeneratorReq

func createGenerators(
	config map[string]HTMLGeneratorReq,
	packageName string,
) ([]FileGeneratorSpec, error) {
	result := make([]FileGeneratorSpec, len(config))
	errs := make([]error, len(config))
	index := 0
	for k, v := range config {
		generator, err := CreateGenerator(v)
		result[index] = FileGeneratorSpec{
			k,
			packageName,
			generator,
		}
		errs[index] = err
		index++
	}
	return result, errors.Join(errs...)
}

func setSpecName(name string, data GeneratorConfig) GeneratorConfig {
	for k, v := range data {
		v.SpecName = name
		data[k] = v
	}
	return data
}
