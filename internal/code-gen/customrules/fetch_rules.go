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
		Operations: OperationRules{
			"constructor": {
				Arguments: ArgumentRules{
					"init": {Variadic: true},
				},
			},
			"append": {HasError: true},
			"set":    {HasError: true},
		},
	},
	"Response": {OutputType: OutputTypeStruct},
}
