package model

import (
	"iter"

	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

type ESConstructorData struct {
	Spec          *configuration.IdlInterfaceConfiguration
	CustomRule    customrules.InterfaceRule
	IdlInterface  idl.Interface
	Operations    []ESOperation
	Attributes    []ESAttribute
	Constructor   *ESOperation
	RunCustomCode bool
}

// Return the idl mixin interfaces included in this interface AND that has been
// included in the configuration
func (d ESConstructorData) Includes() []idl.Interface {
	var result []idl.Interface
	for _, i := range d.IdlInterface.Includes {
		if _, configured := d.Spec.DomSpec.Interfaces[i.Name]; configured {
			result = append(result, i)
		}
	}
	return result
}

func (d ESConstructorData) GetInternalPackage() string {
	switch d.Name() {
	case "Event":
		return packagenames.Events
	case "MutationObserver", "MutationRecord":
		return packagenames.DomInterfaces
	case "URLSearchParams":
		return packagenames.URLInterfaces
	default:
		return packagenames.PackageName(d.Spec.DomSpec.Name)
	}
}

func (d ESConstructorData) WrapperFunctionsToInstall() iter.Seq[ESOperation] {
	return func(yield func(ESOperation) bool) {
		for _, op := range d.Operations {
			if !op.MethodCustomization.Ignored && !yield(op) {
				return
			}
		}
	}
}

func (d ESConstructorData) AttributesToInstall() iter.Seq[ESAttribute] {
	return func(yield func(ESAttribute) bool) {
		for _, a := range d.Attributes {
			if !yield(a) {
				return
			}
		}
	}
}

func (d ESConstructorData) OperationCallbackInfos() iter.Seq[ESOperation] {
	return func(yield func(ESOperation) bool) {
		for op := range d.WrapperFunctionsToInstall() {
			if !op.MethodCustomization.CustomImplementation && !yield(op) {
				return
			}
		}
	}
}

func (d ESConstructorData) Name() string { return d.Spec.TypeName }

func (d ESConstructorData) WrappedType() g.Generator {
	idlInterfaceName := d.Name()
	return g.NewTypePackage(idlInterfaceName, d.GetInternalPackage())
}
