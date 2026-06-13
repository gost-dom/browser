package configuration

import "github.com/gost-dom/browser/internal/code-gen/packagenames"

// configureFetchSpecs registers the code-generation customizations for the Fetch
// API types (Request, Response, Body, Headers, and the global fetch operation),
// marking unsupported members as not-implemented and flagging the members that
// have hand-written custom implementations.
func configureFetchSpecs(specs *WebAPIConfig) {
	req := specs.Type("Request")
	// req.OverrideWrappedType = &GoType{Package: packagenames.Fetch, Name: "Request", Pointer: true}
	req.MarkMembersAsNotImplemented(
		"clone", "method", "destination", "referrer", "referrerPolicy",
		"mode", "credentials", "cache", "redirect", "integrity", "keepalive",
		"isReloadNavigation", "isHistoryNavigation", "signal",
		"duplex",
	)

	res := specs.Type("Response")
	res.OverrideWrappedType = &GoType{Package: packagenames.Fetch, Name: "Response", Pointer: true}
	res.SkipConstructor = true
	res.MarkMembersAsNotImplemented(
		"type", "clone", "url", "redirected",
	)
	res.Method("ok").SetCustomImplementation()
	res.Method("statusText").SetCustomImplementation()

	body := specs.Type("Body")
	body.MarkMembersAsNotImplemented(
		"blob", "formData", "bodyUsed",
	)
	body.Method("json").SetCustomImplementation()
	body.Method("text").SetCustomImplementation()
	body.Method("arrayBuffer").SetCustomImplementation()
	body.Method("bytes").SetCustomImplementation()

	headers := specs.Type("Headers")
	headers.MarkMembersAsNotImplemented("getSetCookie")

	scope := specs.Type("WindowOrWorkerGlobalScope")
	scope.Partial = true
	scope.Method("fetch").SetCustomImplementation()
}
