package v8host

import (
	"github.com/gost-dom/browser/html"
)

type htmlDocumentV8Wrapper struct {
	documentV8Wrapper
}

func newHTMLDocumentV8Wrapper(host *V8ScriptHost) htmlDocumentV8Wrapper {
	return htmlDocumentV8Wrapper{*newDocumentV8Wrapper(host)}
}

func createHTMLDocumentPrototype(host *V8ScriptHost) v8Class {
	wrapper := newDocumentV8Wrapper(host)
	builder := newIllegalConstructorBuilder[html.HTMLDocument](host)
	constructor := builder.constructor
	instanceTemplate := constructor.InstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	wrapper.CustomInitializer(
		v8Class{host, constructor, constructor.PrototypeTemplate(), constructor.InstanceTemplate()},
	)
	return newV8Class(host, constructor)
}
