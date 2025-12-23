package interfaces

import (
	"fmt"

	"github.com/gost-dom/code-gen/customrules"
	htmlelements "github.com/gost-dom/code-gen/html-elements"
	"github.com/gost-dom/code-gen/idltransform"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/packagenames"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

func CreateInterfaceFileGenerators(destPackage string) ([]htmlelements.FileGeneratorSpec, error) {
	config, ok := PackageInterfacesConfiguration[destPackage]
	if !ok {
		return nil, nil
	}
	webApi := config.webApi
	spec, err := idl.Load(webApi)
	if err != nil {
		return nil, err
	}
	res := make([]htmlelements.FileGeneratorSpec, len(config.interfaces))
	for i, intf := range config.interfaces {
		idlIntf, ok := spec.Interfaces[intf]
		if !ok {
			return nil, fmt.Errorf("%s (%s): interface not found in idl spec", intf, webApi)
		}
		res[i] = htmlelements.FileGeneratorSpec{
			OutputFile: internal.TypeNameToFileName(intf),
			Package:    packagenames.ExpandPackageName(destPackage),
			Generator:  generateInterface(config.webApi, destPackage, idlIntf),
		}
	}
	return res, nil
}

func GenerateInterface(webApi string, target string, name string) (g.Generator, error) {
	spec, err := idl.Load(webApi)
	if err != nil {
		return nil, err
	}
	idlInterface, ok := spec.Interfaces[name]
	if !ok {
		return nil, fmt.Errorf("GenerateInterface: %s: not found in package: %s", name, webApi)
	}
	return generateInterface(webApi, target, idlInterface), nil
}

func generateInterface(webApi string, target string, idlInterface idl.Interface) g.Generator {
	apiRules := customrules.GetSpecRules(webApi)
	intfRules := apiRules[idlInterface.Name]
	attributes := make([]IdlInterfaceAttribute, 0)
	operations := make([]IdlInterfaceOperation, 0)
	includes := make([]IdlInterfaceInclude, len(idlInterface.Includes))
	iterableTypes := make([]idltransform.IdlType, len(idlInterface.IterableTypes))

	interfaces := make([]idl.Interface, 1+len(idlInterface.Includes))
	interfaces[0] = idlInterface
	copy(interfaces[1:], idlInterface.Includes)
	result := IdlInterface{
		SpecName:  webApi,
		Name:      idlInterface.Name,
		Inherits:  idlInterface.InternalSpec.Inheritance,
		Includes:  includes,
		Rules:     intfRules,
		TargetPkg: target,
	}

	for idx, i := range idlInterface.Includes {
		includes[idx] = IdlInterfaceInclude{i}
	}

	for _, a := range idlInterface.Attributes {
		attributeRule := intfRules.Attributes[a.Name]
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
			Type:     idltransform.IdlType{Type: attrType, TargetPackage: target},
			ReadOnly: a.Readonly,
		})
	}
	for _, o := range idlInterface.Operations {
		if o.Stringifier {
			result.HasStringifier = true
			if o.Name == "" {
				continue
			}
		}
		operationRule := intfRules.Operations[o.Name]
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
			IdlInterfaceOperation{
				o,
				arguments,
				idltransform.IdlType{Type: o.ReturnType, TargetPackage: target},
				operationRule,
				target,
			},
		)
	}
	for i, t := range idlInterface.IterableTypes {
		iterableTypes[i] = idltransform.IdlType{Type: t, TargetPackage: target}
	}
	result.Attributes = attributes
	result.Operations = operations
	result.IterableTypes = iterableTypes
	return result
}

func CreateInterfaces(name string) error {
	files, err := CreateInterfaceFileGenerators(name)
	if err != nil {
		return err
	}
	for _, f := range files {
		if err = htmlelements.WriteFile(f); err != nil {
			return err
		}
	}
	return nil
}
