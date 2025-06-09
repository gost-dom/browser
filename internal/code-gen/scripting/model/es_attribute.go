package model

import "github.com/gost-dom/webref/idl"

type ESAttribute struct {
	Name   string
	Spec   idl.Attribute
	Getter *ESOperation
	Setter *ESOperation
}
