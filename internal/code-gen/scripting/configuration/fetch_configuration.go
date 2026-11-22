package configuration

import "github.com/gost-dom/browser/internal/code-gen/packagenames"

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
	// ok and statusText have hand-written implementations in response.go.
	res.Method("ok").SetCustomImplementation()
	res.Method("statusText").SetCustomImplementation()

	body := specs.Type("Body")
	body.MarkMembersAsNotImplemented(
		"arrayBuffer", "blob", "bytes", "formData", "bodyUsed",
	)
	body.Method("json").SetCustomImplementation()
	// text has a hand-written implementation in body.go.
	body.Method("text").SetCustomImplementation()

	headers := specs.Type("Headers")
	headers.MarkMembersAsNotImplemented("getSetCookie")

	scope := specs.Type("WindowOrWorkerGlobalScope")
	scope.Partial = true
	scope.Method("fetch").SetCustomImplementation()
}
