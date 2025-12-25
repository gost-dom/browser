package gotypes

import "github.com/gost-dom/code-gen/customrules"

type GoType = customrules.GoType

var TimeDuration = GoType{
	Name:    "Duration",
	Package: "time",
}
