package customrules

var xhrRules = SpecRules{
	"XMLHttpRequest": {Operations: OperationRules{
		"getAllResponseHeaders": {HasError: true},
		"send":                  {HasError: true},
		"abort":                 {HasError: true},
		"overrideMimeType":      {HasError: true},
	}},
}
