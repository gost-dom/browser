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
	outputFileName := fmt.Sprintf("%s_generated.go", s.Name)
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

func CreatePackageGenerators(name string) ([]FileGeneratorSpec, error) {
	config, found := PackageConfigs[name]
	if !found {
		return nil, fmt.Errorf("CreatePackageGenerators: No configuration for package %s", name)
	}
	packageName := fmt.Sprintf("%s/%s", packagenames.BASE_PKG, name)
	return createGenerators(config, packageName)
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
