package htmlelements

import (
	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/webref/idl"
)

type IdlType idl.Type

func (s IdlType) Generate() *jen.Statement {
	switch s.Kind {
	case idl.KindSequence:
		return s.generateSequence()
	}
	switch {
	case s.IsString():
		return jen.Id("string")
	case s.IsInt():
		return jen.Id("int")
	}
	switch string(s.Name) {
	case "boolean":
		return jen.Id("bool")
	case "DOMTokenList", "NodeList", "Node":
		return jen.Qual(packagenames.Dom, s.Name)
	case "undefined":
		return nil
	default:
		return jen.Id(s.Name)
	}
}

func (t IdlType) IsString() bool {
	switch t.Name {
	case "DOMString", "USVString":
		return true
	}
	return false
}

func (t IdlType) IsInt() bool {
	switch t.Name {
	case "unsigned long", "long":
		return true
	}
	return false
}

// Nillable determins if values of the type can be nil. This is useful for
// return values, such as `element.GetAttribute()`, where the DOM specifies a
// null return value if the value is not found, but the more idiomatic Go
// signature would be to return (string, bool) rather then *string
func (t IdlType) Nillable() bool {
	switch {
	case t.IsString():
		return false
	}
	return true
}

func (t IdlType) generateSequence() *jen.Statement {
	if t.TypeParam == nil {
		panic("IdlType.generateSequence: TypeParameter is nil for sequence type")
	}
	return jen.Op("[]").Add(IdlType(*t.TypeParam).Generate())
}
