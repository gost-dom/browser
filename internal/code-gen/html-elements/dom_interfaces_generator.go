package htmlelements

var DOMInterfacesPackageConfig = GeneratorConfig{
	"abort_signal": {
		InterfaceName:     "AbortSignal",
		SpecName:          "dom",
		GenerateInterface: true,
	},
	"abort_controller": {
		InterfaceName:     "AbortController",
		SpecName:          "dom",
		GenerateInterface: true,
	},
	"mutation_observer": {
		InterfaceName:     "MutationObserver",
		SpecName:          "dom",
		GenerateInterface: true,
	},
	"mutation_record": {
		InterfaceName: "MutationRecord",
		SpecName:      "dom",
	},
}
