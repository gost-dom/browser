package customrules

import "github.com/gost-dom/webref/idl"

var htmlRules = SpecRules{
	"DOMStringMap": {OutputType: OutputTypeStruct},
	"Location": {
		IsEntity: true,
		Operations: OperationRules{
			"assign":  {HasError: true},
			"replace": {HasError: true},
			"reload":  {HasError: true},
		}},
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
	"WindowOrWorkerGlobalScope": InterfaceRule{}.IgnoreOperations(
		"atob",
		"btoa",
		"createImageBitmap",
		"structuredClone",
		"reportError",
	).IgnoreAttributes("origin", "isSecureContext", "crossOriginIsolated"),
}
