package customrules

import (
	"github.com/gost-dom/browser/internal/code-gen/gotypes"
	"github.com/gost-dom/browser/internal/code-gen/packagenames"
	"github.com/gost-dom/webref/idl"
)

var timeDuration = gotypes.TimeDuration
var taskHandle = gotypes.TaskHandle

var htmlRules = SpecRules{
	"DOMStringMap": {OutputType: OutputTypeStruct},
	"Location": {
		IsEntity: true,
		Operations: OperationRules{
			"assign":  {HasError: true},
			"replace": {HasError: true},
			"reload":  {HasError: true},
		}},
	"Document": {
		Attributes: AttributeRules{
			"body": {SetterHasError: true},
		},
	},
	"Element": {
		Operations: OperationRules{
			"insertAdjacentHTML": {HasError: true},
		},
		Attributes: AttributeRules{
			"outerHTML": {SetterHasError: true},
			"innerHTML": {SetterHasError: true},
		},
	},
	"History": {
		Operations: OperationRules{
			"go":      {HasError: true},
			"back":    {HasError: true},
			"forward": {HasError: true},
			"pushState": {HasError: true, Arguments: ArgumentRules{
				"data":   {Type: idl.Type{Name: "HistoryState"}},
				"unused": {Ignore: true},
				"url":    {ZeroAsDefault: true},
			}},
			"replaceState": {HasError: true, Arguments: ArgumentRules{
				"data":   {Type: idl.Type{Name: "HistoryState"}},
				"unused": {Ignore: true},
				"url":    {ZeroAsDefault: true},
			}},
		},
		Attributes: AttributeRules{
			"scrollRestoration": {NotImplemented: true},
			"state":             {OverrideType: SamePackageType{"HistoryState"}},
		}},
	"HTMLFormElement": {Operations: OperationRules{
		"submit":        {HasError: true},
		"requestSubmit": {HasError: true},
	}},
	"HTMLOrSVGElement": {Operations: OperationRules{
		"focus": {Arguments: ArgumentRules{"options": {Ignore: true}}},
	}},
	"WindowOrWorkerGlobalScope": {Operations: OperationRules{
		"setTimeout": {
			ReturnType: taskHandle,
			Arguments: ArgumentRules{
				"timeout":   {GoType: timeDuration},
				"arguments": {Ignore: true},
			}},
		"setInterval": {
			ReturnType: taskHandle,
			Arguments: ArgumentRules{
				"timeout":   {GoType: timeDuration},
				"arguments": {Ignore: true},
			}},
		"clearTimeout":  {Arguments: ArgumentRules{"id": {GoType: taskHandle}}},
		"clearInterval": {Arguments: ArgumentRules{"id": {GoType: taskHandle}}},
	}},
	"MessageChannel": {InterfacePackage: packagenames.HTMLInterfaces},
	"MessagePort": {
		InterfacePackage: packagenames.HTMLInterfaces,
		IsEntity:         true,
		Operations: OperationRules{
			"postMessage": {Ignore: true},
		}},
}

func init() {
	IgnoreMembers(htmlRules,
		Overrides{"HTMLTextAreaElement": {
			Operations: []string{
				"checkValidity",
				"setCustomValidity",
				"reportValidity",
				"select",
				"setRangeText",
				"setSelectionRange",
			},
			Attributes: []string{
				"autocomplete",
				"cols", "rows",
				"dirName",
				"disabled",
				"minLength", "maxLength",
				"name",
				"placeholder",
				"readOnly",
				"required",
				"wrap",
				"type",
				"defaultValue",
				"textLength",
				"willValidate",
				"validity",
				"form",
				"labels",
				"validationMessage",
				"selectionStart", "selectionEnd", "selectionDirection",
			},
		}})
	IgnoreMembers(htmlRules,
		Overrides{"WindowOrWorkerGlobalScope": {
			Operations: []string{
				"atob",
				"btoa",
				"createImageBitmap",
				"structuredClone",
				"reportError",
			},
			Attributes: []string{
				"origin", "isSecureContext", "crossOriginIsolated",
			},
		}})
}
