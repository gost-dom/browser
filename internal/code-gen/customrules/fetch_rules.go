package customrules

var fetchRules = SpecRules{
	"Request": {
		OutputType: OutputTypeStruct,
		Operations: OperationRules{
			"constructor": {
				Arguments: ArgumentRules{
					"init": {Variadic: true},
				},
			},
		},
		Attributes: AttributeRules{
			"url": {Callable: true},
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
