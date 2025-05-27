package configuration

import "github.com/gost-dom/code-gen/packagenames"

func CreateV8Specs() WebIdlConfigurations {
	specs := CreateSpecs()
	xhrModule := specs.Module("xhr")
	xhr := xhrModule.Type("XMLHttpRequest")
	// TODO: Just need to support non-node objects
	xhr.SkipWrapper = true

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

	urlSpecs := specs.Module("url")
	urlSearchParams := urlSpecs.Type("URLSearchParams")
	urlSearchParams.Strategy = StrategyAbstraction
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

	domSpecs := specs.Module("dom")
	configureMutationObserver(domSpecs)

	domSpecs.Type("NonDocumentTypeChildNode")
	document := domSpecs.Type("Document")
	document.RunCustomCode = true // Set instance properties
	document.MarkMembersAsNotImplemented(
		"createNodeIterator",
		"createTreeWalker",
		"getElementsByTagName",
		"getElementsByTagNameNS",
		"getElementsByClassName",
		"createProcessingInstruction",
		"importNode",
		"adoptNode",
		"createRange",
		"createEvent",
		"implementation",
		"documentURI",
		"doctype",
		"contentType",
		"inputEncoding",
		"charset", "characterSet",
		"compatMode", "URL",
		"createAttributeNS",
		"createElementNS",

		// Custom Implementation
		"createCDATASection",
	)
	document.Method("createElement").SetCustomImplementation()
	document.Method("createTextNode").SetCustomImplementation()

	nodeList := domSpecs.Type("NodeList")
	nodeList.RunCustomCode = true

	parentNode := domSpecs.Type("ParentNode")
	parentNode.Method("children").Ignore()
	parentNode.Method("append").SetCustomImplementation()
	parentNode.Method("prepend").SetCustomImplementation()
	parentNode.Method("replaceChildren").SetCustomImplementation()

	domElement := domSpecs.Type("Element")
	domElement.SkipWrapper = true
	domElement.RunCustomCode = true
	domElement.Method("classList").SetCustomImplementation()

	domElement.MarkMembersAsNotImplemented(
		"hasAttributes",
		"hasAttributeNS",
		"getAttributeNames",
		"getAttributeNS",
		"setAttributeNS",
		"removeAttributeNode",
		"removeAttributeNS",
		"toggleAttribute",
		"toggleAttributeForce",
		"setAttributeNode",
		"setAttributeNodeNS",
		"getAttributeNode",
		"getAttributeNodeNS",
		"getElementsByTagName",
		"getElementsByTagNameNS",
		"getElementsByClassName",
		"insertAdjacentElement",
		"insertAdjacentText",
		"namespaceURI",
		"prefix",
		"localName",
		"shadowRoot",
		"slot",
		"className",
		"decodeShadowRootInit",
		"attachShadow",
	)

	domElement.MarkMembersAsIgnored(
		// HTMX fails if these exist but throw
		"webkitMatchesSelector",
		"closest",
	)

	domTokenList := domSpecs.Type("DOMTokenList")
	domTokenList.RunCustomCode = true
	domTokenList.Method("toggle").SetCustomImplementation()
	domTokenList.Method("supports").SetNotImplemented()

	htmlSpecs := specs.Module("html")
	htmlSpecs.SetMultipleFiles(true)

	htmlTemplateElement := htmlSpecs.Type("HTMLTemplateElement")
	htmlTemplateElement.Method("shadowRootMode").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootDelegatesFocus").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootClonable").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootSerializable").SetNotImplemented()

	htmlElement := htmlSpecs.Type("HTMLElement")
	htmlElement.Method("hidden").Ignore()
	htmlElement.Method("attachInternals").Ignore()
	htmlElement.Method("showPopover").Ignore()
	htmlElement.Method("hidePopover").Ignore()
	htmlElement.Method("togglePopover").Ignore()
	htmlElement.Method("title").Ignore()
	htmlElement.Method("lang").Ignore()
	htmlElement.Method("translate").Ignore()
	htmlElement.Method("dir").Ignore()
	htmlElement.Method("inert").Ignore()
	htmlElement.Method("accessKey").Ignore()
	htmlElement.Method("accessKeyLabel").Ignore()
	htmlElement.Method("draggable").Ignore()
	htmlElement.Method("spellcheck").Ignore()
	htmlElement.Method("writingSuggestions").Ignore()
	htmlElement.Method("autocorrect").Ignore()
	htmlElement.Method("autocapitalize").Ignore()
	htmlElement.Method("innerText").Ignore()
	htmlElement.Method("outerText").Ignore()
	htmlElement.Method("popover").Ignore()

	htmlSvg := htmlSpecs.Type("HTMLOrSVGElement")
	htmlSvg.Method("focus").SetCustomImplementation()
	htmlSvg.Method("dataset").SetNotImplemented()

	form := htmlSpecs.Type("HTMLFormElement")
	form.Method("requestSubmit").Argument("submitter").SetHasDefault()
	form.Method("reset").SetNotImplemented()
	form.Method("checkValidity").SetNotImplemented()
	form.Method("reportValidity").SetNotImplemented()
	form.Method("acceptCharset").SetNotImplemented()
	form.Method("autocomplete").SetNotImplemented()
	form.Method("enctype").SetNotImplemented()
	form.Method("encoding").SetNotImplemented()
	form.Method("target").SetNotImplemented()
	form.Method("rel").SetNotImplemented()
	form.Method("relList").SetNotImplemented()
	form.Method("length").SetNotImplemented()

	form.Method("name").Ignore()
	form.Method("noValidate").Ignore()

	input := htmlSpecs.Type("HTMLInputElement")
	input.Method("select").Ignore()
	input.Method("stepUp").Ignore()
	input.Method("stepDown").Ignore()
	input.Method("reportValidity").Ignore()
	input.Method("selectionRangeDirection").Ignore()
	input.Method("showPicker").Ignore()
	input.Method("accept").Ignore()
	input.Method("alpha").Ignore()
	input.Method("alt").Ignore()
	input.Method("autocomplete").Ignore()
	input.Method("checked").Ignore()
	input.Method("colorSpace").Ignore()
	input.Method("disabled").Ignore()
	input.Method("form").Ignore()
	input.Method("formAction").Ignore()
	input.Method("formMethod").Ignore()
	input.Method("formTarget").Ignore()
	input.Method("formEnctype").Ignore()
	input.Method("files").Ignore()
	input.Method("height").Ignore()
	input.Method("max").Ignore()
	input.Method("list").Ignore()
	input.Method("maxLength").Ignore()
	input.Method("minLength").Ignore()
	input.Method("multiple").Ignore()
	input.Method("setName").Ignore()
	input.Method("setPattern").Ignore()
	input.Method("placeholder").Ignore()
	input.Method("readonly").Ignore()
	input.Method("required").Ignore()
	input.Method("size").Ignore()
	input.Method("src").Ignore()
	input.Method("step").Ignore()
	input.Method("defaultValue").Ignore()
	input.Method("width").Ignore()
	input.Method("validationMessage").Ignore()
	input.Method("labels").Ignore()
	input.Method("selectionStart").Ignore()
	input.Method("selectionEnd").Ignore()
	input.Method("selectionDirection").Ignore()
	input.Method("willValidate").Ignore()
	input.Method("validity").Ignore()
	input.Method("valueAsNumber").Ignore()
	input.Method("valueAsDate").Ignore()
	input.Method("customValidity").Ignore()
	input.Method("setCustomValidity").Ignore()
	input.Method("setRangeText").Ignore()
	input.Method("setSelectionRange").Ignore()
	input.Method("defaultChecked").Ignore()
	input.Method("dirName").Ignore()
	input.Method("formNoValidate").Ignore()
	input.Method("indeterminate").Ignore()
	input.Method("min").Ignore()
	input.Method("name").Ignore()
	input.Method("value").Ignore()
	input.Method("readOnly").Ignore()
	input.Method("pattern").Ignore()

	window := htmlSpecs.Type("Window")

	window.Method("window").SetCustomImplementation()
	window.Method("location").Ignore()
	window.Method("parent").Ignore() // On `Node`
	window.Method("history").SetCustomImplementation()

	window.Method("prompt").Ignore()
	window.Method("close").SetNotImplemented()
	window.Method("stop").SetNotImplemented()
	window.Method("focus").SetNotImplemented()
	window.Method("blur").SetNotImplemented()
	window.Method("open").SetNotImplemented()
	window.Method("alert").SetNotImplemented()
	window.Method("confirm").SetNotImplemented()
	window.Method("postMessage").SetNotImplemented()
	window.Method("print").SetNotImplemented()
	window.Method("self").SetNotImplemented()
	window.Method("name").SetNotImplemented()
	window.Method("personalbar").SetNotImplemented()
	window.Method("locationbar").SetNotImplemented()
	window.Method("menubar").SetNotImplemented()
	window.Method("scrollbars").SetNotImplemented()
	window.Method("statusbar").SetNotImplemented()
	window.Method("status").SetNotImplemented()
	window.Method("toolbar").SetNotImplemented()
	window.Method("navigation").SetNotImplemented()
	window.Method("customElements").SetNotImplemented()
	window.Method("closed").SetNotImplemented()
	window.Method("frames").SetNotImplemented()
	window.Method("navigator").SetNotImplemented()
	window.Method("frames").SetNotImplemented()
	window.Method("top").SetNotImplemented()
	window.Method("opener").SetNotImplemented()
	window.Method("frameElement").SetNotImplemented()
	window.Method("clientInformation").SetNotImplemented()
	window.Method("originAgentCluster").SetNotImplemented()
	window.Method("length").SetNotImplemented()

	history := htmlSpecs.Type("History")
	history.OverrideWrappedType = &GoType{
		Name:    "History",
		Package: packagenames.HTMLInterfaces,
	}
	// We need to customize the inner type. This is a struct pointer
	// history.SkipWrapper = true
	history.Method("go").Argument("delta").HasDefaultValue("defaultDelta")
	history.Method("pushState").Argument("url").HasDefaultValue("defaultUrl")
	history.Method("pushState").Argument("unused").Ignore()
	history.Method("replaceState").Argument("url").HasDefaultValue("defaultUrl")
	history.Method("replaceState").Argument("unused").Ignore()
	history.Method("scrollRestoration").Ignore()
	// history.Method("state").SetEncoder("toJSON")

	htmlSpecs.Type("HTMLHyperlinkElementUtils")

	anchor := htmlSpecs.Type("HTMLAnchorElement")
	// anchor.wra
	anchor.Method("download").Ignore()
	anchor.Method("Ping").Ignore()
	anchor.Method("ping").Ignore()
	anchor.Method("rel").Ignore()
	anchor.Method("hreflang").Ignore()
	anchor.Method("type").Ignore()
	anchor.Method("referrerPolicy").Ignore()
	anchor.Method("relList").Ignore()
	anchor.Method("text").Ignore()

	// htmlSpecs.Type("HTMLHyperlinkElementUtils")

	return specs
}

func configureMutationObserver(domSpecs *WebIdlConfiguration) {
	domSpecs.Type("MutationObserver")
	domSpecs.Type("MutationRecord")
}
