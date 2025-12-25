package model

import (
	"iter"

	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/idltransform"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/configuration"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

type ESConstructorData struct {
	Spec          *configuration.WebIDLConfig
	CustomRule    customrules.InterfaceRule
	IdlInterface  idl.Interface
	Operations    []Callback
	Attributes    []ESAttribute
	Constructor   *Callback
	RunCustomCode bool
}

func (d ESConstructorData) WriteConstructor() bool {
	return d.InstallConstructor() && !d.Spec.SkipConstructor
}

func (d ESConstructorData) InstallConstructor() bool {
	if d.IdlInterface.Mixin || d.IdlInterface.Partial {
		return false
	}
	return true
}

func (d ESConstructorData) AllowConstructor() bool {
	// You _can_ create a Document, but not HTMLDocument, nor other nodes.
	if d.IdlInterface.Name == "Document" {
		return true
	}
	if IsNodeType(d.IdlInterface.Name) {
		return false
	}
	if d.Constructor == nil {
		return false
	}
	return true
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
	if pkg := idltransform.InternalPackage(d.Name()); pkg != "" {
		return pkg
	}
	switch d.Name() {
	case "Event", "EventTarget":
		return packagenames.Events
	case "MutationObserver", "MutationRecord":
		return packagenames.DomInterfaces
	case "URLSearchParams", "URL":
		return packagenames.URLInterfaces
	case "XMLHttpRequest":
		return packagenames.HTMLInternal
	case "FormData":
		return packagenames.Html
	default:
		return packagenames.PackageName(d.Spec.DomSpec.Name)
	}
}

func (d ESConstructorData) WrapperFunctionsToInstall() iter.Seq[Callback] {
	return func(yield func(Callback) bool) {
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

func (d ESConstructorData) OperationCallbackInfos() iter.Seq[Callback] {
	return func(yield func(Callback) bool) {
		for op := range d.WrapperFunctionsToInstall() {
			if !op.MethodCustomization.CustomImplementation && !yield(op) {
				return
			}
		}
	}
}

func (d ESConstructorData) Name() string { return d.Spec.TypeName }
func (d ESConstructorData) Extends() string {
	return d.IdlInterface.Inheritance
}

func (d ESConstructorData) WrappedType() g.Generator {
	if override := d.Spec.OverrideWrappedType; override != nil {
		res := g.NewTypePackage(override.Name, override.Package)
		if override.Pointer {
			res = g.Type{Generator: res.Pointer()}
		}
		return res
	}
	var typeName string
	var interfacePackage string
	if name := d.CustomRule.OverrideTypeName; name != "" {
		typeName = name
	} else {
		typeName = d.Name()
	}
	if name := d.CustomRule.InterfacePackage; name != "" {
		interfacePackage = string(name)
	} else {
		interfacePackage = d.GetInternalPackage()
	}
	res := g.NewTypePackage(typeName, interfacePackage)
	if d.CustomRule.OutputType == customrules.OutputTypeStruct {
		return res.Pointer()
	}
	return res
}
