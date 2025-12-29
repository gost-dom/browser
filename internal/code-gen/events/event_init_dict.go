package events

import (
	"fmt"

	"github.com/gost-dom/code-gen/idltransform"
	"github.com/gost-dom/code-gen/internal"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

func GenerateEventInit(name, pkg string) (g.Generator, error) {
	spec, err := idl.Load(pkg)
	if err != nil {
		return nil, fmt.Errorf("GenerateEventInit: load pkg %s: %w", pkg, err)
	}
	dict, ok := spec.Dictionaries[name]
	if !ok {
		return nil, fmt.Errorf("GenerateEventInit: %s: dictionary not found", name)
	}
	return GenerateEventInitDict(name, dict), nil
}

func GenerateEventInitDict(name string, dict idl.Dictionary) g.Generator {
	res := g.NewStruct(g.Id(name))
	res.Embed(g.Id(dict.Inheritance))
	for _, entry := range dict.Entries {
		t := idltransform.IdlType{Type: entry.Value}
		res.Field(g.Id(internal.IdlNameToGoName(entry.Key)), t)
	}
	return res
}
