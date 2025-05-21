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

	"github.com/gost-dom/code-gen/customrules/typerule"
	. "github.com/gost-dom/code-gen/customrules/typerule"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/webref/idl"
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
}

type OperationRules map[string]OperationRule

const (
	DomInterfaces = Package(packagenames.DomInterfaces)
)

type OperationRule struct {
	// By default, an operation is assumed to not generate an error. Override
	// the behaviour by setting this to true.
	HasError    bool
	Arguments   ArgumentRules
	DocComments string
	ReturnType  *TypeRule
}

func (r OperationRule) Argument(name string) ArgumentRule {
	if r.Arguments == nil {
		return ArgumentRule{}
	}
	return r.Arguments[name]
}

type ArgumentRules map[string]ArgumentRule

type ArgumentRule struct {
	Type     idl.Type
	Variadic bool
	Ignore   bool
}

func (r ArgumentRule) OverridesType() bool { return !reflect.ValueOf(r.Type).IsZero() }

type AttributeTypeRule struct {
	Name    string
	Package string
}

var parentNodeQueryOperation = OperationRule{HasError: true}
var parentNodeOperation = OperationRule{
	HasError: true,
	DocComments: `Note that the IDL operation accepts either string or node values. This interface
requires an explicit a [Node]. Use [Document.CreateText] to convert a string to
a Node.

See also: https://developer.mozilla.org/en-US/docs/Web/API/Element`,
	Arguments: ArgumentRules{
		"nodes": {Type: idl.Type{Name: "Node"}},
	}}

// The Go functions that correspond to a specific event generally return a bool,
// which corresponds to the return value of DispatchEvent, indicating if the
// event was cancelled or not.
var eventOperation = OperationRule{ReturnType: typerule.Bool}

var rules = CustomRules{
	"url": {
		"URL": {Operations: OperationRules{
			"toJSON": {HasError: true},
		}},
	},
	"dom": {
		"DOMTokenList": {Operations: OperationRules{
			"add": {HasError: true},
		}},
		"Node": {Operations: OperationRules{
			"insertBefore": {HasError: true},
			"appendChild":  {HasError: true},
			"removeChild":  {HasError: true},
		}},
		"Element": {Operations: OperationRules{
			"matches": {HasError: true},
		}},
		"ParentNode": {Operations: OperationRules{
			"append":           parentNodeOperation,
			"prepend":          parentNodeOperation,
			"replaceChildren":  parentNodeOperation,
			"querySelector":    parentNodeQueryOperation,
			"querySelectorAll": parentNodeQueryOperation,
		}},
		"MutationObserver": {
			InterfacePackage: DomInterfaces,
			Operations: OperationRules{
				"observe": {
					HasError: true,
					Arguments: ArgumentRules{
						"options": {
							Type:     idl.Type{Name: "func(*MutationObserverInit)"},
							Variadic: true,
						},
					}},
			}},
		"MutationRecord": {
			InterfacePackage: DomInterfaces,
			OutputType:       OutputTypeStruct,
		},
	},
	"html": {
		"Location": {Operations: OperationRules{
			"assign":  {HasError: true},
			"replace": {HasError: true},
			"reload":  {HasError: true},
		}},
		"History": {Operations: OperationRules{
			"go":           {HasError: true},
			"back":         {HasError: true},
			"forward":      {HasError: true},
			"pushState":    {HasError: true},
			"replaceState": {HasError: true},
		}},
		"HTMLFormElement": {Operations: OperationRules{
			"submit":        {HasError: true},
			"requestSubmit": {HasError: true},
		}},
		"HTMLOrSVGElement": {Operations: OperationRules{
			"focus": {Arguments: ArgumentRules{"options": {Ignore: true}}},
		}},
	},
	"xhr": {
		"XMLHttpRequest": {Operations: OperationRules{
			"getAllResponseHeaders": {HasError: true},
			"send":                  {HasError: true},
			"abort":                 {HasError: true},
			"overrideMimeType":      {HasError: true},
		}},
	},
}

func GetSpecRules(specName string) SpecRules {
	if res, ok := rules[specName]; ok {
		return res
	}
	return make(SpecRules)
}
