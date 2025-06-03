package v8host

import (
	"github.com/gost-dom/browser/html"

	v8 "github.com/gost-dom/v8go"
)

type htmlDocumentV8Wrapper struct {
	documentV8Wrapper
}

func newHTMLDocumentV8Wrapper(host *V8ScriptHost) htmlDocumentV8Wrapper {
	return htmlDocumentV8Wrapper{*newDocumentV8Wrapper(host)}
}

func createHTMLDocumentPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newDocumentV8Wrapper(host)
	builder := newIllegalConstructorBuilder[html.HTMLDocument](host)
	constructor := builder.constructor
	instanceTemplate := constructor.InstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	wrapper.CustomInitializer(
		v8Class{host, constructor, constructor.PrototypeTemplate(), constructor.InstanceTemplate()},
	)
	return constructor
}
