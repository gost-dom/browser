// Package customrules specify how IDL interfaces are implemented by Go types.
//
// The same IDL interfaces are used to generate different types of code, e.g.
//   - Go interfaces in the DOM implementation
//   - Possibly Go types implementing those interfaces
//   - JavaScript wrappers on top of those interfaces (for both V8 and Goja)
//
// Not all relevant properties exist in the IDL specification, e.g., if an
// operation can result in an error or not. This package autments the IDL specs
// with this missing information, e.g., so generated JavaScript wrappers only
// add error handling to the operations that actually produce errors.
package customrules

import (
	"reflect"

	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/webref/idl"
)

const (
	DomInterfaces = Package(packagenames.DomInterfaces)
)

// CustomRules define the rules pr. IDL Spec. The key identifies the name of the
// spec, which corresponds to the file name it was loaded from. E.g., "dom"
// represent the types described in "dom.idl" in the webref specifications
type CustomRules map[string]SpecRules

type Package string

// SpecRules define the rules for all the types in a single IDL specification
// file. The key is the name of the interface the rule applies to
type SpecRules map[string]InterfaceRule

// Specifies how the type is represented in the internal Go model.
type OutputType string

const (
	OutputTypeInterface OutputType = "interface"
	OutputTypeStruct    OutputType = "struct"
)

// InterfaceRule specifies the rules for a specific interface or interface
// mixin.
type InterfaceRule struct {
	InterfacePackage Package
	OutputType       OutputType
	Operations       OperationRules
	Attributes       AttributeRules
}

type OperationRules map[string]OperationRule

type OperationRule struct {
	// By default, an operation is assumed to not generate an error. Override
	// the behaviour by setting this to true.
	HasError    bool
	Arguments   ArgumentRules
	DocComments string
}

func (r OperationRule) Argument(name string) ArgumentRule {
	if r.Arguments == nil {
		return ArgumentRule{}
	}
	return r.Arguments[name]
}

type ArgumentRules map[string]ArgumentRule

type ArgumentRule struct {
	Type          idl.Type
	Variadic      bool
	Ignore        bool
	ZeroAsDefault bool // Whether the "zero" value, e.g. empty string, is default
}

func (r ArgumentRule) OverridesType() bool {
	return !reflect.ValueOf(r.Type).IsZero()
}

var rules = CustomRules{
	"dom":     domRules,
	"fetch":   fetchRules,
	"html":    htmlRules,
	"url":     urlRules,
	"xhr":     xhrRules,
	"streams": streamsRules,
}

func GetSpecRules(specName string) SpecRules {
	if res, ok := rules[specName]; ok {
		return res
	}
	return make(SpecRules)
}

type AttributeRules map[string]AttributeRule

type AttributeRule struct {
	// NotImplemented indicates that the IDL does not have an implementation in
	// Go.
	NotImplemented bool
	OverrideType   IdlTyper
}

// IdlTyper is the interface with the IdlType() method that can generate an
// idl.Type value.
//
// The intended use case is when the type is somewhat vague, and a more specific
// type is helpful. The motivating example is History.state, that has the type
// "any" in the IDL specification, but valid values are JSON serializable data,
// so this is internally stored as a string.
type IdlTyper interface{ IdlType() idl.Type }

type SamePackageType struct{ Name string }

func (t SamePackageType) IdlType() idl.Type {
	return idl.Type{Kind: idl.KindSimple, Name: t.Name}
}
