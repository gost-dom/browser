package configuration

import "github.com/gost-dom/code-gen/packagenames"

func configureURLSpecs(urlSpecs *WebAPIConfig) {
	urlSearchParams := urlSpecs.Type("URLSearchParams")
	urlSearchParams.SkipConstructor = true

	url := urlSpecs.Type("URL")
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
	)
}
