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
	target  string
	idlType idl.Interface
	type_   g.Type
	rules   customrules.InterfaceRule
}

func (g baseGenerator) newIdlType(t idl.Type) idltransform.IdlType {
	return idltransform.IdlType{Type: t, TargetPackage: g.target}
}

func CreateGenerator(req HTMLGeneratorReq, target string) (baseGenerator, error) {
	html, err := idl.Load(req.SpecName)
	specRules := customrules.GetSpecRules(req.SpecName)
	return baseGenerator{
		req,
		target,
		html.Interfaces[req.InterfaceName],
		g.NewType(toStructName(req.InterfaceName)),
		specRules[req.InterfaceName],
	}, err
}

func (gen baseGenerator) Generate() *jen.Statement {
	list := generators.StatementList()

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
		result.Field(
			g.Id(field),
			idltransform.StructFieldType(gen.newIdlType(a.Type)),
		)
	}
	return result
}

/* -------- htmlElementGenerator -------- */

func toStructName(name string) string {
	return strings.Replace(name, "HTML", "html", 1)
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
		generator, err := CreateGenerator(v, packageName)
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
