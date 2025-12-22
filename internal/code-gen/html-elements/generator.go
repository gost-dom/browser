package htmlelements

import (
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/webref/idl"
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

func packageName(name string) string {
	switch name {
	case "dominterfaces":
		return packagenames.DomInterfaces
	}
	return fmt.Sprintf("%s/%s", packagenames.BASE_PKG, name)

}

func CreatePackageGenerators(name string) (res []FileGeneratorSpec, err error) {
	if config, err := GetPackageGeneratorSpecs(name); err == nil {
		return createGenerators(config, packageName(name))
	}
	return
}

func CreateInterfaceFileGenerators(destPackage string) ([]FileGeneratorSpec, error) {
	config, ok := PackageInterfacesConfiguration[destPackage]
	if !ok {
		return nil, nil
	}
	spec, err := idl.Load(config.webApi)
	if err != nil {
		return nil, err
	}
	res := make([]FileGeneratorSpec, len(config.interfaces))
	for i, intf := range config.interfaces {
		idlIntf := spec.Interfaces[intf]
		res[i] = FileGeneratorSpec{
			OutputFile: internal.TypeNameToFileName(intf),
			Package:    destPackage,
			Generator:  generateInterface(config.webApi, destPackage, idlIntf),
		}
	}
	return res, nil
}

func CreateImplementationPackage(name string) error {
	files, err := CreatePackageGenerators(name)
	if err != nil {
		return err
	}
	interfaceFiles, err := CreateInterfaceFileGenerators(name)
	if err != nil {
		return err
	}
	files = append(files, interfaceFiles...)
	for _, f := range files {
		if err = writeFile(f); err != nil {
			return err
		}
	}
	return nil
}
