package configuration

import "github.com/gost-dom/code-gen/packagenames"

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
		"responseXML",
	)
	xhr.Method("open").SetCustomImplementation()
	xhr.Method("upload").SetCustomImplementation()
	xhr.Method("onreadystatechange").Ignore()

	xhrModule.Type("FormData")
}
