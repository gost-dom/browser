package gojagen

import (
	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/packagenames"
	g "github.com/gost-dom/generators"
)

const (
	gojaSrc = packagenames.Goja
)

var (
	gojaFc      = g.Raw(jen.Qual(gojaSrc, "FunctionCall"))
	gojaValue   = g.Raw(jen.Qual(gojaSrc, "Value"))
	gojaObj     = g.Raw(jen.Op("*").Qual(gojaSrc, "Object"))
	gojaRuntime = g.Raw(jen.Op("*").Qual(gojaSrc, "Runtime"))
	gojaCbCtx   = g.NewType("callbackContext").Pointer()
	flagTrue    = g.Raw(jen.Qual(gojaSrc, "FLAG_TRUE"))
)
