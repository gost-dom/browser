package configuration

import "github.com/gost-dom/code-gen/packagenames"

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
