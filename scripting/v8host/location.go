package v8host

import (
	"github.com/gost-dom/browser/html"
	v8 "github.com/tommie/v8go"
)

func createLocationPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	builder := newIllegalConstructorBuilder[html.Location](host)
	builder.instanceLookup = func(ctx *V8ScriptContext, this *v8.Object) (html.Location, error) {
		location := ctx.window.Location()
		return location, nil
	}
	helper := builder.NewPrototypeBuilder()
	helper.CreateReadonlyProp("hash", html.Location.Hash)
	helper.CreateReadonlyProp("host", html.Location.Host)
	helper.CreateReadonlyProp("hostname", html.Location.Hostname)
	helper.CreateReadonlyProp("href", html.Location.Href)
	helper.CreateReadonlyProp("origin", html.Location.Origin)
	helper.CreateReadonlyProp("pathname", html.Location.Pathname)
	helper.CreateReadonlyProp("port", html.Location.Port)
	helper.CreateReadonlyProp("protocol", html.Location.Protocol)
	helper.CreateReadonlyProp("search", html.Location.Search)
	return builder.constructor
}
