package configuration

import (
	"fmt"

	"github.com/gost-dom/code-gen/packagenames"
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
	default:
		panic(fmt.Sprintf("bad spec: %s", spec))
	}

	return specs
}

func configureXHRSpecs(xhrModule *WebAPIConfig) {
	xhrEventTarget := xhrModule.Type("XMLHttpRequestEventTarget")
	xhrEventTarget.OverrideWrappedType = &GoType{
		Package: packagenames.Events,
		Name:    "EventTarget",
	}

	xhr := xhrModule.Type("XMLHttpRequest")

	// TODO: Just need to support non-node objects
	// xhr.SkipWrapper = true

	xhr.MarkMembersAsNotImplemented(
		"readyState",
		"responseType",
		"responseXML",
	)
	xhr.Method("open").SetCustomImplementation()
	xhr.Method("upload").SetCustomImplementation()
	xhr.Method("onreadystatechange").Ignore()

	formData := xhrModule.Type("FormData")
	formData.RunCustomCode = true
}

func configureURLSpecs(urlSpecs *WebAPIConfig) {
	urlSearchParams := urlSpecs.Type("URLSearchParams")
	urlSearchParams.SkipConstructor = true
	urlSearchParams.RunCustomCode = true
	// urlSearchParams.Method("get").SetCustomImplementation()

	url := urlSpecs.Type("URL")
	// TODO: Just need to use a different base class for non-nodes
	url.SkipWrapper = true
	url.OverrideWrappedType = &GoType{Package: packagenames.URL, Name: "URL", Pointer: true}
	url.MarkMembersAsNotImplemented(
		"setHref",
		"setProtocol",
		"username",
		"password",
		"setHost",
		"setPort",
		"setHostname",
		"setPathname",
		"searchParams",
		"setHash",
		"setSearch",
	)

}
