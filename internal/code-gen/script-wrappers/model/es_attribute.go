package model

type ESAttribute struct {
	Name   string
	Getter *ESOperation
	Setter *ESOperation
}
