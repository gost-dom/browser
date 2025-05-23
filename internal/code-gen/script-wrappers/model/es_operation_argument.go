package model

import (
	"fmt"

	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
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
	return a.Optional && !hasDefault && !a.CustomRule.Variadic
}

func (a ESOperationArgument) DefaultValueInGo() (name string, ok bool) {
	ok = a.Optional && a.ArgumentSpec.HasDefault
	if defaultValue := a.ArgumentSpec.DefaultValue; defaultValue != "" {
		name = defaultValue
	} else {
		name = fmt.Sprintf("default%s", a.IdlArg.Type.Name)
	}
	return
}

func (a ESOperationArgument) GoTypeName() string {
	switch a.IdlArg.Type.Name {
	case "DOMString", "USVString", "ByteString":
		return "string"
	default:
		return a.IdlArg.Type.Name
	}
}
