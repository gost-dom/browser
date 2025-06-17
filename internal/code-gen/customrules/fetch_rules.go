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
}
