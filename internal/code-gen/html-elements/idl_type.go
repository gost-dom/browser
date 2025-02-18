package htmlelements

import (
	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/webref/idl"
)

type IdlType idl.Type

func (s IdlType) Generate() *jen.Statement {
	switch string(s.Name) {
	case "boolean":
		return jen.Id("bool")
	case "DOMString", "USVString":
		return jen.Id("string")
	case "unsigned long":
		return jen.Id("int")
	case "DOMTokenList":
		return jen.Qual(packagenames.Dom, "DOMTokenList")
	case "undefined":
		return nil
	default:
		return jen.Id(s.Name)
	}
}
