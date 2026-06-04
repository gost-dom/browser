package customrules

var cssomViewRules = SpecRules{
	"Range": {
		Operations: OperationRules{
			"getClientRects": {Ignore: true},
		},
	},
	"Element": {
		Operations: OperationRules{
			"getClientRects":  {Ignore: true},
			"checkVisibility": {Ignore: true},
			"scrollIntoView":  {Ignore: true},
			"scroll":          {Ignore: true},
			"scrollTo":        {Ignore: true},
			"scrollBy":        {Ignore: true},
		},
		Attributes: AttributeRules{
			"scrollTop":      {Ignore: true},
			"scrollLeft":     {Ignore: true},
			"scrollWidth":    {Ignore: true},
			"scrollHeight":   {Ignore: true},
			"clientTop":      {Ignore: true},
			"clientLeft":     {Ignore: true},
			"clientWidth":    {Ignore: true},
			"clientHeight":   {Ignore: true},
			"currentCSSZoom": {Ignore: true},
		},
	},
}
