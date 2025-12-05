package customrules

var fetchRules = SpecRules{
	"Request": {
		Operations: OperationRules{
			"constructor": {
				Arguments: ArgumentRules{
					"init": {Variadic: true},
				},
			},
		},
	},
	"Headers": {
		OutputType: OutputTypeStruct,
		Operations: OperationRules{
			"constructor": {
				Arguments: ArgumentRules{
					"init": {Variadic: true},
				},
			},
		},
	},
	"Response": {OutputType: OutputTypeStruct},
}
