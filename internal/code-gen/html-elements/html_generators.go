package htmlelements

var HTMLAnchorElementSpecs = HTMLGeneratorReq{
	InterfaceName:     "HTMLAnchorElement",
	SpecName:          "html",
	GenerateInterface: true,
	// GenerateAttributes: true,
}

var HTMLPackageConfig = setSpecName("html", GeneratorConfig{
	"location": {
		InterfaceName:     "Location",
		GenerateInterface: true,
	},
	"dom_string_list": {
		InterfaceName:     "DOMStringList",
		GenerateInterface: true,
	},
	"html_anchor_element": {
		InterfaceName:     "HTMLAnchorElement",
		GenerateInterface: true,
	},
	"html_hyper_link_element_utils": {
		InterfaceName:     "HTMLHyperlinkElementUtils",
		GenerateInterface: true,
	},
	"html_or_svg_element": {
		InterfaceName:     "HTMLOrSVGElement",
		GenerateInterface: true,
	},
})

func CreateHTMLGenerators() ([]FileGeneratorSpec, error) {
	return createGenerators(
		HTMLPackageConfig,
		"github.com/stroiman/go-dom/browser/html",
	)
}
