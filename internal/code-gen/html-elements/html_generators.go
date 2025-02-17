package htmlelements

var HTMLAnchorElementSpecs = HTMLGeneratorReq{
	InterfaceName:     "HTMLAnchorElement",
	SpecName:          "html",
	GenerateInterface: true,
	// GenerateAttributes: true,
}

var HTMLPackageConfig = map[string]HTMLGeneratorReq{
	"location": {
		InterfaceName:     "Location",
		SpecName:          "html",
		GenerateInterface: true,
	},
	"dom_string_list": {
		InterfaceName:     "DOMStringList",
		SpecName:          "html",
		GenerateInterface: true,
	},
	"html_anchor_element": {
		InterfaceName:     "HTMLAnchorElement",
		SpecName:          "html",
		GenerateInterface: true,
	},
	"html_hyper_link_element_utils": {
		InterfaceName:     "HTMLHyperlinkElementUtils",
		SpecName:          "html",
		GenerateInterface: true,
	},
}

func CreateHTMLGenerators() ([]FileGeneratorSpec, error) {
	return createGenerators(
		HTMLPackageConfig,
		"github.com/stroiman/go-dom/browser/html",
	)
}
