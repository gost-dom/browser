package v8gen

import (
	"github.com/dave/jennifer/jen"

	g "github.com/gost-dom/generators"
)

const v8 = "github.com/gost-dom/v8go"

var (
	v8FunctionTemplatePtr     = g.NewTypePackage("FunctionTemplate", v8).Pointer()
	v8FunctionCallbackInfoPtr = g.NewTypePackage("FunctionCallbackInfo", v8).Pointer()
	v8ObjectTemplatePtr       = g.NewTypePackage("ObjectTemplate", v8).Pointer()
	v8Value                   = g.NewTypePackage("Value", v8).Pointer()
	v8ReadOnly                = g.Raw(jen.Qual(v8, "ReadOnly"))
	v8None                    = g.Raw(jen.Qual(v8, "None"))
	v8CbCtx                   = g.NewType("jsCallbackContext")
	v8Class                   = g.NewType("jsClass")
	scriptHostPtr             = g.NewType("jsScriptEngine")
	// scriptHostPtr             = g.NewType("V8ScriptHost").Pointer()
)

// Provides helpers for functions that needs an iso as the first argument
type v8Iso struct{ g.Value }

func (iso v8Iso) NewFunctionTemplate(cb g.Generator) g.Generator {
	return g.NewValuePackage("NewFunctionTemplateWithError", v8).Call(iso.Value, cb)
}
