package htmlelements

var DOMPackageConfig = GeneratorConfig{
	"non_document_type_child_node": {
		InterfaceName:     "NonDocumentTypeChildNode",
		SpecName:          "dom",
		GenerateInterface: true,
	},
	"parent_node": {
		InterfaceName:     "ParentNode",
		SpecName:          "dom",
		GenerateInterface: true,
	},
	/* Doesn't generate an All() function
	"html_collection": {
		InterfaceName:     "HTMLCollection",
		SpecName:          "dom",
		GenerateInterface: true,
	},
	*/
}

func CreateDOMGenerators() ([]FileGeneratorSpec, error) {
	return createGenerators(
		DOMPackageConfig,
		"github.com/stroiman/go-dom/browser/dom",
	)
}
