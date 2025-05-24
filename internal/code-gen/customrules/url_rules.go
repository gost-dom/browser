package customrules

var urlRules = SpecRules{
	"URL": {Operations: OperationRules{
		"toJSON": {HasError: true},
	}},
}
