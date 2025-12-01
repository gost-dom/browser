package scripting

import (
	"fmt"

	"github.com/gost-dom/code-gen/packagenames"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

var (
	// Exported values from: scripting/internal/js

	jsParseSetterArg     = g.NewValuePackage("ParseSetterArg", packagenames.JS)
	jsConsumeArg         = g.NewValuePackage("ConsumeArgument", packagenames.JS)
	jsConsumeOptionalArg = g.NewValuePackage("ConsumeOptionalArg", packagenames.JS)
	jsConsumeRestArgs    = g.NewValuePackage("ConsumeRestArguments", packagenames.JS)
	jsThisLogAttr        = g.NewValuePackage("ThisLogAttr", packagenames.JS)
	jsArgsLogAttr        = g.NewValuePackage("ArgsLogAttr", packagenames.JS)
	jsLogAttr            = g.NewValuePackage("LogAttr", packagenames.JS)

	logErrAttr = g.NewValuePackage("ErrAttr", packagenames.Log)

	jsAs            = g.NewTypePackage("As", packagenames.JS)
	jsValue         = g.NewTypePackage("Value", packagenames.JS).TypeParam(g.Id("T"))
	jsClass         = g.NewTypePackage("Class", packagenames.JS).TypeParam(g.Id("T"))
	jsCbCtx         = g.NewTypePackage("CallbackContext", packagenames.JS).TypeParam(g.Id("T"))
	jsScriptEngine  = g.NewTypePackage("ScriptEngine", packagenames.JS).TypeParam(g.Id("T"))
	jsRegisterClass = g.NewValuePackage("RegisterClass", packagenames.JS)
	jsClassBuilder  = g.NewValuePackage("ClassBuilder", packagenames.JS)

	errorsFirst = g.NewValuePackage("First", packagenames.Errors)

	slogString = g.NewValuePackage("String", packagenames.StdSlog)

	// Exported values from: scripting/internal/codec

	zeroValue = g.NewValuePackage("ZeroValue", packagenames.Codec)

	// Codecs
	EncodeCallbackErrorf = g.NewValuePackage("EncodeCallbackErrorf", packagenames.Codec)
)

var IdlTypeUndefined = idl.Type{Name: "undefined", Kind: idl.KindSimple}

func ConstructorNameForInterface(name string) string {
	return fmt.Sprintf("New%s", name)
}
