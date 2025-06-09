package scripting

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/configuration"
	. "github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

func writeGenerator(writer io.Writer, packagePath string, generator g.Generator) error {
	file := jen.NewFilePath(packagePath)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportAlias(packagenames.V8go, "v8")
	file.ImportAlias(packagenames.Goja, "g")
	file.Add(generator.Generate())
	return file.Render(writer)
}

type ScriptWrapperModulesGenerator struct {
	Specs            configuration.WebIdlConfigurations
	PackagePath      string
	TargetGenerators TargetGenerators
}

func (gen ScriptWrapperModulesGenerator) createPrototypeGenerator(
	typeGenerationInformation ESConstructorData,
) g.Generator {
	return PrototypeWrapperGenerator{
		Platform: gen.TargetGenerators,
		Data:     typeGenerationInformation,
	}
}

func (gen ScriptWrapperModulesGenerator) writeModuleTypes(spec *configuration.WebAPIConfig) error {
	data, err := idl.Load(spec.Name)
	if err != nil {
		return err
	}
	types := spec.GetTypesSorted()
	errs := make([]error, len(types))
	for i, specType := range types {
		outputFileName := fmt.Sprintf("%s_generated.go", typeNameToFileName(specType.TypeName))
		if writer, err := os.Create(outputFileName); err != nil {
			errs[i] = err
		} else {
			typeGenerationInformation := createData(data, specType)
			errs[i] = writeGenerator(writer, gen.PackagePath, gen.createPrototypeGenerator(typeGenerationInformation))
		}
	}
	return errors.Join(errs...)
}

var matchKnownWord = regexp.MustCompile("(HTML|URL|DOM)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func typeNameToFileName(name string) string {
	snake := matchKnownWord.ReplaceAllString(name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (gen ScriptWrapperModulesGenerator) GenerateScriptWrappers() error {
	specs := slices.Collect(maps.Values(gen.Specs))
	errs := make([]error, len(specs))
	for i, spec := range specs {
		fmt.Println("Generate module", spec.Name)
		errs[i] = gen.writeModuleTypes(spec)
	}
	return errors.Join(errs...)
}
