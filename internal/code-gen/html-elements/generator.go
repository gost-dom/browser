package htmlelements

import (
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/packagenames"
)

func writeFile(s FileGeneratorSpec) error {
	jf := jen.NewFilePath(s.Package)
	jf.HeaderComment("This file is generated. Do not edit.")
	jf.Add(s.Generator.Generate())
	outputFileName := fmt.Sprintf("%s_generated.go", s.OutputFile)
	if writer, err := os.Create(outputFileName); err != nil {
		return err
	} else {
		defer writer.Close()
		if err = jf.Render(writer); err != nil {
			return err
		}
	}
	return nil
}

func GetPackageGeneratorSpecs(packageName string) (result GeneratorConfig, err error) {
	result, found := PackageConfigs[packageName]
	if !found {
		err = fmt.Errorf("CreatePackageGenerators: No configuration for package %s", packageName)
	}
	return
}

func CreatePackageGenerators(name string) (res []FileGeneratorSpec, err error) {
	if config, err := GetPackageGeneratorSpecs(name); err == nil {
		packageName := fmt.Sprintf("%s/%s", packagenames.BASE_PKG, name)
		return createGenerators(config, packageName)
	}
	return
}

func CreateImplementationPackage(name string) error {
	files, err := CreatePackageGenerators(name)
	if err != nil {
		return err
	}
	for _, f := range files {
		if err = writeFile(f); err != nil {
			return err
		}
	}
	return nil
}

// func GenerateHTMLElements() error {
// 	files, err := CreateHTMLGenerators()
// 	if err != nil {
// 		return err
// 	}
// 	for _, f := range files {
// 		if err = writeFile(f); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func GenerateDOMTypes() error {
// 	files, err := CreateDOMGenerators()
// 	if err != nil {
// 		return err
// 	}
// 	for _, f := range files {
// 		if err = writeFile(f); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
