package configuration

import (
	"fmt"
)

func CreateV8SpecsForSpec(spec string) WebIdlConfigurations {
	specs := NewWrapperGeneratorsSpec()

	switch spec {
	case "dom":
		ConfigureDOMSpecs(&specs)
	case "uievents":
		ConfigureEventSpecs(&specs)
	case "html":
		ConfigureHTMLSpecs(&specs)
	case "xhr":
		configureXHRSpecs(specs.Module("xhr"))
	case "url":
		configureURLSpecs(specs.Module("url"))
	case "fetch":
		configureFetchSpecs(specs.Module("fetch"))
	case "streams":
		configureStreamsSpecs(specs.Module("streams"))
	default:
		panic(fmt.Sprintf("bad spec: %s", spec))
	}

	return specs
}
