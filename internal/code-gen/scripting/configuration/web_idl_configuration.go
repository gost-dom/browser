package configuration

import (
	"cmp"
	"slices"
)

// WebAPIConfig configures the generation of JavaScript bindings for a specific
// web API.
type WebAPIConfig struct {
	// Name is the name of the specification, corresponds to the file name.
	Name string
	// Interfaces defines the names of the specified interfaces for which to
	// generate specifications
	Interfaces map[string]*WebIDLConfig
}

func NewWebAPIConfig(name string) *WebAPIConfig {
	return &WebAPIConfig{
		Name:       name,
		Interfaces: make(map[string]*WebIDLConfig),
	}
}

func (spec WebAPIConfig) GetTypesSorted() []*WebIDLConfig {
	types := make([]*WebIDLConfig, len(spec.Interfaces))
	idx := 0
	for _, t := range spec.Interfaces {
		types[idx] = t
		idx++
	}
	slices.SortFunc(types, func(x, y *WebIDLConfig) int {
		return cmp.Compare(x.TypeName, y.TypeName)
	})
	return types
}
