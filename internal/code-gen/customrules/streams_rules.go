package customrules

var streamsRules = SpecRules{
	"ReadableStream": {Operations: OperationRules{
		"constructor": {
			Arguments: ArgumentRules{
				"underlyingSource": {ZeroAsDefault: true},
				"strategy":         {Variadic: true},
			},
		},
		"getReader": {
			Arguments: ArgumentRules{"options": {Variadic: true}},
		},
	}},
}
