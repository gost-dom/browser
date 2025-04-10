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

func (o ESOperation) Encoder() string {
	if e := o.MethodCustomization.Encoder; e != "" {
		return e
	}
	converter := "to"
	if o.LegacyRetType.Nullable {
		converter += "Nullable"
	}
	converter += IdlNameToGoName(o.LegacyRetType.TypeName)
	return converter
}

func (o ESOperation) RetTypeName() string {
	return IdlNameToGoName(o.LegacyRetType.TypeName)
}
