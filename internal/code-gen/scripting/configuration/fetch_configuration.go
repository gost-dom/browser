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

	res := specs.Type("Response")
	res.OverrideWrappedType = &GoType{Package: packagenames.Fetch, Name: "Response", Pointer: true}
	res.SkipConstructor = true
	res.MarkMembersAsNotImplemented(
		"body", "type", "clone", "url", "redirected", "ok", "statusText", "headers",
	)

	body := specs.Type("Body")
	body.MarkMembersAsNotImplemented(
		"arrayBuffer", "blob", "bytes", "formData", "text", "body", "bodyUsed",
	)
	body.Method("json").SetCustomImplementation()
}
