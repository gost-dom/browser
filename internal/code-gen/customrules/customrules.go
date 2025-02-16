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

import "github.com/gost-dom/webref/idl"

// CustomRules define the rules pr. IDL Spec. The key identifies the name of the
// spec, which corresponds to the file name it was loaded from. E.g., "dom"
// represent the types described in "dom.idl" in the webref specifications
type CustomRules map[string]SpecRules

// SpecRules define the rules for all the types in a single IDL specification
// file. The key is the name of the interface the rule applies to
type SpecRules map[string]InterfaceRule

// InterfaceRule specifies the rules for a specific interface or interface
// mixin.
type InterfaceRule struct {
	Operations OperationRules
}

type OperationRules map[string]OperationRule

type OperationRule struct {
	// By default, an operation is assumed to not generate an error. Override
	// the behaviour by setting this to true.
	HasError   bool
	Attributes AttributeRules
}

type AttributeRules map[string]AttributeRule

type AttributeRule struct {
	Type idl.Type
}

type AttributeTypeRule struct {
	Name    string
	Package string
}

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
			"append": {HasError: true, Attributes: AttributeRules{
				"nodes": {Type: idl.Type{Name: "Node"}},
			}},
			"prepend": {HasError: true, Attributes: AttributeRules{
				"nodes": {Type: idl.Type{Name: "Node"}},
			}},
			"replaceChildren": {HasError: true, Attributes: AttributeRules{
				"nodes": {Type: idl.Type{Name: "Node"}},
			}},
		}},
	},
	"html": {
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
