package scripting

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/configuration"
	g "github.com/gost-dom/generators"
)

func writeGenerator(writer io.Writer, packagePath string, generator g.Generator) error {
	file := jen.NewFilePath(packagePath)
	file.HeaderComment("This file is generated. Do not edit.")
	file.Add(generator.Generate())
	return file.Render(writer)
}

func writePackageFiles(packagePath string, spec *configuration.WebAPIConfig) error {
	data, extra, err := configuration.LoadSpecs(spec)
	if err != nil {
		return err
	}
	types := spec.Types()
	errs := make([]error, len(types))
	for i, specType := range types {
		outputFileName := fmt.Sprintf(
			"%s_generated.go",
			internal.TypeNameToFileName(specType.TypeName),
		)
		if writer, err := os.Create(outputFileName); err != nil {
			errs[i] = err
		} else {
			defer writer.Close()
			typeGenerationInformation, err := createData(data, specType, extra)
			if err != nil {
				return err
			}
			gen := ScriptingFileGenerator{
				Data: typeGenerationInformation,
			}
			errs[i] = writeGenerator(writer, packagePath, gen)
		}
	}
	return errors.Join(errs...)
}

func CreateJavaScriptMappings(webAPI string) error {
	specs := slices.Collect(maps.Values(configuration.CreateV8SpecsForSpec(webAPI)))
	errs := make([]error, len(specs))
	for i, spec := range specs {
		fmt.Println("Generate module", spec.Name)
		errs[i] = writePackageFiles(packagenames.ScriptPackageName(webAPI), spec)
	}
	return errors.Join(errs...)
}
