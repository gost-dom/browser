package configuration

import (
	"github.com/gost-dom/webref/idl"
)

// WebAPIConfig configures the generation of JavaScript bindings for a specific
// web API.
type WebAPIConfig struct {
	// Name is the name of the specification, corresponds to the file name.
	Name string
	// Interfaces defines the names of the specified interfaces for which to
	// generate specifications
	Interfaces map[string]*WebIDLConfig
	// PartialSearchModules tells which "modules" or "web api specs" to search
	// for partial interfaces.
	//
	// Example: "interface Element" is defined in dom.idl, but innerHTML and
	// outerHTML is defined in a partial interface in html.idl.
	PartialSearchModules []string
}

func NewWebAPIConfig(name string) *WebAPIConfig {
	return &WebAPIConfig{
		Name:       name,
		Interfaces: make(map[string]*WebIDLConfig),
	}
}

func (spec WebAPIConfig) Types() []*WebIDLConfig {
	types := make([]*WebIDLConfig, len(spec.Interfaces))
	idx := 0
	for _, t := range spec.Interfaces {
		types[idx] = t
		idx++
	}
	return types
}

func (c *WebAPIConfig) AddSearchModule(name string) {
	c.PartialSearchModules = append(c.PartialSearchModules, name)
}

// LoadSpec loads relevant [idl.Spec] for the configured package
func LoadSpecs(c *WebAPIConfig) (spec idl.Spec, err error) {
	if spec, err = idl.Load(c.Name); err != nil {
		return
	}
	return
}

func (s *WebAPIConfig) Type(typeName string) *WebIDLConfig {
	if result, ok := s.Interfaces[typeName]; ok {
		return result
	}
	result := &WebIDLConfig{
		DomSpec:  s,
		TypeName: typeName,
	}
	result.ensureMap()
	s.Interfaces[typeName] = result
	return result
}
