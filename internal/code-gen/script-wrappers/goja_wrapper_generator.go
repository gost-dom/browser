package wrappers

import (
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
)

func NewGojaWrapperModuleGenerator() ScriptWrapperModulesGenerator {
	specs := configuration.CreateSpecs()
	dom := specs.Module("dom")
	domNode := dom.Type("Node")
	domNode.Method("childNodes").SetNotImplemented()

	html := specs.Module("html")
	html.SetMultipleFiles(true)
	location := html.Type("Location")
	location.Method("setHref").SetNotImplemented()
	location.Method("setSearch").SetNotImplemented()
	location.Method("setHash").SetNotImplemented()
	location.Method("setPathname").SetNotImplemented()
	location.Method("setHost").SetNotImplemented()
	location.Method("setHostname").SetNotImplemented()
	location.Method("setPort").SetNotImplemented()
	location.Method("setProtocol").SetNotImplemented()
	location.Method("assign").SetNotImplemented()
	location.Method("reload").SetNotImplemented()
	location.Method("replace").SetNotImplemented()
	location.Method("ancestorOrigins").SetNotImplemented()

	return ScriptWrapperModulesGenerator{
		Specs:            specs,
		PackagePath:      packagenames.Gojahost,
		TargetGenerators: GojaTargetGenerators{},
	}
}
