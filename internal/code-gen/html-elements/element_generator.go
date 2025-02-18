package htmlelements

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gost-dom/code-gen/customrules"
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

	InterfaceName       string
	SpecName            string
	GenerateStruct      bool
	GenerateConstructor bool
	GenerateInterface   bool
	GenerateAttributes  bool
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

func (gen baseGenerator) GenerateInterface() g.Generator {
	idlInterface := gen.idlType
	attributes := make([]IdlInterfaceAttribute, 0)
	operations := make([]IdlInterfaceOperation, 0)
	includes := make([]IdlInterfaceInclude, len(idlInterface.Includes))

	interfaces := make([]idl.Interface, 1+len(gen.idlType.Includes))
	interfaces[0] = gen.idlType
	copy(interfaces[1:], gen.idlType.Includes)

	for idx, i := range gen.idlType.Includes {
		includes[idx] = IdlInterfaceInclude{i}
	}

	for _, a := range gen.idlType.Attributes {
		attributes = append(attributes, IdlInterfaceAttribute{
			Name:     a.Name,
			Type:     IdlType(a.Type),
			ReadOnly: a.Readonly,
		})
	}
	for _, o := range gen.idlType.Operations {
		operationRule := gen.rules.Operations[o.Name]
		operations = append(
			operations,
			IdlInterfaceOperation{o, IdlType(o.ReturnType), operationRule.HasError},
		)
	}
	// }
	return IdlInterface{
		Name:       gen.idlType.Name,
		Inherits:   gen.idlType.InternalSpec.Inheritance,
		Attributes: attributes,
		Operations: operations,
		Includes:   includes,
		Rules:      gen.rules,
	}
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
	Name      string
	Package   string
	Generator g.Generator
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
			generator.GenerateInterface(),
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
