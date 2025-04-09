package model

import (
	"fmt"

	"github.com/gost-dom/code-gen/script-wrappers/configuration"
	"github.com/gost-dom/webref/idl/legacy"
)

type ESOperationArgument struct {
	Name         string
	Type         string
	Optional     bool
	Variadic     bool
	IdlType      legacy.IdlTypes
	ArgumentSpec configuration.ESMethodArgument
	Ignore       bool
}

func (a ESOperationArgument) OptionalInGo() bool {
	hasDefault := a.ArgumentSpec.HasDefault
	return a.Optional && !hasDefault
}

func (a ESOperationArgument) DefaultValueInGo() (name string, ok bool) {
	ok = a.Optional && a.ArgumentSpec.HasDefault
	if defaultValue := a.ArgumentSpec.DefaultValue; defaultValue != "" {
		name = defaultValue
	} else {
		name = fmt.Sprintf("default%s", a.Type)
	}
	return
}
