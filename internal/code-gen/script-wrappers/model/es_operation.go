package model

import (
	"strings"

	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/idltransform"
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

type ESOperation struct {
	Name                 string
	Spec                 idl.Operation
	NotImplemented       bool
	RetType              idl.Type
	HasError             bool
	CustomImplementation bool
	MethodCustomization  configuration.ESMethodWrapper
	Arguments            []ESOperationArgument
}

func (o ESOperation) CallbackMethodName() string {
	return idl.SanitizeName(o.Name)
}

func (op ESOperation) GetHasError() bool {
	return op.HasError
}

func (op ESOperation) HasResult() bool {
	return op.RetType.Name != "undefined"
}

func IsNodeType(typeName string) bool {
	loweredName := strings.ToLower(typeName)
	switch loweredName {
	case "node":
		return true
	case "document":
		return true
	case "documentfragment":
		return true
	}
	if strings.HasSuffix(loweredName, "element") {
		return true
	}
	return false
}

func idlTypeNameToGoName(t idl.Type) string {
	switch t.Name {
	case "DOMString", "USVString", "ByteString":
		{
			return "string_"
		}
	default:
		return t.Name
	}
}

func (o ESOperation) EncodeAsSimpleJSLookup() bool {
	if IsNodeType(o.RetType.Name) {
		return true
	}
	switch o.RetType.Name {
	case "Attr", "NodeList", "HTMLFormControlsCollection":
		return true
	default:
		return false
	}
}

func (o ESOperation) Encoder(data ESConstructorData) string {
	if e := o.MethodCustomization.Encoder; e != "" {
		return e
	}
	t := o.RetType
	converter := "to"
	if t.Kind == idl.KindSequence {
		converter += "Sequence"
		t = *t.TypeParam
	}
	if t.Nullable && !idltransform.IdlType(t).Nillable() {
		if data.CustomRule.OutputType == customrules.OutputTypeStruct {
			converter += "Nullable"
		} else {
			converter += "Nillable"
		}
	}
	if o.EncodeAsSimpleJSLookup() {
		converter += "JSWrapper"
	} else {
		converter += IdlNameToGoName(idlTypeNameToGoName(t))
	}
	return converter
}

func (o ESOperation) RetValues(data ESConstructorData) []g.Generator {
	if !o.HasResult() {
		return nil
	}
	t := o.RetType
	res := g.Id("result")
	hasValue := g.Id("hasValue")
	if t.Nullable && !idltransform.IdlType(t).Nillable() {
		if data.CustomRule.OutputType == customrules.OutputTypeStruct {
			return g.List(res)
		} else {
			return g.List(res, hasValue)
		}
	}
	return g.List(res)
}

func (o ESOperation) RetTypeName() string {
	return o.RetType.Name
}
