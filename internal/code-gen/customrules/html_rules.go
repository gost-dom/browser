package customrules

var htmlRules = SpecRules{
	"Location": {Operations: OperationRules{
		"assign":  {HasError: true},
		"replace": {HasError: true},
		"reload":  {HasError: true},
	}},
	"History": {
		Operations: OperationRules{
			"go":           {HasError: true},
			"back":         {HasError: true},
			"forward":      {HasError: true},
			"pushState":    {HasError: true},
			"replaceState": {HasError: true},
		},
		Attributes: AttributeRules{
			"scrollRestoration": {NotImplemented: true},
		}},
	"HTMLFormElement": {Operations: OperationRules{
		"submit":        {HasError: true},
		"requestSubmit": {HasError: true},
	}},
	"HTMLOrSVGElement": {Operations: OperationRules{
		"focus": {Arguments: ArgumentRules{"options": {Ignore: true}}},
	}},
}
