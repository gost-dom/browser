package model

import (
	"fmt"

	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/idltransform"
	"github.com/gost-dom/code-gen/scripting/configuration"
	"github.com/gost-dom/webref/idl"
)

type ESOperationArgument struct {
	Name         string
	IdlArg       idl.Argument
	Type         string
	Optional     bool
	Variadic     bool
	ArgumentSpec configuration.ESMethodArgument
	CustomRule   customrules.ArgumentRule
	Ignore       bool
}

func (a ESOperationArgument) OptionalInGo() bool {
	hasDefault := a.ArgumentSpec.HasDefault
	return a.Optional && !hasDefault && !a.VariadicInGo()
}

func (a ESOperationArgument) VariadicInGo() bool {
	return a.CustomRule.Variadic || a.Variadic
}

func (a ESOperationArgument) idlType() idltransform.IdlType {
	return idltransform.NewIdlType(a.IdlArg.Type)
}

func (a ESOperationArgument) IsString() bool  { return a.idlType().IsString() }
func (a ESOperationArgument) IsBoolean() bool { return a.idlType().IsBoolean() }
func (a ESOperationArgument) IsInt() bool     { return a.idlType().IsInt() }

func (a ESOperationArgument) DefaultValueInGo() (name string, ok bool) {
	ok = a.Optional && a.ArgumentSpec.HasDefault
	if defaultValue := a.ArgumentSpec.DefaultValue; defaultValue != "" {
		name = defaultValue
	} else {
		if n := a.IdlArg.Type.Name; n == "" {
			name = fmt.Sprintf("default%s", IdlNameToGoName(a.Name))
		} else {
			name = fmt.Sprintf("default%s", a.IdlArg.Type.Name)
		}
	}
	return
}

// Returns whether the argument is specified as nullable in the IDL
// specification.
func (a ESOperationArgument) NullableInIDL() bool {
	return a.IdlArg.Type.Nullable
}

func (a ESOperationArgument) HasDefault() bool {
	return a.ArgumentSpec.HasDefault
}
