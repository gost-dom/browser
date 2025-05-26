package customrules

import "github.com/gost-dom/webref/idl"

var xhrRules = SpecRules{
	"XMLHttpRequest": {Operations: OperationRules{
		"getAllResponseHeaders": {HasError: true},
		"send":                  {HasError: true},
		"abort":                 {HasError: true},
		"overrideMimeType":      {HasError: true},
	}},
	"FormData": {
		OutputType: OutputTypeStruct,
		Operations: OperationRules{
			"append": {Arguments: ArgumentRules{
				"value": {Type: idl.Type{Name: "FormDataValue"}},
			}},
			"set": {Arguments: ArgumentRules{
				"value": {Type: idl.Type{Name: "FormDataValue"}},
			}},
		}},
}
