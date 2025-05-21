package model

import (
	"strings"

	htmlelements "github.com/gost-dom/code-gen/html-elements"
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
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

func encoderForIDLType(t idl.Type) string {
	converter := "to"
	if t.Kind == idl.KindSequence {
		converter += "Sequence"
		t = *t.TypeParam
	}
	if t.Nullable && !htmlelements.IdlType(t).Nillable() {
		converter += "Nullable"
	}
	converter += IdlNameToGoName(idlTypeNameToGoName(t))
	return converter
}

func (o ESOperation) Encoder() string {
	if e := o.MethodCustomization.Encoder; e != "" {
		return e
	}
	return encoderForIDLType(o.RetType)
}

func (o ESOperation) RetTypeName() string {
	return o.RetType.Name
}
