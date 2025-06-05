package wrappers

import (
	"github.com/gost-dom/code-gen/packagenames"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

var IdlTypeUndefined = idl.Type{Name: "undefined", Kind: idl.KindSimple}

var callbackContextType = g.NewTypePackage("CallbackContext", packagenames.JS)
var callbackReturnType = g.NewTypePackage("CallbackRVal", packagenames.JS)
var As = g.NewTypePackage("As", packagenames.JS)

var JSValue = g.NewTypePackage("Value", packagenames.JS)
var decoders = g.NewValuePackage("Decoders", packagenames.JS)
var JSRegister = g.NewValuePackage("RegisterClass", packagenames.JS)
var JSClassBiulder = g.NewValuePackage("ClassBuilder", packagenames.JS)
