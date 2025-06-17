package configuration

import "github.com/gost-dom/code-gen/packagenames"

func configureFetchSpecs(specs *WebAPIConfig) {
	req := specs.Type("Request")
	req.OverrideWrappedType = &GoType{Package: packagenames.Fetch, Name: "Request", Pointer: true}
	req.MarkMembersAsNotImplemented(
		"clone", "method", "headers", "destination", "referrer", "referrerPolicy",
		"mode", "credentials", "cache", "redirect", "integrity", "keepalive",
		"isReloadNavigation", "isHistoryNavigation", "signal",
		"duplex",
	)
}
