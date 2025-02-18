// This file is generated. Do not edit.

package url

type URLSearchParams interface {
	Size() int
	Append(string, string)
	Delete(string, string)
	Get(string) string
	GetAll(string)
	Has(string, string) bool
	Set(string, string)
	Sort()
}
