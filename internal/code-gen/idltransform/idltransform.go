package idltransform

import (
	"slices"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/packagenames"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

// IdlType wraps an [idl.Type] value and can generate the proper type
// specification.
type IdlType struct {
	idl.Type
}

// Creates a new IdlType that will always be qualified
func NewIdlType(t idl.Type) IdlType { return IdlType{t} }

func InternalPackage(name string) string {
	switch name {
	case "AbortController", "AbortSignal":
		return packagenames.DomInterfaces
	case "EventHandler", "EventTarget":
		return packagenames.Events
	case "DOMTokenList", "NodeList", "Node":
		return packagenames.Dom
	default:
		return ""
	}
}

func TypeGen(name string) g.Generator {
	if pkg := InternalPackage(name); pkg != "" /* && pkg != targetPkg */ {
		return g.NewTypePackage(name, pkg)
	}
	return g.Id(name)
}

func (s IdlType) Generate() *jen.Statement {
	if pkg := InternalPackage(s.Name); pkg != "" /* && pkg != s.TargetPackage */ {
		return jen.Qual(pkg, s.Name)
	}
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
	case "DOMString", "USVString", "ByteString":
		return true
	}
	return false
}

func (t IdlType) IsInt() bool {
	switch t.Name {
	case "unsigned long", "long", "unsigned short", "short":
		return true
	}
	return false
}

func (t IdlType) IsBoolean() bool {
	switch t.Name {
	case "boolean":
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
	return jen.Op("[]").Add(IdlType{Type: *t.TypeParam}.Generate())
}

func FilterType(t idl.Type) idl.Type {
	if t.Kind != idl.KindUnion {
		return t
	}
	t.Types = slices.DeleteFunc(slices.Clone(t.Types), unsupportedType)
	if len(t.Types) > 1 {
		return t
	}
	if len(t.Types) == 0 {
		panic("Filtering type resulted in nothing")
	}
	return t.Types[0]
}

func unsupportedType(t idl.Type) bool {
	switch t.Name {
	case "TrustedHTML":
		// TrustedHTML isn't prioritized and isn't widely supported yet
		// https://github.com/gost-dom/browser/issues/151
		return true
	default:
		return false
	}
}

// StructFieldType represents the type of a struct field
type StructFieldType struct {
	idl.Type
}

func (s StructFieldType) Generate() *jen.Statement {
	if IdlType(s).IsString() && s.Nullable {
		return jen.Op("*").Id("string")
	}
	return IdlType(s).Generate()
}
