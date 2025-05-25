package wrappers

import (
	"github.com/gost-dom/code-gen/packagenames"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

var IdlTypeUndefined = idl.Type{Name: "undefined", Kind: idl.KindSimple}

var callbackContextType = g.NewTypePackage("CallbackContext", packagenames.JSAbstraction)
var callbackReturnType = g.NewTypePackage("CallbackRVal", packagenames.JSAbstraction)
var As = g.NewTypePackage("As", packagenames.JS)
var CallbackRVal = g.NewTypePackage("CallbackRVal", packagenames.JS)
var decoders = g.NewValuePackage("Decoders", packagenames.JSAbstraction)
