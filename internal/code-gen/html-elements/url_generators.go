package htmlelements

var URLPackageConfig = setSpecName("url", GeneratorConfig{
	"url": {
		InterfaceName:     "URL",
		GenerateInterface: true,
	},
	"url_search_params": {
		InterfaceName:     "URLSearchParams",
		GenerateInterface: true,
	},
})
