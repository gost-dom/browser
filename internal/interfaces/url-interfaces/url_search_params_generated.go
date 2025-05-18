// This file is generated. Do not edit.

package urlinterfaces

import (
	"fmt"
	"iter"
)

type URLSearchParams interface {
	fmt.Stringer
	Size() int
	Append(string, string)
	Delete(string)
	DeleteValue(string, string)
	Get(string) (string, bool)
	GetAll(string) []string
	Has(string) bool
	HasValue(string, string) bool
	Set(string, string)
	Sort()
	All() iter.Seq2[string, string]
}
