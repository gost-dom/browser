package model

import (
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
	"github.com/gost-dom/webref/idl"
)

type ESOperation struct {
	Name                 string
	NotImplemented       bool
	LegacyRetType        idl.RetType
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

func encoderForIDLType(t idl.Type) string {
	converter := "to"
	if t.Kind == idl.KindSequence {
		converter += "Sequence"
		t = *t.TypeParam
	}
	if t.Nullable {
		converter += "Nullable"
	}
	converter += IdlNameToGoName(t.Name)
	return converter
}

func (o ESOperation) Encoder() string {
	if e := o.MethodCustomization.Encoder; e != "" {
		return e
	}
	return encoderForIDLType(o.RetType)
}

func (o ESOperation) RetTypeName() string {
	return IdlNameToGoName(o.LegacyRetType.TypeName)
}
