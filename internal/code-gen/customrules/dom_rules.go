package customrules

import "github.com/gost-dom/webref/idl"

var domRules = SpecRules{
	"Event": {OutputType: OutputTypeStruct},
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
						Type:     idl.Type{Name: "ObserveOption"},
						Variadic: true,
					},
				}},
		}},
	"MutationRecord": {
		InterfacePackage: DomInterfaces,
		OutputType:       OutputTypeStruct,
	},
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
