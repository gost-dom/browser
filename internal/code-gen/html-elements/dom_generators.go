package htmlelements

var DOMPackageConfig = GeneratorConfig{
	"parent_node": {
		InterfaceName:     "ParentNode",
		SpecName:          "dom",
		GenerateInterface: true,
	},
	"html_collection": {
		InterfaceName:     "HTMLCollection",
		SpecName:          "dom",
		GenerateInterface: true,
	},
}

func CreateDOMGenerators() ([]FileGeneratorSpec, error) {
	return createGenerators(
		DOMPackageConfig,
		"github.com/stroiman/go-dom/browser/dom",
	)
}
