package configuration

import (
	"fmt"
)

func CreateV8SpecsForSpec(spec string) WebIdlConfigurations {
	specs := NewWrapperGeneratorsSpec()

	switch spec {
	case "dom":
		ConfigureDOMSpecs(specs.Module("dom"))
	case "uievents":
		ConfigureEventSpecs(specs.Module("uievents"))
		ConfigurePointerEventSpecs(specs.Module("pointerevents4"))
	case "html":
		ConfigureHTMLSpecs(specs.Module("html"))
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
