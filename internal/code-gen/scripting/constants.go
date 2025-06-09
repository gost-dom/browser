package scripting

import (
	"github.com/gost-dom/code-gen/packagenames"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

var (
	IdlTypeUndefined = idl.Type{Name: "undefined", Kind: idl.KindSimple}

	callbackContextType = g.NewTypePackage("CallbackContext", packagenames.JS)
	callbackReturnType  = g.NewTypePackage("CallbackRVal", packagenames.JS)
	As                  = g.NewTypePackage("As", packagenames.JS)
	JSValue             = g.NewTypePackage("Value", packagenames.JS).TypeParam(g.Id("T"))
	decoders            = g.NewValuePackage("Decoders", packagenames.JS)
	JSRegister          = g.NewValuePackage("RegisterClass", packagenames.JS)
	JSClassBiulder      = g.NewValuePackage("ClassBuilder", packagenames.JS)

	v8Class = g.NewTypePackage("Class", packagenames.JS).TypeParam(g.Id("T"))
	v8CbCtx = g.NewTypePackage("CallbackContext", packagenames.JS).TypeParam(g.Id("T"))

	scriptEngine         = g.NewTypePackage("ScriptEngine", packagenames.JS).TypeParam(g.Id("T"))
	parseSetterArg       = g.NewValuePackage("ParseSetterArg", packagenames.JS)
	ConsumeArgument      = g.NewValuePackage("ConsumeArgument", packagenames.JS)
	ConsumeOptionalArg   = g.NewValuePackage("ConsumeOptionalArg", packagenames.JS)
	ConsumeRestArguments = g.NewValuePackage("ConsumeRestArguments", packagenames.JS)
	ZeroValue            = g.NewValuePackage("ZeroValue", packagenames.Codec)
)
